package routes

import (
	"daoduy.com/hoc-golang/controllers"
	"daoduy.com/hoc-golang/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	authRouters := r.Group("/auth")
	{
		authRouters.POST("/login", controllers.Login)
		authRouters.Use(middleware.AuthMiddleware())
		{
			authRouters.POST("/logout", controllers.Logout)
		}
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controllers.CreateUser)

		userRoutes.Use(middleware.AuthMiddleware())
		{
			adminRoutes := userRoutes.Group("/")
			adminRoutes.Use(middleware.AuthMiddleware("ADMIN"))
			{
				adminRoutes.GET("/", controllers.GetAllUsers)
				adminRoutes.GET("/:id", controllers.GetUserById)
				// adminRoutes.DELETE("/:id", controllers.DeleteUser)
			}

			userRoutes.GET("/myinfo", controllers.MyInfoUser)
		}
	}

	addressRoutes := r.Group("/address")
	{
		addressRoutes.Use(middleware.AuthMiddleware())
		{
			addressRoutes.GET("/myaddress", controllers.MyAddress)
			addressRoutes.POST("/add", controllers.CreateAddress)
			addressRoutes.PUT("/update/:id", controllers.UpdateAddress)
		}
	}

	return r
}
