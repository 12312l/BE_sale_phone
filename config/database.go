package config

import (
	"fmt"
	"log"

	"daoduy.com/hoc-golang/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Nga31072004@@tcp(127.0.0.1:3306)/workmanagement?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Không thể kết nối database: ", err)
	}

	DB = database
	// Tự động tạo bảng nếu chưa có
	DB.AutoMigrate(
		&models.User{},
		&models.InvalidatedToken{},
		&models.Address{},
		&models.Category{},
		&models.Product{},
		&models.ProductVariant{},
		&models.Image{},
		&models.Color{},
		&models.Sepecification{},
		&models.SepecificationDetail{},
	)
	createDefaultAdmin()

	fmt.Println("✅ Kết nối database thành công!")
}

// --- Thêm ngay trong file này ---
func createDefaultAdmin() {
	var count int64
	DB.Model(&models.User{}).Where("username = ?", "admin").Count(&count)

	if count == 0 {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		admin := models.User{
			Username: "admin",
			Password: string(hashed),
			Email:    "admin@example.com",
			Role:     "ADMIN",
			FullName: "Administrator",
		}
		if err := DB.Create(&admin).Error; err != nil {
			log.Println("❌ Lỗi khi tạo tài khoản admin mặc định:", err)
		} else {
			log.Println("✅ Đã tạo tài khoản admin mặc định (username: admin, password: admin)")
		}
	}
}
