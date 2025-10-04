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

// GetProducts godoc
// @Summary      List all products
// @Description  Get all products from the database
// @Tags         products
// @Produce      json
// @Success      200  {array}   model.Product
// @Failure      500  {object}  dto.Response
// @Router       /products [get]
func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product in the database
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      dto.CreateProductRequest  true  "Product to create"
// @Success      201      {object}  model.Product
// @Failure      400      {object}  dto.Response
// @Failure      500      {object}  dto.Response
// @Router       /product [post]
func (p *productController) CreateProduct(ctx *gin.Context) {
	var req dto.CreateProductRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{Message: err.Error()})
		return
	}

	product := model.Product{
		Name:  req.ProductName,
		Price: req.Price,
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

// GetProductByID godoc
// @Summary      Get product by ID
// @Description  Retrieve a single product by its ID
// @Tags         products
// @Produce      json
// @Param        productID  path      int  true  "Product ID"
// @Success      200        {object}  model.Product
// @Failure      400        {object}  dto.Response
// @Failure      404        {object}  dto.Response
// @Failure      500        {object}  dto.Response
// @Router       /product/{productID} [get]
func (p *productController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("productID")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, dto.Response{Message: "The product id cannot be null"})
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{Message: "The product id must be an integer"})
		return
	}

	product, err := p.productUseCase.GetProductByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{Message: err.Error()})
		return
	}

	if product == nil {
		ctx.JSON(http.StatusNotFound, dto.Response{Message: "The product was not found in the database"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// UpdateProduct godoc
// @Summary      Update an existing product
// @Description  Update an existing product in the database
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        productID  path      int                     true  "Product ID"
// @Param        product    body      dto.UpdateProductRequest  true  "Product data to update"
// @Success      200        {object}  model.Product
// @Failure      400        {object}  dto.Response
// @Failure      500        {object}  dto.Response
// @Router       /product/{productID} [put]
func (p *productController) UpdateProduct(ctx *gin.Context) {
	idParam := ctx.Param("productID")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{Message: "Invalid product ID"})
		return
	}

	var req dto.UpdateProductRequest
	if err = ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{Message: err.Error()})
		return
	}

	product := model.Product{
		ID:    productID,
		Name:  req.ProductName,
		Price: req.Price,
	}

	updatedProduct, err := p.productUseCase.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product from the database
// @Tags         products
// @Produce      json
// @Param        productID  path      int  true  "Product ID"
// @Success      200        {object}  dto.Response
// @Failure      400        {object}  dto.Response
// @Failure      500        {object}  dto.Response
// @Router       /product/{productID} [delete]
func (p *productController) DeleteProduct(ctx *gin.Context) {
	idParam := ctx.Param("productID")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{Message: "Invalid product ID"})
		return
	}

	if err := p.productUseCase.DeleteProduct(productID); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Message: "Product deleted successfully"})
}
