package sqlite

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect is a method used to connect to sqlite
func Connect() *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open("./ecommerce.db"),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("Failed to connect to database.")
	}

	return db
}
