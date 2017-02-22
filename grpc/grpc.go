package grpc

import (
	"errors"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/artefactual-labs/archivematica-workflow/endpoints"
	"github.com/artefactual-labs/archivematica-workflow/service/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

// NewHandler returns a handler that makes a set of endpoints available as a
// gRPC WorkflowServer.
func NewHandler(ctx context.Context, eps endpoints.Endpoints, logger log.Logger) pb.WorkflowServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{
		workflow: grpctransport.NewServer(
			ctx,
			eps.WorkflowGetEndpoint,
			DecodeGRPCWorkflowGetRequest,
			EncodeGRPCWorkflowGetResponse,
			options...,
		),
	}
}

// grpcServer is our implementation of pb.WorkflowServer
type grpcServer struct {
	workflow grpctransport.Handler
}

func (s *grpcServer) WorkflowGet(ctx context.Context, req *pb.WorkflowGetRequest) (*pb.WorkflowGetResponse, error) {
	_, rep, err := s.workflow.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WorkflowGetResponse), nil
}

// Primarily useful in a client.

// EncodeGRPCWorkflowGetRequest is a transport/grpc.EncodeRequestFunc that
// converts a user-domain workflowGet request to a gRPC workflowGet request.
func EncodeGRPCWorkflowGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(endpoints.WorkflowGetRequest)
	return &pb.WorkflowGetRequest{Id: req.ID}, nil
}

// DecodeGRPCWorkflowGetResponse is a transport/grpc.DecodeResponseFunc that
// converts a gRPC workflowGet response to a user-domain workflowGet response.
func DecodeGRPCWorkflowGetResponse(_ context.Context, grpcResp interface{}) (interface{}, error) {
	resp := grpcResp.(*pb.WorkflowGetResponse)
	return endpoints.WorkflowGetResponse{Workflow: resp.Workflow, Err: str2err(resp.Error)}, nil
}

// Primarily useful in a server.

// EncodeGRPCWorkflowGetResponse is a transport/grpc.EncodeResponseFunc that
// converts a user-domain workflowGet response to a gRPC workflowGet response.
func EncodeGRPCWorkflowGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.WorkflowGetResponse)
	return &pb.WorkflowGetResponse{Workflow: resp.Workflow, Error: err2str(resp.Err)}, nil
}

// DecodeGRPCWorkflowGetRequest is a transport/grpc.DecodeRequestFunc that
// converts a gRPC workflow request to a user-domain workflowGet request.
func DecodeGRPCWorkflowGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.WorkflowGetRequest)
	return endpoints.WorkflowGetRequest{ID: req.Id}, nil
}

// These annoying helper functions are required to translate Go error types to
// and from strings, which is the type we use in our IDLs to represent errors.
// There is special casing to treat empty strings as nil errors.

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
