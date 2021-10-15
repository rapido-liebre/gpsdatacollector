package receiver

import (
	"context"

	"github.com/rapido-liebre/gpsDataCollector/internal"
)

type Service interface {
	AddCoordinates(ctx context.Context, point *internal.Coordinates) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}

type Repository interface {
	AddCoordinates(ctx context.Context, data *internal.GpsData) error
}