package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	"github.com/artefactual-labs/archivematica-workflow/endpoints"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// NewHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHandler(ctx context.Context, eps endpoints.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []httptransport.ServerOption{}

	r.Methods("GET").Path("/workflows/{id}").Handler(httptransport.NewServer(
		ctx,
		eps.WorkflowGetEndpoint,
		decodeGetWorkflowRequest,
		encodeGenericResponse,
		options...,
	))
	return r
}

func decodeGetWorkflowRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return endpoints.WorkflowGetRequest{ID: id}, nil
}

// encodeGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func encodeGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
