package routes

import (
	"go-crud-fiber/internal/user"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// 1. Global Prefix for API versioning
	api := app.Group("api/v1")

	// 2. Initialize the User Module via Wire
	// Note: InitializeUserModule is the function we defined in wire.go
	userHandler := user.InitUserModule(db)

	// 3. Register User Routes
	userGroup := api.Group("/users")
	userGroup.Post("/register", userHandler.Register)
}
