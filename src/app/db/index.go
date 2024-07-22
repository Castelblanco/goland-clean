package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresConnect() {
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseInt(p, 10, 32)

	if err != nil {
		panic(err)
	}

	var DSN = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		log.Println("DB Connected")
	}

	DB = db
}
