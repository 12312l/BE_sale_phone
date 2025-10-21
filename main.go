package main

import (
	"time"

	"daoduy.com/hoc-golang/config"
	_ "daoduy.com/hoc-golang/docs" // 👈 Quan trọng: import docs tự sinh bởi swag
	"daoduy.com/hoc-golang/middleware"
	"daoduy.com/hoc-golang/routes"

	"github.com/gin-contrib/cors"
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
// @BasePath /

// ✅ Định nghĩa BearerAuth cho toàn bộ API
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Nhập token theo định dạng: Bearer <your_token>

// ✅ Áp dụng bảo mật mặc định (tự động cho tất cả API, trừ API login/register)
// @Security BearerAuth

func main() {
	// Kết nối DB
	config.ConnectDatabase()

	r := gin.Default()
	r.Use(middleware.GlobalExceptionHandler())

		// Cấu hình CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))


	// Khởi tạo router
	routes.SetupRouter(r)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
