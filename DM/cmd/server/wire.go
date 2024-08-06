// //go:build wireinject
// // +build wireinject

// package wire

// import (
// 	"DM/internal/biz"
// 	"DM/internal/data"
// 	"DM/internal/server"
// 	"DM/internal/service"

// 	"github.com/go-kratos/kratos/v2/transport/http"

// 	"github.com/google/wire"
// 	"gorm.io/gorm"
// )

// func InitializeDepartmentService(db *gorm.DB) (*service.DepartmentService, error) {
// 	wire.Build(
// 		data.NewDepartmentRepository,
// 		biz.NewDepartmentService,
// 		service.NewDepartmentService,
// 	)
// 	return &service.DepartmentService{}, nil
// }

// // func InitializeGRPCServer(deptService *service.DepartmentService) *grpc.Server {
// // 	wire.Build(server.NewGRPCServer)
// // 	return &grpc.Server{}
// // }

// func InitializeHTTPServer(deptService *service.DepartmentService) *http.Server {
// 	wire.Build(server.NewHTTPServer)
// 	return &http.Server{}
// }
//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"DM/internal/biz"
	"DM/internal/conf"
	"DM/internal/data"
	"DM/internal/server"
	"DM/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
