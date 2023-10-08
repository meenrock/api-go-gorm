package routers

import (
	"restapi/controllers"
	queue "restapi/controllers/queue"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	indexRoutes := r.Group("/")
	{
		indexRoutes.GET("/", controllers.WelcomeMessage)
	}

	UserRoutes := r.Group("/user")
	{
		UserRoutes.GET("/", controllers.WelcomeMessage)
		UserRoutes.GET("/getall", controllers.GetAllUser)
		UserRoutes.GET("/usersnew", controllers.GetUserFavFood)
		UserRoutes.GET("/get/:id", controllers.GetUserById)
		UserRoutes.POST("/create", controllers.CreateUser)
		UserRoutes.PUT("/update/:id", controllers.UpdateUserById)
		UserRoutes.DELETE("/delete/:id", controllers.DeleteUserById)
	}

	MealQueue := r.Group("/meal")
	{
		MealQueue.GET("/get", queue.CreateComment)
	}

	return r
}
