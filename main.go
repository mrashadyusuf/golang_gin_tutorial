package main

import (
	"go-restapi-gin/controllers/productcontroller"
	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/search", productcontroller.Search)
	r.GET("/api/product/:id", productcontroller.Show)	

	r.POST("/api/product", productcontroller.Create)
	r.POST("/api/test", productcontroller.Test)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
