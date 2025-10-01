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

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productID, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID

	return product, nil
}

func (pu *ProductUsecase) GetProductByID(idProduct int) (model.Product, error) {
	product, err := pu.repository.GetProductByID(idProduct)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}
