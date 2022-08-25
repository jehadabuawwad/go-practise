package routes

import (
	"theSiS/web-service-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("api/v1/products", controllers.CreateProduct)
	app.Get("api/v1/products", controllers.GetProducts)
	app.Put("api/v1/products/:productId", controllers.UpdateProduct)
	app.Delete("api/v1/products/:productId", controllers.DeleteProduct)
}
