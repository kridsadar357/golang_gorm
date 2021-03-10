package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&Todo{})

	DB = database
}
