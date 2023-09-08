package database

import (
	"restapi/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDBPostgres() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=192.168.1.184 port=5432 user=postgres dbname=joyjoydevelopment password=meen sslmode=disable")
	if err != nil {
		return nil, err
	}

	// defer db.Close()
	db.AutoMigrate(&models.User{})

	return db, nil
}
