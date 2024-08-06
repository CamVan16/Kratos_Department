package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewDepartmentUC, NewSubDepartmentUC, NewEmployeeUC, NewUserUC)
