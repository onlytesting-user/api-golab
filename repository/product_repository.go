// Package repository contains the data access layer responsible for interacting
// with the database and persisting domain entities.
package repository

import (
	"database/sql"
)

// ProductRepository provides methods to perform operations on products
// within the database.
type ProductRepository struct {
	connection *sql.DB
}

// NewProductRepository creates and returns a new instance of ProductRepository
// with the given database connection.
func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}
