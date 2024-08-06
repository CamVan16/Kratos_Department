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
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, deptService *service.DepartmentService, subDepartment *service.SubDepartmentService, employee *service.EmployeeService, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			AuthMiddleware("admin", "manager"),
		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterDepartmentServiceHTTPServer(srv, deptService)
	v2.RegisterSubDepartmentServiceHTTPServer(srv, subDepartment)
	v3.RegisterEmployeeServiceHTTPServer(srv, employee)
	v4.RegisterUserServiceHTTPServer(srv, user)
	return srv
}
