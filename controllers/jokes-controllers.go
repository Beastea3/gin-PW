package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/beastea/golang-gin/services"
	"strconv"
)

func JokeHandler(c *gin.Context) {
	c.Header("Content-Type","application/json")
	c.JSON(http.StatusOK, services.JokeHandler())
}

func LikeJoke(c *gin.Context) {
	c.Header("Content-Type","application/json")
	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		c.JSON(http.StatusOK, services.LikeJoke(jokeid))
	}  else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}