package main

import (
	"fmt"

	"github.com/FrancoGroenewald/CP2/API/appConfig"
	"github.com/FrancoGroenewald/CP2/API/appRoutes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("main")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	appRoutes.RoutePaths(app)
	appConfig.Connection()
	app.Listen(":3000")

}
