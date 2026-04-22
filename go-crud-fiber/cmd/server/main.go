package main

import (
	"go-crud-fiber/internal/platform"
	"go-crud-fiber/internal/routes"
	"go-crud-fiber/internal/user"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "go-crud-fiber",
		ErrorHandler: platform.ErrorHandler,
	})

	// 1. Standard Middlewares
	app.Use(recover.New()) // Prevents server from crashing on panic
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	// 2. Connect to the database
	db := platform.InitDB()

	// 3. Migrate the database schema
	db.AutoMigrate(&user.User{})

	// 4. Load Routes
	routes.SetupRoutes(app, db)

	// 5. Graceful Shutdown Implementation
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Printf("Server failed to start: %v", err)
		}
	}()

	<-quit
	log.Println("Gracefully shutting down server...")

	sqlDB, _ := db.DB()
	if sqlDB != nil {
		sqlDB.Close()
	}

	if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
