// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: ggezchain/acl/tx.proto

package acl

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
	Msg_UpdateParams_FullMethodName     = "/ggezchain.acl.Msg/UpdateParams"
	Msg_Init_FullMethodName             = "/ggezchain.acl.Msg/Init"
	Msg_UpdateSuperAdmin_FullMethodName = "/ggezchain.acl.Msg/UpdateSuperAdmin"
	Msg_AddAdmin_FullMethodName         = "/ggezchain.acl.Msg/AddAdmin"
	Msg_DeleteAdmin_FullMethodName      = "/ggezchain.acl.Msg/DeleteAdmin"
	Msg_AddAuthority_FullMethodName     = "/ggezchain.acl.Msg/AddAuthority"
	Msg_UpdateAuthority_FullMethodName  = "/ggezchain.acl.Msg/UpdateAuthority"
	Msg_DeleteAuthority_FullMethodName  = "/ggezchain.acl.Msg/DeleteAuthority"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	Init(ctx context.Context, in *MsgInit, opts ...grpc.CallOption) (*MsgInitResponse, error)
	UpdateSuperAdmin(ctx context.Context, in *MsgUpdateSuperAdmin, opts ...grpc.CallOption) (*MsgUpdateSuperAdminResponse, error)
	AddAdmin(ctx context.Context, in *MsgAddAdmin, opts ...grpc.CallOption) (*MsgAddAdminResponse, error)
	DeleteAdmin(ctx context.Context, in *MsgDeleteAdmin, opts ...grpc.CallOption) (*MsgDeleteAdminResponse, error)
	AddAuthority(ctx context.Context, in *MsgAddAuthority, opts ...grpc.CallOption) (*MsgAddAuthorityResponse, error)
	UpdateAuthority(ctx context.Context, in *MsgUpdateAuthority, opts ...grpc.CallOption) (*MsgUpdateAuthorityResponse, error)
	DeleteAuthority(ctx context.Context, in *MsgDeleteAuthority, opts ...grpc.CallOption) (*MsgDeleteAuthorityResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Init(ctx context.Context, in *MsgInit, opts ...grpc.CallOption) (*MsgInitResponse, error) {
	out := new(MsgInitResponse)
	err := c.cc.Invoke(ctx, Msg_Init_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateSuperAdmin(ctx context.Context, in *MsgUpdateSuperAdmin, opts ...grpc.CallOption) (*MsgUpdateSuperAdminResponse, error) {
	out := new(MsgUpdateSuperAdminResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateSuperAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddAdmin(ctx context.Context, in *MsgAddAdmin, opts ...grpc.CallOption) (*MsgAddAdminResponse, error) {
	out := new(MsgAddAdminResponse)
	err := c.cc.Invoke(ctx, Msg_AddAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteAdmin(ctx context.Context, in *MsgDeleteAdmin, opts ...grpc.CallOption) (*MsgDeleteAdminResponse, error) {
	out := new(MsgDeleteAdminResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddAuthority(ctx context.Context, in *MsgAddAuthority, opts ...grpc.CallOption) (*MsgAddAuthorityResponse, error) {
	out := new(MsgAddAuthorityResponse)
	err := c.cc.Invoke(ctx, Msg_AddAuthority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateAuthority(ctx context.Context, in *MsgUpdateAuthority, opts ...grpc.CallOption) (*MsgUpdateAuthorityResponse, error) {
	out := new(MsgUpdateAuthorityResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateAuthority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteAuthority(ctx context.Context, in *MsgDeleteAuthority, opts ...grpc.CallOption) (*MsgDeleteAuthorityResponse, error) {
	out := new(MsgDeleteAuthorityResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteAuthority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	Init(context.Context, *MsgInit) (*MsgInitResponse, error)
	UpdateSuperAdmin(context.Context, *MsgUpdateSuperAdmin) (*MsgUpdateSuperAdminResponse, error)
	AddAdmin(context.Context, *MsgAddAdmin) (*MsgAddAdminResponse, error)
	DeleteAdmin(context.Context, *MsgDeleteAdmin) (*MsgDeleteAdminResponse, error)
	AddAuthority(context.Context, *MsgAddAuthority) (*MsgAddAuthorityResponse, error)
	UpdateAuthority(context.Context, *MsgUpdateAuthority) (*MsgUpdateAuthorityResponse, error)
	DeleteAuthority(context.Context, *MsgDeleteAuthority) (*MsgDeleteAuthorityResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) Init(context.Context, *MsgInit) (*MsgInitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedMsgServer) UpdateSuperAdmin(context.Context, *MsgUpdateSuperAdmin) (*MsgUpdateSuperAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSuperAdmin not implemented")
}
func (UnimplementedMsgServer) AddAdmin(context.Context, *MsgAddAdmin) (*MsgAddAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAdmin not implemented")
}
func (UnimplementedMsgServer) DeleteAdmin(context.Context, *MsgDeleteAdmin) (*MsgDeleteAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAdmin not implemented")
}
func (UnimplementedMsgServer) AddAuthority(context.Context, *MsgAddAuthority) (*MsgAddAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAuthority not implemented")
}
func (UnimplementedMsgServer) UpdateAuthority(context.Context, *MsgUpdateAuthority) (*MsgUpdateAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthority not implemented")
}
func (UnimplementedMsgServer) DeleteAuthority(context.Context, *MsgDeleteAuthority) (*MsgDeleteAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthority not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Init(ctx, req.(*MsgInit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateSuperAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateSuperAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateSuperAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateSuperAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateSuperAdmin(ctx, req.(*MsgUpdateSuperAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddAdmin(ctx, req.(*MsgAddAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteAdmin(ctx, req.(*MsgDeleteAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddAuthority)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddAuthority(ctx, req.(*MsgAddAuthority))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateAuthority)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateAuthority(ctx, req.(*MsgUpdateAuthority))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteAuthority)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteAuthority(ctx, req.(*MsgDeleteAuthority))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ggezchain.acl.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _Msg_Init_Handler,
		},
		{
			MethodName: "UpdateSuperAdmin",
			Handler:    _Msg_UpdateSuperAdmin_Handler,
		},
		{
			MethodName: "AddAdmin",
			Handler:    _Msg_AddAdmin_Handler,
		},
		{
			MethodName: "DeleteAdmin",
			Handler:    _Msg_DeleteAdmin_Handler,
		},
		{
			MethodName: "AddAuthority",
			Handler:    _Msg_AddAuthority_Handler,
		},
		{
			MethodName: "UpdateAuthority",
			Handler:    _Msg_UpdateAuthority_Handler,
		},
		{
			MethodName: "DeleteAuthority",
			Handler:    _Msg_DeleteAuthority_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ggezchain/acl/tx.proto",
}
