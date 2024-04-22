package database

import (
	"github.com/davidgrldo/go-restapi-fiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3308)/go-fiber-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migrate the model
	db.AutoMigrate(&models.Product{})

	return db
}
