package main

import (
	"express-to-gin/middleware"
	"express-to-gin/routes"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setupLogOutputToFile() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutputToFile()

	server := gin.New()

	server.Use(gin.Recovery())
	server.Use(middleware.Logger())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	routes.InitTodos(server)
	routes.InitComments(server)

	server.Run(":8000")
}
