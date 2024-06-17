package main

import (
	"parking-management/config"
	"parking-management/routes"
	"parking-management/utils"
	"parking-management/database"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database
	database.InitDB()

	// Initialize Echo instance
	e := echo.New()

	// Initialize Validator
	utils.InitValidator()

	// Initialize routes
	routes.InitRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
