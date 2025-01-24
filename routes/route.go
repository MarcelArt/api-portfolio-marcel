package routes

import (
	"time"

	"github.com/MarcelArt/api-portfolio-marcel/config"
	view_handlers "github.com/MarcelArt/api-portfolio-marcel/handlers/view"
	api_routes "github.com/MarcelArt/api-portfolio-marcel/routes/api"
	view_routes "github.com/MarcelArt/api-portfolio-marcel/routes/view"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))

	app.Static("/scripts", "./public/static/scripts")

	app.Get("/", view_handlers.HelloWorldView)

	if config.Env.ServerENV != "prod" {
		view_routes.SetupDevToolsRoutes(app)
	}

	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))

	app.Get("/metrics", monitor.New())

	// authMiddleware := middlewares.NewAuthMiddleware(repositories.NewUserRepo(database.GetDB()))

	api := app.Group("/api")
	api_routes.SetupPortfolioRoutes(api)
	// api_routes.SetupUserRoutes(api, authMiddleware)
	// api_routes.SetupAuthorizedDeviceRoutes(api, authMiddleware)
}
