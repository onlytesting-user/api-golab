package main

import (
	"github.com/gin-gonic/gin"
	"github.com/onlytesting-user/api-golab/db"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	_ = dbConnection

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	server.Run(":8000")
}
