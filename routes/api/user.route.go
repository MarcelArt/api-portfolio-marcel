package api_routes

import (
	"github.com/MarcelArt/api-portfolio-marcel/database"
	api_handlers "github.com/MarcelArt/api-portfolio-marcel/handlers/api"
	"github.com/MarcelArt/api-portfolio-marcel/middlewares"
	"github.com/MarcelArt/api-portfolio-marcel/repositories"
	"github.com/MarcelArt/api-portfolio-marcel/services"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router, auth *middlewares.AuthMiddleware) {
	h := api_handlers.NewUserHandler(
		repositories.NewUserRepo(database.GetDB()),
		repositories.NewAuthorizedDeviceRepo(database.GetDB()),
		services.NewMailService(),
	)

	g := api.Group("/user")
	g.Get("/", auth.ProtectedAPI, h.Read)
	g.Get("/:id", auth.ProtectedAPI, h.GetByID)
	g.Get("/verify/:id", h.Verify)

	g.Post("/", h.Create)
	g.Post("/login", h.Login)
	g.Post("/refresh", h.Refresh)

	g.Put("/:id", auth.ProtectedAPI, h.Update)
	g.Delete("/:id", auth.ProtectedAPI, h.Delete)
}
