package controller

import (
	"github.com/gofiber/fiber/v2"
)

// get all blot
func GetAllBlog(c *fiber.Ctx) error {
	context := fiber.Map{"statusText": "ok", "msg": "list"}

	c.Status(200)
	return c.JSON(context)
}

// create blog
func CreateBlog(c *fiber.Ctx) error {
	context := fiber.Map{"statusText": "ok", "msg": "list"}

	c.Status(200)
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
