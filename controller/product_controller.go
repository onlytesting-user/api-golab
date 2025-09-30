// Package controller provides the controllers responsible for handling
// requests and coordinating between the HTTP layer and use cases.
package controller

type productController struct{
	//usecase
}

func NewProductController() productController {
	return productController{}
}
