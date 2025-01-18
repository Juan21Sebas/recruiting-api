package main

import (
	"log"

	"RECRUITING-API/cmd/config/db"
	server "RECRUITING-API/internal/infrastructure/api/candidates"
)

func main() {
	dbConfig := db.DBConfig{
		Host:     "localhost",
		Port:     "3306",
		Database: "recruiting_db",
		User:     "juansebastiansanchez",
		Password: "JuanAdmin123$",
	}

	database, err := db.NewMySQLDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	server.RunServer(database)
}
