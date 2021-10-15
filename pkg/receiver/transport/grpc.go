package transport

import (
	"context"
	pb "github.com/rapido-liebre/gpsDataCollector/api/v1/pb"

	"github.com/rapido-liebre/gpsDataCollector/internal"
	"github.com/rapido-liebre/gpsDataCollector/pkg/receiver/endpoints"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	addCoordinates grpctransport.Handler
	serviceStatus grpctransport.Handler
	//fix for https://youtrack.jetbrains.com/issue/GO-11357
	pb.UnimplementedGpsDataCollectorServer
}

//func (g *grpcServer) AddCoordinates(ctx context.Context, a *interface{}) (*interface{}, error) {

func (g *grpcServer) AddCoordinates(ctx context.Context, r *pb.AddCoordinatesRequest) (*pb.AddCoordinatesReply, error) {
	//panic("implement me")
	_, rep, err := g.addCoordinates.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddCoordinatesReply), nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, r *pb.ServiceStatusRequest) (*pb.ServiceStatusReply, error) {
	//panic("implement me")
	_, rep, err := g.serviceStatus.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ServiceStatusReply), nil
}

func NewGRPCServer(ep endpoints.Endpoints) pb.GpsDataCollectorServer {
	return &grpcServer{
		addCoordinates: grpctransport.NewServer(
			ep.AddCoordinatesEndpoint,
			decodeGRPCAddCoordinatesRequest,
			decodeGRPCAddCoordinatesResponse,
		),
		serviceStatus: grpctransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeGRPCServiceStatusRequest,
			decodeGRPCServiceStatusResponse,
		),
	}
}

func decodeGRPCAddCoordinatesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddCoordinatesRequest)
	point := &internal.Coordinates{
		Latitude: req.Coordinates.Latitude,
		Longitude: req.Coordinates.Longitude,
	}
	return endpoints.AddCoordinatesRequest{Coordinates: point}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}

func decodeGRPCAddCoordinatesResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.AddCoordinatesReply)
	return endpoints.AddCoordinatesResponse{InsertedId: reply.InsertedId, Err: reply.Err}, nil
}

func decodeGRPCServiceStatusResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.ServiceStatusReply)
	return endpoints.ServiceStatusResponse{Code: int(reply.Code), Err: reply.Err}, nil
}
