package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"github.com/davidgrldo/go-restapi-fiber/database"
	"github.com/davidgrldo/go-restapi-fiber/routes"
)

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Initialize database connection
	db := database.InitDatabase()

	// Defer closing the database connection
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error getting underlying DB connection: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	// Define routes
	routes.InitRoutes(app, db)

	// Start the Fiber server
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server gracefully stopped")
}
