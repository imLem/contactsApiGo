package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "contacts/routes"
)

func welcome(c *fiber.Ctx) error {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "success":  true,
        "message": "You are at the endpoint 😉",
    })
}

func setupRoutes(app *fiber.App) {

    // moved from main method
    app.Get("/", welcome)

    api := app.Group("/api")

    api.Get("", func(c *fiber.Ctx) error {
      return c.Status(fiber.StatusOK).JSON(fiber.Map{
          "success": true,
          "message": "You are at the api endpoint 😉",
      })
  })

  routes.TodoRoute(api.Group("/contacts"))
}

func main() {
    app := fiber.New()
    app.Use(logger.New())

    // setup routes
    setupRoutes(app) // new

    // Listen on server 8000 and catch error if any
    err := app.Listen(":8000")

    // handle error
    if err != nil {
        panic(err)
    }
}
