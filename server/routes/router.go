package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamtonmoy0/go-react-blog-app/controller"
)

func BlogRouter(app *fiber.App) {
	app.Get("/all-blog", controller.GetAllBlog)
	app.Post("/create-blog", controller.CreateBlog)
	app.Put("/update-blog", controller.UpdateBlog)
	app.Delete("/delete-blog", controller.DeleteBlog)
}
