package routes

import (
	"contacts/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetContacts)
	route.Post("", controllers.CreateContact)
	route.Put("/:id", controllers.UpdateContact)
	route.Delete("/:id", controllers.DeleteContact)
	route.Get("/:id", controllers.GetContact)
}
