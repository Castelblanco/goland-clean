package main

import (
	"log"

	"github.com/Castelblanco/goland-clean/src/app/db"
	"github.com/Castelblanco/goland-clean/src/app/server"
	taskModelsSql "github.com/Castelblanco/goland-clean/src/modules/tasks/infrastructure/storages/sql/models"
	userModelsSql "github.com/Castelblanco/goland-clean/src/modules/users/infrastructure/storages/sql/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.PostgresConnect()
	userModelsSql.Migrations()
	taskModelsSql.Migrations()
	server.Server()
}
