package main

import (
	handlers "sample01/handlers"
	"sample01/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectAndSyncDB()
}
func main() {
	router := gin.Default()
	router.GET("/GetData", handlers.GetHandler)
	router.POST("/Discount", handlers.DiscountHandler)
	router.POST("/RegisterHandler", handlers.RegisterHandler)
	router.Run(":9090")

}
