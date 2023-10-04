package routers

import (
	"restapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	r.GET("/", controllers.WelcomeMessage)
	r.GET("/users", controllers.GetAllUser)
	r.GET("/usersnew", controllers.GetUserFavFood)
	r.GET("/users/:id", controllers.GetUserById)
	r.POST("/create", controllers.CreateUser)
	r.PUT("/update/:id", controllers.UpdateUserById)
	r.DELETE("/delete/:id", controllers.DeleteUserById)

	return r
}
