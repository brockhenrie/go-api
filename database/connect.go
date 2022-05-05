package database
import (
	"api/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


var DB *gorm.DB


func Connect(){
	db, err := gorm.Open(mysql.Open("root:$$B@ckspace1!!@/go_admin"), &gorm.Config{})
	if err != nil{
		panic("could not connect to database")
	}

	DB = db

	db.AutoMigrate(&models.User{},&models.Role{}, &models.Permission{})
}