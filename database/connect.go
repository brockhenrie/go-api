package database

import (
	"api/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


var DB *gorm.DB


func Connect(){
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3300)/go_admin"), &gorm.Config{})
	if err != nil{
		panic("could not connect to database")
	}

	DB = db

	db.AutoMigrate(&models.User{},&models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}