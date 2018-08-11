package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/beastea/golang-gin/controllers"
	"gopkg.in/mgo.v2"
	"github.com/beastea/golang-gin/db"
	"github.com/beastea/golang-gin/middlewares"
	"os"
)

const (
	PORT = "3000"
)

var jokesCollection *mgo.Collection

func init() {
	db.Connect()
}

func main() {
	router := gin.Default()

	router.Use(middlewares.ErrorHandler)
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	api.Use(middlewares.Connect)
	{
		api.GET("/", controllers.Welcome)
		api.GET("/joke", controllers.Api)
		api.POST("/uploadFiles", controllers.UploadFiles)
		api.GET("/jokes", controllers.ShowJokes)
		api.POST("/joke", controllers.CreateJoke)
	}

	port :=  PORT

	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	router.Run(":" + port)
}