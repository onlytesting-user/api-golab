// Package usecase provides the application business logic, orchestrating
// operations between controllers and repositories.
package usecase

import "github.com/onlytesting-user/api-golab/model"

type ProductUsecase struct {
	// Repository
}

func NewProductUsecase() ProductUsecase {
	return ProductUsecase{}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
