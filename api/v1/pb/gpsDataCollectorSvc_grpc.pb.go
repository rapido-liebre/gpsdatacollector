// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GpsDataCollectorClient is the client API for GpsDataCollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GpsDataCollectorClient interface {
	AddCoordinates(ctx context.Context, in *AddCoordinatesRequest, opts ...grpc.CallOption) (*AddCoordinatesReply, error)
	ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error)
}

type gpsDataCollectorClient struct {
	cc grpc.ClientConnInterface
}

func NewGpsDataCollectorClient(cc grpc.ClientConnInterface) GpsDataCollectorClient {
	return &gpsDataCollectorClient{cc}
}

func (c *gpsDataCollectorClient) AddCoordinates(ctx context.Context, in *AddCoordinatesRequest, opts ...grpc.CallOption) (*AddCoordinatesReply, error) {
	out := new(AddCoordinatesReply)
	err := c.cc.Invoke(ctx, "/GpsDataCollector/AddCoordinates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gpsDataCollectorClient) ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error) {
	out := new(ServiceStatusReply)
	err := c.cc.Invoke(ctx, "/GpsDataCollector/ServiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GpsDataCollectorServer is the server API for GpsDataCollector service.
// All implementations must embed UnimplementedGpsDataCollectorServer
// for forward compatibility
type GpsDataCollectorServer interface {
	AddCoordinates(context.Context, *AddCoordinatesRequest) (*AddCoordinatesReply, error)
	ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error)
	mustEmbedUnimplementedGpsDataCollectorServer()
}

// UnimplementedGpsDataCollectorServer must be embedded to have forward compatible implementations.
type UnimplementedGpsDataCollectorServer struct {
}

func (UnimplementedGpsDataCollectorServer) AddCoordinates(context.Context, *AddCoordinatesRequest) (*AddCoordinatesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCoordinates not implemented")
}
func (UnimplementedGpsDataCollectorServer) ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceStatus not implemented")
}
func (UnimplementedGpsDataCollectorServer) mustEmbedUnimplementedGpsDataCollectorServer() {}

// UnsafeGpsDataCollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GpsDataCollectorServer will
// result in compilation errors.
type UnsafeGpsDataCollectorServer interface {
	mustEmbedUnimplementedGpsDataCollectorServer()
}

func RegisterGpsDataCollectorServer(s grpc.ServiceRegistrar, srv GpsDataCollectorServer) {
	s.RegisterService(&GpsDataCollector_ServiceDesc, srv)
}

func _GpsDataCollector_AddCoordinates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCoordinatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GpsDataCollectorServer).AddCoordinates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GpsDataCollector/AddCoordinates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GpsDataCollectorServer).AddCoordinates(ctx, req.(*AddCoordinatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GpsDataCollector_ServiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GpsDataCollectorServer).ServiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GpsDataCollector/ServiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GpsDataCollectorServer).ServiceStatus(ctx, req.(*ServiceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GpsDataCollector_ServiceDesc is the grpc.ServiceDesc for GpsDataCollector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GpsDataCollector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GpsDataCollector",
	HandlerType: (*GpsDataCollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCoordinates",
			Handler:    _GpsDataCollector_AddCoordinates_Handler,
		},
		{
			MethodName: "ServiceStatus",
			Handler:    _GpsDataCollector_ServiceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/pb/gpsDataCollectorSvc.proto",
}
