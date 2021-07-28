package appRoutes

import (
	"github.com/FrancoGroenewald/CP2/API/myAgents"
	"github.com/gofiber/fiber/v2"
)

func RoutePaths(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/login", myAgents.Login)
	app.Get("/myagents", myAgents.MyAgents)

}
