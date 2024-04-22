package product

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/davidgrldo/go-restapi-fiber/models"
)

func CreateProduct(db *gorm.DB, c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return err
	}
	db.Create(&product)
	return c.JSON(product)
}

func GetAllProducts(db *gorm.DB, c *fiber.Ctx) error {
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		return handleDBError(c, err)
	}

	if len(products) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No products found",
		})
	}

	return c.JSON(products)
}

func GetProductByID(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		return handleDBError(c, err)
	}
	return c.JSON(product)
}

func UpdateProductByID(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		return handleDBError(c, err)
	}

	newProduct := new(models.Product)
	if err := c.BodyParser(newProduct); err != nil {
		return err
	}

	product.Name = newProduct.Name
	product.Price = newProduct.Price
	db.Save(&product)

	return c.JSON(product)
}

func DeleteProductByID(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	if err := db.Delete(&models.Product{}, id).Error; err != nil {
		return handleDBError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "Product with id " + id + " deleted successfully",
	})
}

// Helper function to handle database errors
func handleDBError(c *fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal server error",
	})
}
