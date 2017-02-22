package service

import (
	"encoding/json"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"github.com/artefactual-labs/archivematica-workflow/service/dist"
	"github.com/artefactual-labs/archivematica-workflow/service/pb"
)

// Service provides operations on Archivematica workflows.
type Service interface {
	WorkflowGet(ctx context.Context, id string) (*pb.WorkflowData, error)
}

var (
	// ErrNotFound is used when the workflow cannot be found.
	ErrNotFound = errors.New("not found")
)

// New returns a basic Service with all of the expected middlewares wired in.
func New(logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService(logger)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

// NewBasicService returns a naïve, stateless implementation of Service.
func NewBasicService(logger log.Logger) Service {
	return basicService{
		logger: logger,
	}
}

// basicService is our implementation of Service.
type basicService struct {
	logger log.Logger
}

func (s basicService) WorkflowGet(_ context.Context, id string) (*pb.WorkflowData, error) {
	path := "workflows/" + id + ".json"
	jsonBlob, err := dist.Asset(path)
	if err != nil {
		return nil, errors.Wrapf(ErrNotFound, "Cannot read file %s", path)
	}
	wfd := &pb.WorkflowData{}
	err = json.Unmarshal(jsonBlob, wfd)
	if err != nil {
		return nil, err
	}
	return wfd, nil
}
