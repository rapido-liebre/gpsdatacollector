package endpoints

import (
	"context"
	"errors"
	"os"

	"github.com/rapido-liebre/gpsDataCollector/internal"
	"github.com/rapido-liebre/gpsDataCollector/pkg/receiver"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Endpoints struct {
	AddCoordinatesEndpoint endpoint.Endpoint
	ServiceStatusEndpoint  endpoint.Endpoint
}

func NewEndpointSet(svc receiver.Service) Endpoints {
	return Endpoints{
		AddCoordinatesEndpoint: MakeAddCoordinatesEndpoint(svc),
		ServiceStatusEndpoint:  MakeServiceStatusEndpoint(svc),
	}
}

func MakeAddCoordinatesEndpoint(svc receiver.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddCoordinatesRequest)
		insertedId, err := svc.AddCoordinates(ctx, req.Coordinates)
		if err != nil {
			return AddCoordinatesResponse{InsertedId: insertedId, Err: err.Error()}, nil
		}
		return AddCoordinatesResponse{InsertedId: insertedId, Err: ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc receiver.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}

func (s *Endpoints) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})
	svcStatusResp := resp.(ServiceStatusResponse)
	if err != nil {
		return svcStatusResp.Code, err
	}
	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Code, nil
}

func (s *Endpoints) AddCoordinates(ctx context.Context, doc *internal.Coordinates) (string, error) {
	resp, err := s.AddCoordinatesEndpoint(ctx, AddCoordinatesRequest{Coordinates: doc})
	if err != nil {
		return "", err
	}
	adResp := resp.(AddCoordinatesResponse)
	if adResp.Err != "" {
		return "", errors.New(adResp.Err)
	}
	return adResp.InsertedId, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts_eps", log.DefaultTimestampUTC)
}
