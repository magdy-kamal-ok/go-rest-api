package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setup Db connection")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disabled", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	db, err := gorm.Open("postgress", connectionStr)
	if err != nil {
		return db, err
	}
	if err := db.DB().Ping(); err != nil {
		return db, err
	}
	return db, nil
}
