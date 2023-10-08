package routers

// import (
// 	"restapi/controllers"

// 	"github.com/gin-gonic/gin"
// )

// // func SetupRouter() *gin.Engine {
// // 	r := gin.Default()

// // 	// User routes
// // 	r.GET("/", controllers.WelcomeMessage)
// // 	r.GET("/users", controllers.GetAllUser)
// // 	r.GET("/usersnew", controllers.GetUserFavFood)
// // 	r.GET("/users/:id", controllers.GetUserById)
// // 	r.POST("/create", controllers.CreateUser)
// // 	r.PUT("/update/:id", controllers.UpdateUserById)
// // 	r.DELETE("/delete/:id", controllers.DeleteUserById)

// // 	return r
// // }

// func SetupUserRoutes(r *gin.RouterGroup, ctrl controllers.UserControllerInterface) {
// 	UserRoutes := r.Group("/user")
// 	{
// 		UserRoutes.GET("/", controllers.WelcomeMessage)
// 		UserRoutes.GET("/users", controllers.GetAllUser)
// 		UserRoutes.GET("/usersnew", controllers.GetUserFavFood)
// 		UserRoutes.GET("/users/:id", controllers.GetUserById)
// 		UserRoutes.POST("/create", controllers.CreateUser)
// 		UserRoutes.PUT("/update/:id", controllers.UpdateUserById)
// 		UserRoutes.DELETE("/delete/:id", controllers.DeleteUserById)
// 	}
// }
