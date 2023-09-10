package main

import (
	"webServerGin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/getendpoint", handlers.GetEndPoint)
	r.POST("/disc", handlers.NumHandler)
	r.Run(":8080")
}
