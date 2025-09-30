// Package controller provides the controllers responsible for handling
// requests and coordinating between the HTTP layer and use cases.
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlytesting-user/api-golab/model"
)

type productController struct{
	//usecase
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			ID: 1,
			Name: "Batata Frita",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, products)

}
