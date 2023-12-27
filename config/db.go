package config

import (
	"golang_basic_gin_sept_2023/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_basic_gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed To Connect Database")
	}

	DB.AutoMigrate(&models.Department{}, &models.Position{}, &models.Inventory{}, &models.Archive{}, &models.Employee{})
}
