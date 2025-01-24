package api_routes

import (
	api_handlers "github.com/MarcelArt/api-portfolio-marcel/handlers/api"
	"github.com/MarcelArt/api-portfolio-marcel/services"
	"github.com/gofiber/fiber/v2"
)

func SetupPortfolioRoutes(api fiber.Router) {
	h := api_handlers.NewPortfolioHandler(services.NewMailService())

	g := api.Group("/portfolio")
	g.Post("/email", h.SendMessage)
}
