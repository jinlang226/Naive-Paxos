// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: paxos.proto

package __

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
	Acceptor_Prepare_FullMethodName = "/message.Acceptor/Prepare"
	Acceptor_Accept_FullMethodName  = "/message.Acceptor/Accept"
)

// AcceptorClient is the client API for Acceptor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AcceptorClient interface {
	Prepare(ctx context.Context, in *MsgArgs, opts ...grpc.CallOption) (*MsgReply, error)
	Accept(ctx context.Context, in *MsgArgs, opts ...grpc.CallOption) (*MsgReply, error)
}

type acceptorClient struct {
	cc grpc.ClientConnInterface
}

func NewAcceptorClient(cc grpc.ClientConnInterface) AcceptorClient {
	return &acceptorClient{cc}
}

func (c *acceptorClient) Prepare(ctx context.Context, in *MsgArgs, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := c.cc.Invoke(ctx, Acceptor_Prepare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acceptorClient) Accept(ctx context.Context, in *MsgArgs, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := c.cc.Invoke(ctx, Acceptor_Accept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AcceptorServer is the server API for Acceptor service.
// All implementations must embed UnimplementedAcceptorServer
// for forward compatibility
type AcceptorServer interface {
	Prepare(context.Context, *MsgArgs) (*MsgReply, error)
	Accept(context.Context, *MsgArgs) (*MsgReply, error)
	mustEmbedUnimplementedAcceptorServer()
}

// UnimplementedAcceptorServer must be embedded to have forward compatible implementations.
type UnimplementedAcceptorServer struct {
}

func (UnimplementedAcceptorServer) Prepare(context.Context, *MsgArgs) (*MsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (UnimplementedAcceptorServer) Accept(context.Context, *MsgArgs) (*MsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accept not implemented")
}
func (UnimplementedAcceptorServer) mustEmbedUnimplementedAcceptorServer() {}

// UnsafeAcceptorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AcceptorServer will
// result in compilation errors.
type UnsafeAcceptorServer interface {
	mustEmbedUnimplementedAcceptorServer()
}

func RegisterAcceptorServer(s grpc.ServiceRegistrar, srv AcceptorServer) {
	s.RegisterService(&Acceptor_ServiceDesc, srv)
}

func _Acceptor_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcceptorServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Acceptor_Prepare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcceptorServer).Prepare(ctx, req.(*MsgArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acceptor_Accept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcceptorServer).Accept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Acceptor_Accept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcceptorServer).Accept(ctx, req.(*MsgArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// Acceptor_ServiceDesc is the grpc.ServiceDesc for Acceptor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Acceptor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.Acceptor",
	HandlerType: (*AcceptorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _Acceptor_Prepare_Handler,
		},
		{
			MethodName: "Accept",
			Handler:    _Acceptor_Accept_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paxos.proto",
}

const (
	Learner_Learn_FullMethodName = "/message.Learner/Learn"
)

// LearnerClient is the client API for Learner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LearnerClient interface {
	Learn(ctx context.Context, in *MsgArgs, opts ...grpc.CallOption) (*MsgReply, error)
}

type learnerClient struct {
	cc grpc.ClientConnInterface
}

func NewLearnerClient(cc grpc.ClientConnInterface) LearnerClient {
	return &learnerClient{cc}
}

func (c *learnerClient) Learn(ctx context.Context, in *MsgArgs, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := c.cc.Invoke(ctx, Learner_Learn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LearnerServer is the server API for Learner service.
// All implementations must embed UnimplementedLearnerServer
// for forward compatibility
type LearnerServer interface {
	Learn(context.Context, *MsgArgs) (*MsgReply, error)
	mustEmbedUnimplementedLearnerServer()
}

// UnimplementedLearnerServer must be embedded to have forward compatible implementations.
type UnimplementedLearnerServer struct {
}

func (UnimplementedLearnerServer) Learn(context.Context, *MsgArgs) (*MsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Learn not implemented")
}
func (UnimplementedLearnerServer) mustEmbedUnimplementedLearnerServer() {}

// UnsafeLearnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LearnerServer will
// result in compilation errors.
type UnsafeLearnerServer interface {
	mustEmbedUnimplementedLearnerServer()
}

func RegisterLearnerServer(s grpc.ServiceRegistrar, srv LearnerServer) {
	s.RegisterService(&Learner_ServiceDesc, srv)
}

func _Learner_Learn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnerServer).Learn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learner_Learn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnerServer).Learn(ctx, req.(*MsgArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// Learner_ServiceDesc is the grpc.ServiceDesc for Learner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Learner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.Learner",
	HandlerType: (*LearnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Learn",
			Handler:    _Learner_Learn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paxos.proto",
}
