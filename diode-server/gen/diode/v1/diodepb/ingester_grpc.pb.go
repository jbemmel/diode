// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: diode/v1/ingester.proto

package diodepb

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

const (
	IngesterService_Ingest_FullMethodName = "/diode.v1.IngesterService/Ingest"
)

// IngesterServiceClient is the client API for IngesterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IngesterServiceClient interface {
	// Ingests data into the system
	Ingest(ctx context.Context, in *IngestRequest, opts ...grpc.CallOption) (*IngestResponse, error)
}

type ingesterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIngesterServiceClient(cc grpc.ClientConnInterface) IngesterServiceClient {
	return &ingesterServiceClient{cc}
}

func (c *ingesterServiceClient) Ingest(ctx context.Context, in *IngestRequest, opts ...grpc.CallOption) (*IngestResponse, error) {
	out := new(IngestResponse)
	err := c.cc.Invoke(ctx, IngesterService_Ingest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngesterServiceServer is the server API for IngesterService service.
// All implementations must embed UnimplementedIngesterServiceServer
// for forward compatibility
type IngesterServiceServer interface {
	// Ingests data into the system
	Ingest(context.Context, *IngestRequest) (*IngestResponse, error)
	mustEmbedUnimplementedIngesterServiceServer()
}

// UnimplementedIngesterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIngesterServiceServer struct {
}

func (UnimplementedIngesterServiceServer) Ingest(context.Context, *IngestRequest) (*IngestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ingest not implemented")
}
func (UnimplementedIngesterServiceServer) mustEmbedUnimplementedIngesterServiceServer() {}

// UnsafeIngesterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IngesterServiceServer will
// result in compilation errors.
type UnsafeIngesterServiceServer interface {
	mustEmbedUnimplementedIngesterServiceServer()
}

func RegisterIngesterServiceServer(s grpc.ServiceRegistrar, srv IngesterServiceServer) {
	s.RegisterService(&IngesterService_ServiceDesc, srv)
}

func _IngesterService_Ingest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngesterServiceServer).Ingest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IngesterService_Ingest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngesterServiceServer).Ingest(ctx, req.(*IngestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IngesterService_ServiceDesc is the grpc.ServiceDesc for IngesterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IngesterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "diode.v1.IngesterService",
	HandlerType: (*IngesterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ingest",
			Handler:    _IngesterService_Ingest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "diode/v1/ingester.proto",
}
