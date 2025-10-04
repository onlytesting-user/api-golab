// Package dto contains Data Transfer Objects (DTOs) used to define
// input and output structures for the API, especially HTTP requests
// and responses.
package dto

type CreateProductRequest struct {
    ProductName string  `json:"product_name" example:"Keyboard"`
    Price       float64 `json:"price" example:"50.00"`
}

type UpdateProductRequest struct {
    ProductName string  `json:"product_name" example:"Keyboard"`
    Price       float64 `json:"price" example:"50.00"`
}
