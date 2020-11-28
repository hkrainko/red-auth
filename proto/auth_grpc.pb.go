// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	GetAuthUrl(ctx context.Context, in *GetAuthUrlRequest, opts ...grpc.CallOption) (*GetAuthUrlResponse, error)
	CallBack(ctx context.Context, in *CallBackRequest, opts ...grpc.CallOption) (*CallBackResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) GetAuthUrl(ctx context.Context, in *GetAuthUrlRequest, opts ...grpc.CallOption) (*GetAuthUrlResponse, error) {
	out := new(GetAuthUrlResponse)
	err := c.cc.Invoke(ctx, "/AuthService/GetAuthUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CallBack(ctx context.Context, in *CallBackRequest, opts ...grpc.CallOption) (*CallBackResponse, error) {
	out := new(CallBackResponse)
	err := c.cc.Invoke(ctx, "/AuthService/CallBack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	GetAuthUrl(context.Context, *GetAuthUrlRequest) (*GetAuthUrlResponse, error)
	CallBack(context.Context, *CallBackRequest) (*CallBackResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) GetAuthUrl(context.Context, *GetAuthUrlRequest) (*GetAuthUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthUrl not implemented")
}
func (UnimplementedAuthServiceServer) CallBack(context.Context, *CallBackRequest) (*CallBackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallBack not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_GetAuthUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAuthUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/GetAuthUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAuthUrl(ctx, req.(*GetAuthUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CallBack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallBackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CallBack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/CallBack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CallBack(ctx, req.(*CallBackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthUrl",
			Handler:    _AuthService_GetAuthUrl_Handler,
		},
		{
			MethodName: "CallBack",
			Handler:    _AuthService_CallBack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}
