package main

import (
	"crud/controllers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api")
	{
		client.GET("/story/:id", controllers.Read)
		//client.POST("/story/create", controllers.Create)
		//client.PATCH("/story/update/:id", controllers.Update)
		//client.DELETE("/story/:id", controllers.Delete)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run("http://localhost:8080") // Ứng dụng chạy tại cổng 8080
}
