package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rapido-liebre/gpsDataCollector/config"
	"github.com/rapido-liebre/gpsDataCollector/internal/database"
	"gorm.io/gorm"

	pb "github.com/rapido-liebre/gpsDataCollector/api/v1/pb"

	"github.com/rapido-liebre/gpsDataCollector/pkg/receiver"
	"github.com/rapido-liebre/gpsDataCollector/pkg/receiver/endpoints"
	"github.com/rapido-liebre/gpsDataCollector/pkg/receiver/transport"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
)

var (
	logger log.Logger
	db     *gorm.DB
	cfg    *config.Config
)

func main() {
	var (
		repo        = database.NewRepo(db, logger)
		service     = receiver.NewService(repo, logger)
		eps         = endpoints.NewEndpointSet(service)
		httpHandler = transport.NewHTTPHandler(eps)
		grpcServer  = transport.NewGRPCServer(eps)
	)
	service = receiver.NewloggingService(logger, service)

	var g group.Group
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", cfg.Ports.HttpAddr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(
			func() error {
				logger.Log("transport", "HTTP", "addr", cfg.Ports.HttpAddr)
				return http.Serve(httpListener, httpHandler)
			}, func(error) {
				httpListener.Close()
			},
		)
	}
	{
		// The gRPC listener mounts the Go kit gRPC server we created.
		grpcListener, err := net.Listen("tcp", cfg.Ports.GrpcAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(
			func() error {
				logger.Log("transport", "gRPC", "addr", cfg.Ports.GrpcAddr)
				// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
				// the here demonstrated zipkin tracing middleware.
				baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
				pb.RegisterGpsDataCollectorServer(baseServer, grpcServer)
				return baseServer.Serve(grpcListener)
			}, func(error) {
				grpcListener.Close()
			},
		)
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(
			func() error {
				c := make(chan os.Signal, 1)
				signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
				select {
				case sig := <-c:
					return fmt.Errorf("received signal %s", sig)
				case <-cancelInterrupt:
					return nil
				}
			}, func(error) {
				close(cancelInterrupt)
			},
		)
	}
	logger.Log("exit", g.Run())
}

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	cfg = config.NewConfig(logger)
	if cfg == nil {
		log.With(logger, "ERROR: reading configuration failed")
		os.Exit(2)
	}

	dbConn, err := database.Init(
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Database.User, cfg.Database.Password,
	)
	dbSQL, ok := dbConn.DB()
	if ok != nil {
		defer func() {
			err := dbSQL.Close()
			if err != nil {
				logger.Log("ERROR: failed to close the database connection ", err.Error())
			}
		}()
	}
	if err != nil {
		logger.Log(fmt.Sprintf("FATAL: failed to load db with error: %s", err.Error()))
		os.Exit(3)
	}
	db = dbConn
}
