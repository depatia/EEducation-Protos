// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: user.proto

package user

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
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	SetPermissionLevel(ctx context.Context, in *SetPermissionLevelRequest, opts ...grpc.CallOption) (*SetPermissionLevelResponse, error)
	GetPermissionLevel(ctx context.Context, in *GetPermissionLevelRequest, opts ...grpc.CallOption) (*GetPermissionLevelResponse, error)
	FillUserProfile(ctx context.Context, in *FillUserProfileRequest, opts ...grpc.CallOption) (*FillUserProfileResponse, error)
	ChangeUserStatus(ctx context.Context, in *ChangeUserStatusRequest, opts ...grpc.CallOption) (*ChangeUserStatusResponse, error)
	IsUserActive(ctx context.Context, in *IsUserActiveRequest, opts ...grpc.CallOption) (*IsUserActiveResponse, error)
	GetStudentsByClassname(ctx context.Context, in *GetStudentsByClassnameRequest, opts ...grpc.CallOption) (*GetStudentsByClassnameResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	out := new(UpdatePasswordResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetPermissionLevel(ctx context.Context, in *SetPermissionLevelRequest, opts ...grpc.CallOption) (*SetPermissionLevelResponse, error) {
	out := new(SetPermissionLevelResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/SetPermissionLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetPermissionLevel(ctx context.Context, in *GetPermissionLevelRequest, opts ...grpc.CallOption) (*GetPermissionLevelResponse, error) {
	out := new(GetPermissionLevelResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetPermissionLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FillUserProfile(ctx context.Context, in *FillUserProfileRequest, opts ...grpc.CallOption) (*FillUserProfileResponse, error) {
	out := new(FillUserProfileResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/FillUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeUserStatus(ctx context.Context, in *ChangeUserStatusRequest, opts ...grpc.CallOption) (*ChangeUserStatusResponse, error) {
	out := new(ChangeUserStatusResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ChangeUserStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IsUserActive(ctx context.Context, in *IsUserActiveRequest, opts ...grpc.CallOption) (*IsUserActiveResponse, error) {
	out := new(IsUserActiveResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/IsUserActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetStudentsByClassname(ctx context.Context, in *GetStudentsByClassnameRequest, opts ...grpc.CallOption) (*GetStudentsByClassnameResponse, error) {
	out := new(GetStudentsByClassnameResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetStudentsByClassname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	SetPermissionLevel(context.Context, *SetPermissionLevelRequest) (*SetPermissionLevelResponse, error)
	GetPermissionLevel(context.Context, *GetPermissionLevelRequest) (*GetPermissionLevelResponse, error)
	FillUserProfile(context.Context, *FillUserProfileRequest) (*FillUserProfileResponse, error)
	ChangeUserStatus(context.Context, *ChangeUserStatusRequest) (*ChangeUserStatusResponse, error)
	IsUserActive(context.Context, *IsUserActiveRequest) (*IsUserActiveResponse, error)
	GetStudentsByClassname(context.Context, *GetStudentsByClassnameRequest) (*GetStudentsByClassnameResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserServiceServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedUserServiceServer) SetPermissionLevel(context.Context, *SetPermissionLevelRequest) (*SetPermissionLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPermissionLevel not implemented")
}
func (UnimplementedUserServiceServer) GetPermissionLevel(context.Context, *GetPermissionLevelRequest) (*GetPermissionLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissionLevel not implemented")
}
func (UnimplementedUserServiceServer) FillUserProfile(context.Context, *FillUserProfileRequest) (*FillUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FillUserProfile not implemented")
}
func (UnimplementedUserServiceServer) ChangeUserStatus(context.Context, *ChangeUserStatusRequest) (*ChangeUserStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserStatus not implemented")
}
func (UnimplementedUserServiceServer) IsUserActive(context.Context, *IsUserActiveRequest) (*IsUserActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUserActive not implemented")
}
func (UnimplementedUserServiceServer) GetStudentsByClassname(context.Context, *GetStudentsByClassnameRequest) (*GetStudentsByClassnameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentsByClassname not implemented")
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

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetPermissionLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPermissionLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetPermissionLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/SetPermissionLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetPermissionLevel(ctx, req.(*SetPermissionLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetPermissionLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetPermissionLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetPermissionLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetPermissionLevel(ctx, req.(*GetPermissionLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FillUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FillUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FillUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FillUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FillUserProfile(ctx, req.(*FillUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ChangeUserStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeUserStatus(ctx, req.(*ChangeUserStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IsUserActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsUserActiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IsUserActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/IsUserActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IsUserActive(ctx, req.(*IsUserActiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetStudentsByClassname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentsByClassnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetStudentsByClassname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetStudentsByClassname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetStudentsByClassname(ctx, req.(*GetStudentsByClassnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserService_UpdatePassword_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _UserService_Validate_Handler,
		},
		{
			MethodName: "SetPermissionLevel",
			Handler:    _UserService_SetPermissionLevel_Handler,
		},
		{
			MethodName: "GetPermissionLevel",
			Handler:    _UserService_GetPermissionLevel_Handler,
		},
		{
			MethodName: "FillUserProfile",
			Handler:    _UserService_FillUserProfile_Handler,
		},
		{
			MethodName: "ChangeUserStatus",
			Handler:    _UserService_ChangeUserStatus_Handler,
		},
		{
			MethodName: "IsUserActive",
			Handler:    _UserService_IsUserActive_Handler,
		},
		{
			MethodName: "GetStudentsByClassname",
			Handler:    _UserService_GetStudentsByClassname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
