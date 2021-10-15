package endpoints

import "github.com/rapido-liebre/gpsDataCollector/internal"

type AddCoordinatesRequest struct {
	Coordinates *internal.Coordinates `json:"coordinates"`
}

type AddCoordinatesResponse struct {
	InsertedId string `json:"insertedId"`
	Err      string `json:"err,omitempty"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
