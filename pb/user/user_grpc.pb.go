// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.11.2
// source: pb/user/user.proto

package userpb

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

// UserHandlerClient is the client API for UserHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserHandlerClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UserInfo(ctx context.Context, in *ID, opts ...grpc.CallOption) (*UserInfoRes, error)
}

type userHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserHandlerClient(cc grpc.ClientConnInterface) UserHandlerClient {
	return &userHandlerClient{cc}
}

func (c *userHandlerClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/UserHandler/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHandlerClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/UserHandler/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHandlerClient) UserInfo(ctx context.Context, in *ID, opts ...grpc.CallOption) (*UserInfoRes, error) {
	out := new(UserInfoRes)
	err := c.cc.Invoke(ctx, "/UserHandler/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserHandlerServer is the server API for UserHandler service.
// All implementations must embed UnimplementedUserHandlerServer
// for forward compatibility
type UserHandlerServer interface {
	Login(context.Context, *LoginReq) (*LoginRes, error)
	Register(context.Context, *RegisterReq) (*emptypb.Empty, error)
	UserInfo(context.Context, *ID) (*UserInfoRes, error)
	mustEmbedUnimplementedUserHandlerServer()
}

// UnimplementedUserHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedUserHandlerServer struct {
}

func (UnimplementedUserHandlerServer) Login(context.Context, *LoginReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserHandlerServer) Register(context.Context, *RegisterReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserHandlerServer) UserInfo(context.Context, *ID) (*UserInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserHandlerServer) mustEmbedUnimplementedUserHandlerServer() {}

// UnsafeUserHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserHandlerServer will
// result in compilation errors.
type UnsafeUserHandlerServer interface {
	mustEmbedUnimplementedUserHandlerServer()
}

func RegisterUserHandlerServer(s grpc.ServiceRegistrar, srv UserHandlerServer) {
	s.RegisterService(&UserHandler_ServiceDesc, srv)
}

func _UserHandler_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserHandler/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHandler_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserHandler/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHandler_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserHandler/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).UserInfo(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// UserHandler_ServiceDesc is the grpc.ServiceDesc for UserHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserHandler",
	HandlerType: (*UserHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserHandler_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserHandler_Register_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _UserHandler_UserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/user/user.proto",
}
