package data

import (
	"DM/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// // ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDepartmentRepository, NewSubDepartmentRepository, NewEmployeeRepository, NewUserRepository)

// Data .
type Data struct {
	// 	// TODO wrapped database client
	DB *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{DB: db}, cleanup, nil
}
