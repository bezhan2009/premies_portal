// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: card_prices.proto

package card_prices

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CardPricesService_UploadCardPricesData_FullMethodName = "/CardPricesService/UploadCardPricesData"
)

// CardPricesServiceClient is the client API for CardPricesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardPricesServiceClient interface {
	UploadCardPricesData(ctx context.Context, in *CardPricesUploadRequest, opts ...grpc.CallOption) (*CardPricesUploadResponse, error)
}

type cardPricesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardPricesServiceClient(cc grpc.ClientConnInterface) CardPricesServiceClient {
	return &cardPricesServiceClient{cc}
}

func (c *cardPricesServiceClient) UploadCardPricesData(ctx context.Context, in *CardPricesUploadRequest, opts ...grpc.CallOption) (*CardPricesUploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CardPricesUploadResponse)
	err := c.cc.Invoke(ctx, CardPricesService_UploadCardPricesData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardPricesServiceServer is the server API for CardPricesService service.
// All implementations must embed UnimplementedCardPricesServiceServer
// for forward compatibility.
type CardPricesServiceServer interface {
	UploadCardPricesData(context.Context, *CardPricesUploadRequest) (*CardPricesUploadResponse, error)
	mustEmbedUnimplementedCardPricesServiceServer()
}

// UnimplementedCardPricesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCardPricesServiceServer struct{}

func (UnimplementedCardPricesServiceServer) UploadCardPricesData(context.Context, *CardPricesUploadRequest) (*CardPricesUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadCardPricesData not implemented")
}
func (UnimplementedCardPricesServiceServer) mustEmbedUnimplementedCardPricesServiceServer() {}
func (UnimplementedCardPricesServiceServer) testEmbeddedByValue()                           {}

// UnsafeCardPricesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardPricesServiceServer will
// result in compilation errors.
type UnsafeCardPricesServiceServer interface {
	mustEmbedUnimplementedCardPricesServiceServer()
}

func RegisterCardPricesServiceServer(s grpc.ServiceRegistrar, srv CardPricesServiceServer) {
	// If the following call pancis, it indicates UnimplementedCardPricesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CardPricesService_ServiceDesc, srv)
}

func _CardPricesService_UploadCardPricesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardPricesUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardPricesServiceServer).UploadCardPricesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardPricesService_UploadCardPricesData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardPricesServiceServer).UploadCardPricesData(ctx, req.(*CardPricesUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardPricesService_ServiceDesc is the grpc.ServiceDesc for CardPricesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardPricesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CardPricesService",
	HandlerType: (*CardPricesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadCardPricesData",
			Handler:    _CardPricesService_UploadCardPricesData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "card_prices.proto",
}
