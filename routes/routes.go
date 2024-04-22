package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	productService "github.com/davidgrldo/go-restapi-fiber/service/product"
)

func InitRoutes(app *fiber.App, db *gorm.DB) {
	// Create a new product
	app.Post("/products", func(c *fiber.Ctx) error {
		return productService.CreateProduct(db, c)
	})

	// Get all products
	app.Get("/products", func(c *fiber.Ctx) error {
		return productService.GetAllProducts(db, c)
	})

	// Get a product by ID
	app.Get("/products/:id", func(c *fiber.Ctx) error {
		return productService.GetProductByID(db, c)
	})

	// Update a product by ID
	app.Put("/products/:id", func(c *fiber.Ctx) error {
		return productService.UpdateProductByID(db, c)
	})

	// Delete a product by ID
	app.Delete("/products/:id", func(c *fiber.Ctx) error {
		return productService.DeleteProductByID(db, c)
	})
}
