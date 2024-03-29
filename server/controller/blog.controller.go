package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtonmoy0/go-react-blog-app/database"
	"github.com/iamtonmoy0/go-react-blog-app/models"
)

// get all blot
func GetAllBlog(c *fiber.Ctx) error {
	db := database.Database()
	var record []models.Blog
	db.Find(&record)

	context := fiber.Map{"data": record}

	c.Status(200)
	return c.JSON(context)
}

// create blog
func CreateBlog(c *fiber.Ctx) error {

	db := database.Database()
	record := new(models.Blog)
	if err := c.BodyParser(&record); err != nil {
		fmt.Println("failed to get the body")
	}

	db.Create(record)

	context := fiber.Map{"msg": "successfully created the blog!"}

	c.Status(201)
	return c.JSON(context)
}

// update blog
func UpdateBlog(c *fiber.Ctx) error {
	context := fiber.Map{"statusText": "ok", "msg": "list"}

	c.Status(200)
	return c.JSON(context)
}

// delete blog
func DeleteBlog(c *fiber.Ctx) error {
	context := fiber.Map{"statusText": "ok", "msg": "list"}

	c.Status(200)
	return c.JSON(context)
}
