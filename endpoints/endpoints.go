package endpoints

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats. It also includes endpoint middlewares.

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/artefactual-labs/archivematica-workflow/service"
	"github.com/artefactual-labs/archivematica-workflow/service/pb"
)

// Endpoints collects all of the endpoints that compose the service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them
// into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	WorkflowGetEndpoint endpoint.Endpoint
}

// New returns an Endpoints that wraps the provided server, and wires in all of
// the expected endpoint middlewares via the various parameters.
func New(svc service.Service, logger log.Logger) Endpoints {
	var workflowGetEndpoint endpoint.Endpoint
	{
		workflowGetEndpoint = MakeWorkflowEndpoint(svc)
		workflowGetEndpoint = LoggingMiddleware(logger)(workflowGetEndpoint)
	}
	return Endpoints{
		WorkflowGetEndpoint: workflowGetEndpoint,
	}
}

// WorkflowGet implements Service. Primarily useful in a client.
func (e Endpoints) WorkflowGet(ctx context.Context, id string) (*pb.WorkflowData, error) {
	req := WorkflowGetRequest{ID: id}
	resp, err := e.WorkflowGetEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(WorkflowGetResponse).Workflow, resp.(WorkflowGetResponse).Err
}

// MakeWorkflowEndpoint returns an endpoint that invokes Workflow on the
// service. Primarily useful in a server.
func MakeWorkflowEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WorkflowGetRequest)
		w, err := s.WorkflowGet(ctx, req.ID)
		return WorkflowGetResponse{Workflow: w, Err: err}, nil
	}
}

// WorkflowGetRequest collects the request parameters for the WorkflowGet
// method.
type WorkflowGetRequest struct {
	ID string
}

// WorkflowGetResponse collects the response values for the WorkflowGet method.
type WorkflowGetResponse struct {
	Workflow *pb.WorkflowData
	Err      error
}
