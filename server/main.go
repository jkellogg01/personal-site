package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.StaticFile("/", "../client/dist/index.html")
	app.Static("/assets", "../client/dist/assets")

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	log.Fatal(
		app.Run(),
	)
}
