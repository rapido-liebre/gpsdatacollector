package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	auth "github.com/go-kit/kit/auth/basic"

	"github.com/rapido-liebre/gpsDataCollector/internal/util"
	"github.com/rapido-liebre/gpsDataCollector/pkg/receiver/endpoints"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(ep endpoints.Endpoints) http.Handler {
	m := mux.NewRouter()
	//m.Use(basicAuthMiddleware)

	m.Methods("POST").Path("/addCoords").Handler(
		httptransport.NewServer(
			auth.AuthMiddleware(
				os.Getenv("AUTH_USER"), os.Getenv("AUTH_PASS"), "Gps Data Collector",
			)(ep.AddCoordinatesEndpoint),
			decodeHTTPAddCoordinatesRequest,
			encodeResponse,
			httptransport.ServerBefore(httptransport.PopulateRequestContext),
		),
	)

	m.Methods("GET").Path("/health").Handler(
		httptransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeHTTPServiceStatusRequest,
			encodeResponse,
		),
	)
	return m
}

func decodeHTTPAddCoordinatesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.AddCoordinatesRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPServiceStatusRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req endpoints.ServiceStatusRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case util.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case util.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"error": err.Error(),
		},
	)
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
