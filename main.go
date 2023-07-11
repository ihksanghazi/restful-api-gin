package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/restful-api-gin/controllers/productcontroller"
	"github.com/ihksanghazi/restful-api-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products/:id", productcontroller.Delete)

	r.Run(":5000")
}