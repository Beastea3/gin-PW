package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func Api(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
		"message": "PONG",
	})
}

func UploadFiles(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	if err := c.SaveUploadedFile(file, "./public/" + file.Filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
}