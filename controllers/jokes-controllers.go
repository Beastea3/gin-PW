package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/beastea/golang-gin/models"
	"fmt"
)

func ShowJokes(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	jokes := []models.Joke{}
	err := db.C(models.CollectionJoke).Find(nil).Sort("-updated_on").All(&jokes)
	if err != nil {
		c.Error(err)
	}
	c.Header("Content-Type","application/json")
	c.JSON(http.StatusOK, gin.H{
		"title": "Jokes",
		"Jokes": jokes,
	})
}

// Create an article
func CreateJoke(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	joke := models.Joke{}
	err := c.Bind(&joke)
	if err != nil {
		c.Error(err)
		return
	}
	fmt.Printf("Happen Here");
	err = db.C(models.CollectionJoke).Insert(joke)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/jokes")
}
//func LikeJoke(c *gin.Context) {
//	c.Header("Content-Type","application/json")
//	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
//		c.JSON(http.StatusOK, services.LikeJoke(jokeid))
//	}  else {
//		c.AbortWithStatus(http.StatusNotFound)
//	}
//}