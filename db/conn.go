// Package db provides the database connection for the application.
package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

func ConnectDB() (*sql.DB, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to: ", os.Getenv("DB_NAME"))

	return db, nil
}
