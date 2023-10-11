package main

import "github.com/rganes5/LCO/webServerGin/handlers"

func main() {
	r := gin.Default()
	r.GET("/getendpoint", handlers.GetEndPoint)
	r.POST("/disc", handlers.NumHandler)
	r.Run(":8080")
}
