// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: pkg/allieflowkitgrpc/allie-flowkit.proto

package allieflowkitgrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ExternalFunctions_ListFunctions_FullMethodName  = "/allieflowkitgrpc.ExternalFunctions/ListFunctions"
	ExternalFunctions_RunFunction_FullMethodName    = "/allieflowkitgrpc.ExternalFunctions/RunFunction"
	ExternalFunctions_StreamFunction_FullMethodName = "/allieflowkitgrpc.ExternalFunctions/StreamFunction"
)

// ExternalFunctionsClient is the client API for ExternalFunctions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ExternalFunctions is a gRPC service that allows for listing and running
// the function available in the externalfunctions package.
type ExternalFunctionsClient interface {
	// Lists all available functions with description, inputs and outputs.
	ListFunctions(ctx context.Context, in *ListFunctionsRequest, opts ...grpc.CallOption) (*ListFunctionsResponse, error)
	// Runs a specified function with provided inputs and returns the function outputs.
	RunFunction(ctx context.Context, in *FunctionInputs, opts ...grpc.CallOption) (*FunctionOutputs, error)
	// Runs a specified function with provided inputs and returns the function output as a stream.
	StreamFunction(ctx context.Context, in *FunctionInputs, opts ...grpc.CallOption) (ExternalFunctions_StreamFunctionClient, error)
}

type externalFunctionsClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalFunctionsClient(cc grpc.ClientConnInterface) ExternalFunctionsClient {
	return &externalFunctionsClient{cc}
}

func (c *externalFunctionsClient) ListFunctions(ctx context.Context, in *ListFunctionsRequest, opts ...grpc.CallOption) (*ListFunctionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFunctionsResponse)
	err := c.cc.Invoke(ctx, ExternalFunctions_ListFunctions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalFunctionsClient) RunFunction(ctx context.Context, in *FunctionInputs, opts ...grpc.CallOption) (*FunctionOutputs, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FunctionOutputs)
	err := c.cc.Invoke(ctx, ExternalFunctions_RunFunction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalFunctionsClient) StreamFunction(ctx context.Context, in *FunctionInputs, opts ...grpc.CallOption) (ExternalFunctions_StreamFunctionClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExternalFunctions_ServiceDesc.Streams[0], ExternalFunctions_StreamFunction_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &externalFunctionsStreamFunctionClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExternalFunctions_StreamFunctionClient interface {
	Recv() (*StreamOutput, error)
	grpc.ClientStream
}

type externalFunctionsStreamFunctionClient struct {
	grpc.ClientStream
}

func (x *externalFunctionsStreamFunctionClient) Recv() (*StreamOutput, error) {
	m := new(StreamOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExternalFunctionsServer is the server API for ExternalFunctions service.
// All implementations must embed UnimplementedExternalFunctionsServer
// for forward compatibility
//
// ExternalFunctions is a gRPC service that allows for listing and running
// the function available in the externalfunctions package.
type ExternalFunctionsServer interface {
	// Lists all available functions with description, inputs and outputs.
	ListFunctions(context.Context, *ListFunctionsRequest) (*ListFunctionsResponse, error)
	// Runs a specified function with provided inputs and returns the function outputs.
	RunFunction(context.Context, *FunctionInputs) (*FunctionOutputs, error)
	// Runs a specified function with provided inputs and returns the function output as a stream.
	StreamFunction(*FunctionInputs, ExternalFunctions_StreamFunctionServer) error
	mustEmbedUnimplementedExternalFunctionsServer()
}

// UnimplementedExternalFunctionsServer must be embedded to have forward compatible implementations.
type UnimplementedExternalFunctionsServer struct {
}

func (UnimplementedExternalFunctionsServer) ListFunctions(context.Context, *ListFunctionsRequest) (*ListFunctionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFunctions not implemented")
}
func (UnimplementedExternalFunctionsServer) RunFunction(context.Context, *FunctionInputs) (*FunctionOutputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunFunction not implemented")
}
func (UnimplementedExternalFunctionsServer) StreamFunction(*FunctionInputs, ExternalFunctions_StreamFunctionServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamFunction not implemented")
}
func (UnimplementedExternalFunctionsServer) mustEmbedUnimplementedExternalFunctionsServer() {}

// UnsafeExternalFunctionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalFunctionsServer will
// result in compilation errors.
type UnsafeExternalFunctionsServer interface {
	mustEmbedUnimplementedExternalFunctionsServer()
}

func RegisterExternalFunctionsServer(s grpc.ServiceRegistrar, srv ExternalFunctionsServer) {
	s.RegisterService(&ExternalFunctions_ServiceDesc, srv)
}

func _ExternalFunctions_ListFunctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFunctionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalFunctionsServer).ListFunctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalFunctions_ListFunctions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalFunctionsServer).ListFunctions(ctx, req.(*ListFunctionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalFunctions_RunFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FunctionInputs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalFunctionsServer).RunFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalFunctions_RunFunction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalFunctionsServer).RunFunction(ctx, req.(*FunctionInputs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalFunctions_StreamFunction_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FunctionInputs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExternalFunctionsServer).StreamFunction(m, &externalFunctionsStreamFunctionServer{ServerStream: stream})
}

type ExternalFunctions_StreamFunctionServer interface {
	Send(*StreamOutput) error
	grpc.ServerStream
}

type externalFunctionsStreamFunctionServer struct {
	grpc.ServerStream
}

func (x *externalFunctionsStreamFunctionServer) Send(m *StreamOutput) error {
	return x.ServerStream.SendMsg(m)
}

// ExternalFunctions_ServiceDesc is the grpc.ServiceDesc for ExternalFunctions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalFunctions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "allieflowkitgrpc.ExternalFunctions",
	HandlerType: (*ExternalFunctionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFunctions",
			Handler:    _ExternalFunctions_ListFunctions_Handler,
		},
		{
			MethodName: "RunFunction",
			Handler:    _ExternalFunctions_RunFunction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamFunction",
			Handler:       _ExternalFunctions_StreamFunction_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/allieflowkitgrpc/allie-flowkit.proto",
}
