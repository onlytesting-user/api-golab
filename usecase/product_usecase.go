// Package usecase provides the application business logic, orchestrating
// operations between controllers and repositories.
package usecase

import (
	"github.com/onlytesting-user/api-golab/model"
	"github.com/onlytesting-user/api-golab/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
