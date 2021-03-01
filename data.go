package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	// dbName := "notes"
	// port := 5432
	// pass := "postgres"
	// user := "postgres"
	// host := "localhost"
	// ssl := "disable"

	// connectionLink := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbName, pass, ssl)

	connectionLink := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connectionLink), &gorm.Config{})
	log.Println(err)

	err = db.AutoMigrate(&note{})
	log.Println(err)
	err = db.AutoMigrate(&todo{})
	log.Println(err)

	return db
}
