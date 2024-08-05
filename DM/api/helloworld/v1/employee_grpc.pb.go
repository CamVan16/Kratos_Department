// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: employee.proto

package v1

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
	EmployeeService_CreateEmployee_FullMethodName         = "/helloworld.v1.EmployeeService/CreateEmployee"
	EmployeeService_GetAllEmployee_FullMethodName         = "/helloworld.v1.EmployeeService/GetAllEmployee"
	EmployeeService_GetEmployeeByID_FullMethodName        = "/helloworld.v1.EmployeeService/GetEmployeeByID"
	EmployeeService_UpdateEmployee_FullMethodName         = "/helloworld.v1.EmployeeService/UpdateEmployee"
	EmployeeService_DeleteEmployee_FullMethodName         = "/helloworld.v1.EmployeeService/DeleteEmployee"
	EmployeeService_GetEmployeeByPhonePass_FullMethodName = "/helloworld.v1.EmployeeService/GetEmployeeByPhonePass"
)

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	CreateEmployee(ctx context.Context, in *CreateEmployeeRequest, opts ...grpc.CallOption) (*CreateEmployeeRespone, error)
	GetAllEmployee(ctx context.Context, in *GetAllEmployeeRequest, opts ...grpc.CallOption) (*GetAllEmployeeRespone, error)
	GetEmployeeByID(ctx context.Context, in *GetEmployeeByIDResquest, opts ...grpc.CallOption) (*GetEmployeeByIDRespone, error)
	UpdateEmployee(ctx context.Context, in *UpdateEmployeeResquest, opts ...grpc.CallOption) (*UpdateEmployeeRespone, error)
	DeleteEmployee(ctx context.Context, in *DeleteEmployeeResquest, opts ...grpc.CallOption) (*DeleteEmployeeRespone, error)
	GetEmployeeByPhonePass(ctx context.Context, in *GetEmployeeByPhonePassRequest, opts ...grpc.CallOption) (*GetEmployeeByPhonePassRespone, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) CreateEmployee(ctx context.Context, in *CreateEmployeeRequest, opts ...grpc.CallOption) (*CreateEmployeeRespone, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateEmployeeRespone)
	err := c.cc.Invoke(ctx, EmployeeService_CreateEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetAllEmployee(ctx context.Context, in *GetAllEmployeeRequest, opts ...grpc.CallOption) (*GetAllEmployeeRespone, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllEmployeeRespone)
	err := c.cc.Invoke(ctx, EmployeeService_GetAllEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetEmployeeByID(ctx context.Context, in *GetEmployeeByIDResquest, opts ...grpc.CallOption) (*GetEmployeeByIDRespone, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEmployeeByIDRespone)
	err := c.cc.Invoke(ctx, EmployeeService_GetEmployeeByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) UpdateEmployee(ctx context.Context, in *UpdateEmployeeResquest, opts ...grpc.CallOption) (*UpdateEmployeeRespone, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEmployeeRespone)
	err := c.cc.Invoke(ctx, EmployeeService_UpdateEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) DeleteEmployee(ctx context.Context, in *DeleteEmployeeResquest, opts ...grpc.CallOption) (*DeleteEmployeeRespone, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteEmployeeRespone)
	err := c.cc.Invoke(ctx, EmployeeService_DeleteEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetEmployeeByPhonePass(ctx context.Context, in *GetEmployeeByPhonePassRequest, opts ...grpc.CallOption) (*GetEmployeeByPhonePassRespone, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEmployeeByPhonePassRespone)
	err := c.cc.Invoke(ctx, EmployeeService_GetEmployeeByPhonePass_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
// All implementations must embed UnimplementedEmployeeServiceServer
// for forward compatibility.
type EmployeeServiceServer interface {
	CreateEmployee(context.Context, *CreateEmployeeRequest) (*CreateEmployeeRespone, error)
	GetAllEmployee(context.Context, *GetAllEmployeeRequest) (*GetAllEmployeeRespone, error)
	GetEmployeeByID(context.Context, *GetEmployeeByIDResquest) (*GetEmployeeByIDRespone, error)
	UpdateEmployee(context.Context, *UpdateEmployeeResquest) (*UpdateEmployeeRespone, error)
	DeleteEmployee(context.Context, *DeleteEmployeeResquest) (*DeleteEmployeeRespone, error)
	GetEmployeeByPhonePass(context.Context, *GetEmployeeByPhonePassRequest) (*GetEmployeeByPhonePassRespone, error)
	mustEmbedUnimplementedEmployeeServiceServer()
}

// UnimplementedEmployeeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEmployeeServiceServer struct{}

func (UnimplementedEmployeeServiceServer) CreateEmployee(context.Context, *CreateEmployeeRequest) (*CreateEmployeeRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) GetAllEmployee(context.Context, *GetAllEmployeeRequest) (*GetAllEmployeeRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) GetEmployeeByID(context.Context, *GetEmployeeByIDResquest) (*GetEmployeeByIDRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeByID not implemented")
}
func (UnimplementedEmployeeServiceServer) UpdateEmployee(context.Context, *UpdateEmployeeResquest) (*UpdateEmployeeRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) DeleteEmployee(context.Context, *DeleteEmployeeResquest) (*DeleteEmployeeRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) GetEmployeeByPhonePass(context.Context, *GetEmployeeByPhonePassRequest) (*GetEmployeeByPhonePassRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeByPhonePass not implemented")
}
func (UnimplementedEmployeeServiceServer) mustEmbedUnimplementedEmployeeServiceServer() {}
func (UnimplementedEmployeeServiceServer) testEmbeddedByValue()                         {}

// UnsafeEmployeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServiceServer will
// result in compilation errors.
type UnsafeEmployeeServiceServer interface {
	mustEmbedUnimplementedEmployeeServiceServer()
}

func RegisterEmployeeServiceServer(s grpc.ServiceRegistrar, srv EmployeeServiceServer) {
	// If the following call pancis, it indicates UnimplementedEmployeeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EmployeeService_ServiceDesc, srv)
}

func _EmployeeService_CreateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).CreateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_CreateEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).CreateEmployee(ctx, req.(*CreateEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetAllEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetAllEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_GetAllEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetAllEmployee(ctx, req.(*GetAllEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetEmployeeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeeByIDResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetEmployeeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_GetEmployeeByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetEmployeeByID(ctx, req.(*GetEmployeeByIDResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployeeResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_UpdateEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, req.(*UpdateEmployeeResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmployeeResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_DeleteEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, req.(*DeleteEmployeeResquest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetEmployeeByPhonePass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeeByPhonePassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetEmployeeByPhonePass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_GetEmployeeByPhonePass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetEmployeeByPhonePass(ctx, req.(*GetEmployeeByPhonePassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmployeeService_ServiceDesc is the grpc.ServiceDesc for EmployeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmployee",
			Handler:    _EmployeeService_CreateEmployee_Handler,
		},
		{
			MethodName: "GetAllEmployee",
			Handler:    _EmployeeService_GetAllEmployee_Handler,
		},
		{
			MethodName: "GetEmployeeByID",
			Handler:    _EmployeeService_GetEmployeeByID_Handler,
		},
		{
			MethodName: "UpdateEmployee",
			Handler:    _EmployeeService_UpdateEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _EmployeeService_DeleteEmployee_Handler,
		},
		{
			MethodName: "GetEmployeeByPhonePass",
			Handler:    _EmployeeService_GetEmployeeByPhonePass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee.proto",
}
