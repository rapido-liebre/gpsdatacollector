package receiver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/lithammer/shortuuid"
	"github.com/rapido-liebre/gpsDataCollector/internal"
)

type receiverService struct {
	repository Repository
	logger     log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	log.With(logger, "ts receiver NewService() ", log.DefaultTimestampUTC)
	return &receiverService{
		repository: repo,
		logger:     logger}
}

func (w *receiverService) AddCoordinates(ctx context.Context, point *internal.Coordinates) (string, error) {
	// authenticate sender, return deviceId
	deviceId := shortuuid.New()

	// add the coordinates entry in the database by calling the database service
	// return error if the point is invalid and/or the database invalid entry error
	data := internal.GpsData{DeviceId: deviceId, Point: *point}
	err := w.repository.AddCoordinates(ctx, &data)
	if err != nil {
		w.logger.Log("receiver", "AddCoordinates", "during insert err:", err)
		//os.Exit(1)
	}
	w.logger.Log("receiver", "Added new coordinates successfully", data.Point)

	return fmt.Sprintf("Success. ID: %s", data.Id), nil
}

func (w *receiverService) ServiceStatus(_ context.Context) (int, error) {
	w.logger.Log("receiver", "Checking the Service health...")
	return http.StatusOK, nil
}
