package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"strings"

	grpcclient "github.com/artefactual-labs/archivematica-workflow/client/grpc"
	"github.com/artefactual-labs/archivematica-workflow/service/pb"
	"github.com/artefactual-labs/archivematica-workflow/version"
)

func main() {
	var (
		grpcAddr     = flag.String("grpc.addr", ":50050", "gRPC (HTTP) address of archivematica-workflow")
		method       = flag.String("method", "", "WorkflowGet, ...")
		printVersion = flag.Bool("version", false, "Print version")
	)
	flag.Parse()

	if *printVersion {
		if version.VERSION == "" {
			fmt.Fprintln(os.Stdout, "archivematica-workflow-cli (unknown version)")
		} else {
			fmt.Fprintf(os.Stdout, "archivematica-workflow-cli %s\n", version.VERSION)
		}
		os.Exit(1)
	}

	if *method == "" {
		fmt.Fprintf(os.Stderr, "You need to use one of the methods available, e.g. WorkflowGet, ...\n")
		fmt.Fprintf(os.Stderr, "For example: archivemica-workflow-cli -method=WorkflowGet default\n")
		os.Exit(1)
	}

	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	svc := grpcclient.New(conn, log.NewNopLogger())

	switch *method {
	case "WorkflowGet":
		id := flag.Args()[0]
		wf, err := svc.WorkflowGet(context.Background(), id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		printWorkflow(wf)
	default:
		fmt.Fprintf(os.Stderr, "error: invalid method %q\n", *method)
		os.Exit(1)
	}
}

func printWorkflow(wfd *pb.WorkflowData) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.DiscardEmptyColumns)

	for _, item := range wfd.WatchedDirectories {
		row := []string{
			fmt.Sprintf("Watched directory: %s", item.Path),
			item.Type.String(),
			fmt.Sprintf("goes to Chain %s", item.ChainId),
		}
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	w.Flush()

	for id, item := range wfd.Chains {
		row := []string{
			fmt.Sprintf("Chain %s", id),
			item.Description["en"],
			fmt.Sprintf("goes to Link %s", item.LinkId),
		}
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	w.Flush()

	for id, item := range wfd.Links {
		row := []string{
			fmt.Sprintf("Link %s", id),
			item.Description["en"],
		}
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	w.Flush()

}
