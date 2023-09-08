package controllers

import (
	"github.com/gin-gonic/gin"
)

func WelcomeMessage(c *gin.Context) {
	c.JSON(201, gin.H{
		"id":      "6670322321",
		"author":  "Supawat Suntornlimsiri",
		"message": "Welcome to my api assignment",
	})
}
