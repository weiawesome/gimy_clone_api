// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: proto/ad_service.proto

package ad_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Advertisement_GetAd_FullMethodName  = "/advertisement.Advertisement/GetAd"
	Advertisement_SaveAd_FullMethodName = "/advertisement.Advertisement/SaveAd"
)

// AdvertisementClient is the client API for Advertisement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdvertisementClient interface {
	GetAd(ctx context.Context, in *GetAdvertisementRequest, opts ...grpc.CallOption) (*GetAdvertisementReply, error)
	SaveAd(ctx context.Context, in *SaveAdvertisementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type advertisementClient struct {
	cc grpc.ClientConnInterface
}

func NewAdvertisementClient(cc grpc.ClientConnInterface) AdvertisementClient {
	return &advertisementClient{cc}
}

func (c *advertisementClient) GetAd(ctx context.Context, in *GetAdvertisementRequest, opts ...grpc.CallOption) (*GetAdvertisementReply, error) {
	out := new(GetAdvertisementReply)
	err := c.cc.Invoke(ctx, Advertisement_GetAd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertisementClient) SaveAd(ctx context.Context, in *SaveAdvertisementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Advertisement_SaveAd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvertisementServer is the server API for Advertisement service.
// All implementations must embed UnimplementedAdvertisementServer
// for forward compatibility
type AdvertisementServer interface {
	GetAd(context.Context, *GetAdvertisementRequest) (*GetAdvertisementReply, error)
	SaveAd(context.Context, *SaveAdvertisementRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAdvertisementServer()
}

// UnimplementedAdvertisementServer must be embedded to have forward compatible implementations.
type UnimplementedAdvertisementServer struct {
}

func (UnimplementedAdvertisementServer) GetAd(context.Context, *GetAdvertisementRequest) (*GetAdvertisementReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAd not implemented")
}
func (UnimplementedAdvertisementServer) SaveAd(context.Context, *SaveAdvertisementRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveAd not implemented")
}
func (UnimplementedAdvertisementServer) mustEmbedUnimplementedAdvertisementServer() {}

// UnsafeAdvertisementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdvertisementServer will
// result in compilation errors.
type UnsafeAdvertisementServer interface {
	mustEmbedUnimplementedAdvertisementServer()
}

func RegisterAdvertisementServer(s grpc.ServiceRegistrar, srv AdvertisementServer) {
	s.RegisterService(&Advertisement_ServiceDesc, srv)
}

func _Advertisement_GetAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdvertisementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServer).GetAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Advertisement_GetAd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServer).GetAd(ctx, req.(*GetAdvertisementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Advertisement_SaveAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveAdvertisementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServer).SaveAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Advertisement_SaveAd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServer).SaveAd(ctx, req.(*SaveAdvertisementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Advertisement_ServiceDesc is the grpc.ServiceDesc for Advertisement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Advertisement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "advertisement.Advertisement",
	HandlerType: (*AdvertisementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAd",
			Handler:    _Advertisement_GetAd_Handler,
		},
		{
			MethodName: "SaveAd",
			Handler:    _Advertisement_SaveAd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ad_service.proto",
}