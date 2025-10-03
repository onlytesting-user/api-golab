// Package controller provides the controllers responsible for handling
// requests and coordinating between the HTTP layer and use cases.
package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/onlytesting-user/api-golab/dto"
	"github.com/onlytesting-user/api-golab/model"
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
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("productID")
	if id == "" {
		response := dto.Response{
			Message: "The product id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := dto.Response{
			Message: "The product id must be an Integer",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := dto.Response{
			Message: "The product was not found in the database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	var product model.Product

	// Captura o ID do path e converte para int
	idParam := ctx.Param("productID")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	product.ID = productID

	// Faz o bind do JSON recebido
	err = ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o usecase para atualizar
	updatedProduct, err := p.productUseCase.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna o produto atualizado
	ctx.JSON(http.StatusOK, updatedProduct)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	idParam := ctx.Param("productID")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	err = p.productUseCase.DeleteProduct(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
