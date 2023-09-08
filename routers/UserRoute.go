package routers

import (
	"restapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	r.GET("/", controllers.WelcomeMessage)
	r.GET("/users")
	r.GET("/users/:id")
	r.POST("/create", controllers.CreateUser)
	r.PUT("/update")
	r.DELETE("/delete/:id")

	return r
}
