package server

import (
	v1 "DM/api/department/v1"
	v3 "DM/api/employee/v1"
	v2 "DM/api/sub_department/v1"
	v4 "DM/api/user/v1"
	"DM/internal/conf"
	"DM/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, department *service.DepartmentService, subDepartment *service.SubDepartmentService, employee *service.EmployeeService, user *service.UserService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			AuthMiddleware("admin", "manager"),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterDepartmentServiceServer(srv, department)
	v2.RegisterSubDepartmentServiceServer(srv, subDepartment)
	v3.RegisterEmployeeServiceServer(srv, employee)
	v4.RegisterUserServiceServer(srv, user)
	return srv
}
