// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cliente

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SolicitudClient is the client API for Solicitud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SolicitudClient interface {
	ShowOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Sample, error)
	MakeOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Confirmation, error)
	GetStatus(ctx context.Context, in *CodigoSeguimiento, opts ...grpc.CallOption) (*Status, error)
}

type solicitudClient struct {
	cc grpc.ClientConnInterface
}

func NewSolicitudClient(cc grpc.ClientConnInterface) SolicitudClient {
	return &solicitudClient{cc}
}

func (c *solicitudClient) ShowOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Sample, error) {
	out := new(Sample)
	err := c.cc.Invoke(ctx, "/cliente.Solicitud/ShowOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solicitudClient) MakeOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/cliente.Solicitud/MakeOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solicitudClient) GetStatus(ctx context.Context, in *CodigoSeguimiento, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/cliente.Solicitud/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolicitudServer is the server API for Solicitud service.
// All implementations must embed UnimplementedSolicitudServer
// for forward compatibility
type SolicitudServer interface {
	ShowOrder(context.Context, *Order) (*Sample, error)
	MakeOrder(context.Context, *Order) (*Confirmation, error)
	GetStatus(context.Context, *CodigoSeguimiento) (*Status, error)
	mustEmbedUnimplementedSolicitudServer()
}

// UnimplementedSolicitudServer must be embedded to have forward compatible implementations.
type UnimplementedSolicitudServer struct {
}

func (UnimplementedSolicitudServer) ShowOrder(context.Context, *Order) (*Sample, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowOrder not implemented")
}
func (UnimplementedSolicitudServer) MakeOrder(context.Context, *Order) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeOrder not implemented")
}
func (UnimplementedSolicitudServer) GetStatus(context.Context, *CodigoSeguimiento) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedSolicitudServer) mustEmbedUnimplementedSolicitudServer() {}

// UnsafeSolicitudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SolicitudServer will
// result in compilation errors.
type UnsafeSolicitudServer interface {
	mustEmbedUnimplementedSolicitudServer()
}

func RegisterSolicitudServer(s *grpc.Server, srv SolicitudServer) {
	s.RegisterService(&_Solicitud_serviceDesc, srv)
}

func _Solicitud_ShowOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolicitudServer).ShowOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cliente.Solicitud/ShowOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolicitudServer).ShowOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Solicitud_MakeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolicitudServer).MakeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cliente.Solicitud/MakeOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolicitudServer).MakeOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Solicitud_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodigoSeguimiento)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolicitudServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cliente.Solicitud/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolicitudServer).GetStatus(ctx, req.(*CodigoSeguimiento))
	}
	return interceptor(ctx, in, info, handler)
}

var _Solicitud_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cliente.Solicitud",
	HandlerType: (*SolicitudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowOrder",
			Handler:    _Solicitud_ShowOrder_Handler,
		},
		{
			MethodName: "MakeOrder",
			Handler:    _Solicitud_MakeOrder_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Solicitud_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cliente.proto",
}
