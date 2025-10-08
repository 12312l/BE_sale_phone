package routes

import (
	"daoduy.com/hoc-golang/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetAllUsers)
		userRoutes.GET("/:id", controllers.GetUserById)
		userRoutes.POST("/", controllers.CreateUser)
		// userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	authRouters := r.Group("/auth")
	{
		authRouters.POST("/login", controllers.Login)
		authRouters.POST("/logout", controllers.Logout)

	}

	return r
}
