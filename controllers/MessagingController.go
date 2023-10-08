package controllers

import (
	"restapi/database"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

func GetKafkaMessage(c *gin.Context) {
	db, err := database.ConnectDBPostgres()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to connect to the database", "detail": err.Error()})
		return
	}
	defer db.Close()

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}

	db.Create(&user)
	c.JSON(201, user)
}
