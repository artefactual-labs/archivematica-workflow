package grpc

import (
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"github.com/artefactual-labs/archivematica-workflow/endpoints"
	grpcsvc "github.com/artefactual-labs/archivematica-workflow/grpc"
	"github.com/artefactual-labs/archivematica-workflow/service"
	"github.com/artefactual-labs/archivematica-workflow/service/pb"
)

// New returns a Service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, logger log.Logger) service.Service {
	var workflowGetEndpoint = grpctransport.NewClient(
		conn,
		"Workflow",
		"WorkflowGet",
		grpcsvc.EncodeGRPCWorkflowGetRequest,
		grpcsvc.DecodeGRPCWorkflowGetResponse,
		pb.WorkflowGetResponse{},
	).Endpoint()

	return endpoints.Endpoints{
		WorkflowGetEndpoint: workflowGetEndpoint,
	}
}
