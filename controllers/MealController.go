package controllers

import (
	"net/http"
	"restapi/database"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

func GetUserFavFood(c *gin.Context) {
	db, err := database.ExecuteQueryPostgres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database", "detail": err.Error()})
		return
	}
	defer db.Close()

	var users []models.User

	userFoodRepo := NewUserFoodRepository(db)

	for i := range users {
		id := users[i].UserUuid

		favFoodRows, err := userFoodRepo.GetUserFavFoodByUserId(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch favorite foods", "detail": err.Error()})
			return
		}

		defer favFoodRows.Close()

	}

	c.JSON(http.StatusOK, users)
}
