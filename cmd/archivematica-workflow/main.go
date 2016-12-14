package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/artefactual-labs/archivematica-workflow/pkg/endpoints"
	grpctransport "github.com/artefactual-labs/archivematica-workflow/pkg/grpc"
	httptransport "github.com/artefactual-labs/archivematica-workflow/pkg/http"
	"github.com/artefactual-labs/archivematica-workflow/pkg/service"
	"github.com/artefactual-labs/archivematica-workflow/pkg/service/pb"
	"github.com/artefactual-labs/archivematica-workflow/pkg/version"
)

func main() {
	var (
		debugAddr    = flag.String("debug.addr", ":40050", "Debug and metrics listen address")
		grpcAddr     = flag.String("grpc.addr", ":50050", "gRPC (HTTP/2) listen address")
		httpAddr     = flag.String("http.addr", ":60050", "HTTP listen address")
		printVersion = flag.Bool("version", false, "Print version")
	)
	flag.Parse()

	if *printVersion {
		if version.VERSION == "" {
			fmt.Fprintln(os.Stdout, "archivematica-workflow (unknown version)")
		} else {
			fmt.Fprintf(os.Stdout, "archivematica-workflow %s\n", version.VERSION)
		}
		os.Exit(1)
	}

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}
	logger.Log("msg", "Hello!", "debugAddr", debugAddr, "grpcAddr", grpcAddr, "httpAddr", httpAddr)
	defer logger.Log("msg", "Goodbye!")

	// Business domain.
	svc := service.New(logger)
	eps := endpoints.New(svc, logger)

	// Mechanical domain.
	errc := make(chan error)
	ctx := context.Background()

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// Debug listener.
	go func() {
		m := http.NewServeMux()
		m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
		m.Handle("/metrics", promhttp.Handler())

		errc <- http.ListenAndServe(*debugAddr, m)
	}()

	// gRPC transport.
	go func() {
		logger := log.NewContext(logger).With("transport", "grpc")

		ln, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errc <- err
			return
		}

		srv := grpctransport.NewHandler(ctx, eps, logger)
		s := grpc.NewServer()
		pb.RegisterWorkflowServer(s, srv)

		errc <- s.Serve(ln)
	}()

	// HTTP transport.
	go func() {
		logger := log.NewContext(logger).With("transport", "http")

		mux := httptransport.NewHandler(ctx, eps, logger)

		errc <- http.ListenAndServe(*httpAddr, mux)
	}()

	// Run!
	logger.Log("exit", <-errc)
}
