package cmd

import (
	"fmt"
	"log"

	"github.com/MarcelArt/api-portfolio-marcel/config"
	"github.com/MarcelArt/api-portfolio-marcel/database"
	"github.com/MarcelArt/api-portfolio-marcel/routes"
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	config.SetupENV()
	database.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Printf("Listening to http://localhost:%s", config.Env.PORT)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.PORT)))
}
