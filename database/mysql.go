package database

import (
	"restapi/models"

	"github.com/jinzhu/gorm"
)

func ConnectDBMySQL() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/testinger?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	println(err)

	defer db.Close()
	db.AutoMigrate(&models.User{})

	return db, nil
}
