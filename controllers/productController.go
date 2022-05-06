package controllers

import (
	"api/database"
	"api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(page,&models.Product{}, database.DB))
}

func CreateProduct(c *fiber.Ctx)error{
	var product models.Product

	if err := c.BodyParser(&product); err != nil{
		return err
	}
	
	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx)error{
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx)error{
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil{
		return err
	}


	database.DB.Model(&product).Updates(product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx)error{
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)

	return c.JSON(fiber.Map{
		"message": "Product successfully deleted",
		"deletedProduct": product,
	})
}

