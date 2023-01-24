// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api.proto

package api

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthDto, error)
	GetUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserDto, error)
	DeleteUserById(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*ResponseMessage, error)
	UpdateUserById(ctx context.Context, in *User, opts ...grpc.CallOption) (*ResponseMessage, error)
	GetAllUser(ctx context.Context, in *Null, opts ...grpc.CallOption) (*ResponseUsersDto, error)
	SignIn(ctx context.Context, in *Authentication, opts ...grpc.CallOption) (*AuthDto, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthDto, error) {
	out := new(AuthDto)
	err := c.cc.Invoke(ctx, "/apiproto.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserDto, error) {
	out := new(UserDto)
	err := c.cc.Invoke(ctx, "/apiproto.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUserById(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/apiproto.UserService/DeleteUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserById(ctx context.Context, in *User, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/apiproto.UserService/UpdateUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUser(ctx context.Context, in *Null, opts ...grpc.CallOption) (*ResponseUsersDto, error) {
	out := new(ResponseUsersDto)
	err := c.cc.Invoke(ctx, "/apiproto.UserService/GetAllUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignIn(ctx context.Context, in *Authentication, opts ...grpc.CallOption) (*AuthDto, error) {
	out := new(AuthDto)
	err := c.cc.Invoke(ctx, "/apiproto.UserService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *User) (*AuthDto, error)
	GetUser(context.Context, *UserID) (*UserDto, error)
	DeleteUserById(context.Context, *UserID) (*ResponseMessage, error)
	UpdateUserById(context.Context, *User) (*ResponseMessage, error)
	GetAllUser(context.Context, *Null) (*ResponseUsersDto, error)
	SignIn(context.Context, *Authentication) (*AuthDto, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *User) (*AuthDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *UserID) (*UserDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUserById(context.Context, *UserID) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserById not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserById(context.Context, *User) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserById not implemented")
}
func (UnimplementedUserServiceServer) GetAllUser(context.Context, *Null) (*ResponseUsersDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUser not implemented")
}
func (UnimplementedUserServiceServer) SignIn(context.Context, *Authentication) (*AuthDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiproto.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiproto.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiproto.UserService/DeleteUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUserById(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiproto.UserService/UpdateUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserById(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiproto.UserService/GetAllUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUser(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Authentication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiproto.UserService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignIn(ctx, req.(*Authentication))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiproto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "DeleteUserById",
			Handler:    _UserService_DeleteUserById_Handler,
		},
		{
			MethodName: "UpdateUserById",
			Handler:    _UserService_UpdateUserById_Handler,
		},
		{
			MethodName: "GetAllUser",
			Handler:    _UserService_GetAllUser_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _UserService_SignIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
