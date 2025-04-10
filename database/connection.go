package database

import (
	"log"
	"portfolio-backend/config"
	"portfolio-backend/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := config.EnvDBConnection()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database: ", err)
	}

	DB = db

	// Auto-migrate
	log.Println("Starting DB auto-migration...")

	err = DB.AutoMigrate(
		&models.Project{},
		&models.User{},
		&models.Contact{},
		&models.CaseStudy{},
		&models.ProductImpact{},
	)
	if err != nil {
		log.Fatal("Auto-migration failed: ", err)
	}

	log.Println("Auto-migration completed")

	// Dummy admin seeding
	var userCount int64
	DB.Model(&models.User{}).Count(&userCount)
	if userCount == 0 {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("admin123"), 12)
		DB.Create(&models.User{
			Name:     "Administrator",
			Email:    "admin@example.com",
			Password: string(hashed),
		})
		config.Logger.Info("✅ Admin user seeded")
	}

	config.Logger.Info("✅ Database connected and schema migrated")
}
