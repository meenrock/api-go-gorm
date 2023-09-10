package controllers

import (
	"restapi/database"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
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

func GetAllUser(c *gin.Context) {
	db, err := database.ConnectDBPostgres()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to connect to the database", "detail": err.Error()})
		return
	}
	defer db.Close()

	var user []models.User
	if err := db.Find(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func GetUserById(c *gin.Context) {
	db, err := database.ConnectDBPostgres()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to connect to the database", "detail": err.Error()})
		return
	}
	defer db.Close()

	userId := c.Param("id")

	var user models.User
	if err := db.First(&user, userId).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func UpdateUserById(c *gin.Context) {
	db, err := database.ConnectDBPostgres()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to connect to the database", "detail": err.Error()})
		return
	}
	defer db.Close()

	userID := c.Param("id")

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}

	db.Save(&user)
	c.JSON(200, user)
}

func DeleteUserById(c *gin.Context) {
	db, err := database.ConnectDBPostgres()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to connect to the database", "detail": err.Error()})
		return
	}
	defer db.Close()

	userID := c.Param("id")

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	db.Delete(&user)
	c.JSON(204, user.FirstName)
}
