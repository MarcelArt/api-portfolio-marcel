package view_handlers

import (
	"github.com/MarcelArt/api-portfolio-marcel/utils"
	"github.com/MarcelArt/api-portfolio-marcel/views/hello"
	"github.com/gofiber/fiber/v2"
)

func HelloWorldView(c *fiber.Ctx) error {
	return utils.Render(c, hello.Show("marcel"))
}
