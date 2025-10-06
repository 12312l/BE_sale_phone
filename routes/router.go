package routes

import (
	"github.com/gin-gonic/gin"
	"daoduy.com/hoc-golang/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetAllUsers)
		// userRoutes.GET("/:id", controllers.GetUserByID)
		userRoutes.POST("/", controllers.CreateUser)
		// userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	authRouters := r.Group("/auth") 
	{
		authRouters.POST("/login", controllers.Login)
	}

	return r
}
