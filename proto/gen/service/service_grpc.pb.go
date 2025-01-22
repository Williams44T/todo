// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: service.proto

package service

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
	Todo_Signup_FullMethodName      = "/service.Todo/Signup"
	Todo_Signin_FullMethodName      = "/service.Todo/Signin"
	Todo_AddTask_FullMethodName     = "/service.Todo/AddTask"
	Todo_GetTask_FullMethodName     = "/service.Todo/GetTask"
	Todo_GetAllTasks_FullMethodName = "/service.Todo/GetAllTasks"
)

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoClient interface {
	Signup(ctx context.Context, in *SignupReq, opts ...grpc.CallOption) (*SignupResp, error)
	Signin(ctx context.Context, in *SigninReq, opts ...grpc.CallOption) (*SigninResp, error)
	AddTask(ctx context.Context, in *AddTaskReq, opts ...grpc.CallOption) (*AddTaskResp, error)
	GetTask(ctx context.Context, in *GetTaskReq, opts ...grpc.CallOption) (*GetTaskResp, error)
	GetAllTasks(ctx context.Context, in *GetAllTasksReq, opts ...grpc.CallOption) (*GetAllTasksResp, error)
}

type todoClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoClient(cc grpc.ClientConnInterface) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) Signup(ctx context.Context, in *SignupReq, opts ...grpc.CallOption) (*SignupResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignupResp)
	err := c.cc.Invoke(ctx, Todo_Signup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) Signin(ctx context.Context, in *SigninReq, opts ...grpc.CallOption) (*SigninResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SigninResp)
	err := c.cc.Invoke(ctx, Todo_Signin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) AddTask(ctx context.Context, in *AddTaskReq, opts ...grpc.CallOption) (*AddTaskResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddTaskResp)
	err := c.cc.Invoke(ctx, Todo_AddTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) GetTask(ctx context.Context, in *GetTaskReq, opts ...grpc.CallOption) (*GetTaskResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskResp)
	err := c.cc.Invoke(ctx, Todo_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) GetAllTasks(ctx context.Context, in *GetAllTasksReq, opts ...grpc.CallOption) (*GetAllTasksResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllTasksResp)
	err := c.cc.Invoke(ctx, Todo_GetAllTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
// All implementations must embed UnimplementedTodoServer
// for forward compatibility.
type TodoServer interface {
	Signup(context.Context, *SignupReq) (*SignupResp, error)
	Signin(context.Context, *SigninReq) (*SigninResp, error)
	AddTask(context.Context, *AddTaskReq) (*AddTaskResp, error)
	GetTask(context.Context, *GetTaskReq) (*GetTaskResp, error)
	GetAllTasks(context.Context, *GetAllTasksReq) (*GetAllTasksResp, error)
	mustEmbedUnimplementedTodoServer()
}

// UnimplementedTodoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTodoServer struct{}

func (UnimplementedTodoServer) Signup(context.Context, *SignupReq) (*SignupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedTodoServer) Signin(context.Context, *SigninReq) (*SigninResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signin not implemented")
}
func (UnimplementedTodoServer) AddTask(context.Context, *AddTaskReq) (*AddTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTask not implemented")
}
func (UnimplementedTodoServer) GetTask(context.Context, *GetTaskReq) (*GetTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTodoServer) GetAllTasks(context.Context, *GetAllTasksReq) (*GetAllTasksResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTasks not implemented")
}
func (UnimplementedTodoServer) mustEmbedUnimplementedTodoServer() {}
func (UnimplementedTodoServer) testEmbeddedByValue()              {}

// UnsafeTodoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServer will
// result in compilation errors.
type UnsafeTodoServer interface {
	mustEmbedUnimplementedTodoServer()
}

func RegisterTodoServer(s grpc.ServiceRegistrar, srv TodoServer) {
	// If the following call pancis, it indicates UnimplementedTodoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Todo_ServiceDesc, srv)
}

func _Todo_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Signup(ctx, req.(*SignupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_Signin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SigninReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Signin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_Signin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Signin(ctx, req.(*SigninReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_AddTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).AddTask(ctx, req.(*AddTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetTask(ctx, req.(*GetTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_GetAllTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTasksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetAllTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Todo_GetAllTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetAllTasks(ctx, req.(*GetAllTasksReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Todo_ServiceDesc is the grpc.ServiceDesc for Todo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _Todo_Signup_Handler,
		},
		{
			MethodName: "Signin",
			Handler:    _Todo_Signin_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _Todo_AddTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _Todo_GetTask_Handler,
		},
		{
			MethodName: "GetAllTasks",
			Handler:    _Todo_GetAllTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
