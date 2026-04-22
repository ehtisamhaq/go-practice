package routes

import (
	"go-crud-fiber/internal/user"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// 1. Global Prefix for API versioning
	api := app.Group("api/v1")

	// 2. Manual Dependency Injection (Express/NestJS Style)
	userRepo := user.NewRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// 3. Register User Routes
	userGroup := api.Group("/users")
	user.RegisterRoutes(userGroup, userHandler)

}
