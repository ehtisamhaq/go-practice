package main

import (
	"go-crud-fiber/internal/platform"
	"go-crud-fiber/internal/routes"
	"go-crud-fiber/internal/user"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "go-crud-fiber",
		ErrorHandler: platform.ErrorHandler,
	})

	// 1. Connect to the database
	db := platform.InitDB()

	// 2. Migrate the database schema
	db.AutoMigrate(&user.User{})

	// 3. Load Routes
	routes.SetupRoutes(app, db)

	log.Fatal(app.Listen(":8080"))
}
