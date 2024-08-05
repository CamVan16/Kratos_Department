package connection

import (
	"DM/internal/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := "root:Camvan12345678#@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB.AutoMigrate(&models.Department{}, &models.SubDepartment{}, models.Employee{}, models.User{})
	DB = database
	return DB
}
