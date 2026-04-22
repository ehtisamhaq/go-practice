package main

import (
	"go-crud-fiber/internal/platform"
	"go-crud-fiber/internal/routes"
	"go-crud-fiber/internal/user"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "go-crud-fiber"})

	// 1. Connect to the database (e.g., PostgreSQL, MySQL)
	db := platform.InitDB()

	// 2. Migrate the database schema (optional, but recommended)
	db.AutoMigrate(&user.RegisterUserDTO{}) // Uncomment when you have a User model

	// 3. Load Routes
	routes.SetupRoutes(app, db)

	log.Fatal(app.Listen(":8080"))
}
