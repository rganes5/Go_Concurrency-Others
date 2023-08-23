package main

import (
	handlers "sample/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/GetData", handlers.GetHandler)
	router.POST("/Discount", handlers.DiscountHandler)
	router.POST("/RegisterHandler", handlers.RegisterHandler)
	router.Run(":9090")

}
