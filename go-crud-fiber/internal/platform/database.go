package platform

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	DATABASE_URL := "postgresql://neondb_owner:npg_oyv9pcqLAQE3@ep-steep-frog-aohem6dg-pooler.c-2.ap-southeast-1.aws.neon.tech/neondb?channel_binding=require&sslmode=require"

	db, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{PrepareStmt: true}) // Recommended for performance with poolers

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Adjusting Connection Pool settings for Serverless environments like Neon
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}
