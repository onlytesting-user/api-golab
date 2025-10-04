package main

import (
	"github.com/gin-gonic/gin"
	"github.com/onlytesting-user/api-golab/controller"
	"github.com/onlytesting-user/api-golab/db"
	docs "github.com/onlytesting-user/api-golab/docs"
	"github.com/onlytesting-user/api-golab/repository"
	"github.com/onlytesting-user/api-golab/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Repository Layer
	ProductRepository := repository.NewProductRepository(dbConnection)
	// Usecase Layer
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)
	// Controller Layer
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productID", ProductController.GetProductByID)
	server.PUT("/product/:productID", ProductController.UpdateProduct)
	server.DELETE("/product/:productID", ProductController.DeleteProduct)

	// Swagger UI
	docs.SwaggerInfo.Title = "API Go Lab Tutoriais"
	docs.SwaggerInfo.Description = "Documentação da API Go Lab Tutoriais"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8000")
}
