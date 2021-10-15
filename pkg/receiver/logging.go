package receiver

import (
	"context"
	"github.com/rapido-liebre/gpsDataCollector/internal"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func NewloggingService(logger log.Logger, next Service) Service { return &loggingMiddleware{} }

func (mw loggingMiddleware) AddCoordinates(ctx context.Context, location *internal.Coordinates) (insertedId string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "AddCoordinates",
			"location", location,
			"insertedId", insertedId,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	insertedId, err = mw.next.AddCoordinates(ctx, location)
	return
}

func (mw loggingMiddleware) ServiceStatus(ctx context.Context) (status int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "ServiceStatus",
			"status", status,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	status, err = mw.next.ServiceStatus(ctx)
	return
}
