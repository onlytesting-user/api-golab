// Package repository contains the data access layer responsible for interacting
// with the database and persisting domain entities.
package repository

import (
	"database/sql"
	"fmt"

	"github.com/onlytesting-user/api-golab/model"
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

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}
