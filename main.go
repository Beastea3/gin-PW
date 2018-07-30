package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/beastea/golang-gin/controllers"
)

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", controllers.Welcome)
		api.GET("/joke", controllers.Api)
		api.POST("/uploadFiles", controllers.UploadFiles)
	}
	api.GET("/jokes", controllers.JokeHandler)
	api.POST("/jokes/like/:jokeID", controllers.LikeJoke)

	router.Run(":3000")
}