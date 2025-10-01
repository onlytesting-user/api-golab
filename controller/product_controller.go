// Package controller provides the controllers responsible for handling
// requests and coordinating between the HTTP layer and use cases.
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlytesting-user/api-golab/usecase"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if (err != nil) {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
