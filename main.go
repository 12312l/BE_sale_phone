package main

import (
	"daoduy.com/hoc-golang/config"
	"daoduy.com/hoc-golang/routes"
	_ "daoduy.com/hoc-golang/docs" // 👈 Quan trọng: import docs tự sinh bởi swag

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

	// Khởi tạo router
	r := routes.SetupRouter()

	// Swagger endpoint —> http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
