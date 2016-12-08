package service

import (
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/artefactual-labs/archivematica-workflow/pkg/service/pb"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

// LoggingMiddleware takes a logger as a dependency and returns a
// ServiceMiddleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func (mw loggingMiddleware) WorkflowGet(ctx context.Context, id string) (_ *pb.WorkflowData, err error) {
	defer func() {
		mw.logger.Log("method", "WorkflowGet", "id", id, "err", err)
	}()
	return mw.next.WorkflowGet(ctx, id)
}
