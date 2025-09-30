// Package usecase provides the application business logic, orchestrating
// operations between controllers and repositories.
package usecase

type ProductUsecase struct {
	// Repository
}

func NewProductUsecase() ProductUsecase {
	return ProductUsecase{}
}
