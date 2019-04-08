package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// InitDB conn
func InitDB() *gorm.DB {
	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}

	dbType := os.Getenv("DB_TYPE")

	db, err := gorm.Open(dbType, "./data.db")

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	return db
}
