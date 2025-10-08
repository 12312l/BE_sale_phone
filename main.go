package main

import (
	"daoduy.com/hoc-golang/config"
	_ "daoduy.com/hoc-golang/docs" // 👈 Quan trọng: import docs tự sinh bởi swag
	"daoduy.com/hoc-golang/middleware"
	"daoduy.com/hoc-golang/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Học Golang API
// @version 1.0
// @description Đây là tài liệu API cho project học Golang
// @termsOfService http://swagger.io/terms/

// @contact.name Duy
// @contact.url http://github.com/daoduy
// @contact.email duy@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath
func main() {
	// Kết nối DB
	config.ConnectDatabase()

	r := gin.Default()
	r.Use(middleware.GlobalExceptionHandler())

	// Khởi tạo router
	routes.SetupRouter(r)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
