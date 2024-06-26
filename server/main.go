package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/iamtonmoy0/go-react-blog-app/database"
	"github.com/iamtonmoy0/go-react-blog-app/routes"
)

func main() {
	db := database.Database()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	// Check if the connection is successful by performing a simple query
	var result int
	db.Raw("SELECT 1").Scan(&result)
	fmt.Println("Database connection successful!")

	app := fiber.New()
	// logger
	app.Use(logger.New())
	// cors
	app.Use(cors.New())
	// Setup routes
	routes.BlogRouter(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
	// Start the server
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	} else {
		fmt.Println("Server started on port 3000")
	}

}
