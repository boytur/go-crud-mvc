package config

import (
	"github.com/boytur/go-crud-mvc/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:password@tcp(localhost:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&models.User{})
	DB = database
}
