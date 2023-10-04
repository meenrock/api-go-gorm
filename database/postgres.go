package database

import (
	"database/sql"
	"log"
	"os"
	"restapi/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func ConnectDBPostgres() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env unable to load")
	}

	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("SERVER_NAME")

	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPwd + " sslmode=disable"

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// defer db.Close()
	db.AutoMigrate(&models.User{})

	return db, nil
}

func ExecuteQueryPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", "connection-string")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db, err
}
