package database

import (
	"log"
	"os"
	"restapi/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func ConnectDBMySQL() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("env unable to load")
	}

	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("SERVER_NAME")

	dsn := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	println(err)

	defer db.Close()
	db.AutoMigrate(&models.User{})

	return db, nil
}
