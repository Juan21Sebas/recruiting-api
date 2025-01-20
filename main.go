package main

import (
	"log"

	"RECRUITING-API/cmd/config/db"
	server "RECRUITING-API/internal/infrastructure/api/candidates"
)

func main() {
	dbConfig := db.DBConfig{
		Host:     "",
		Port:     "",
		Database: "",
		User:     "",
		Password: "",
	}

	database, err := db.NewMySQLDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	server.RunServer(database)
}
