package api_handlers

import (
	"fmt"

	"github.com/MarcelArt/api-portfolio-marcel/config"
	"github.com/MarcelArt/api-portfolio-marcel/models"
	"github.com/MarcelArt/api-portfolio-marcel/services"
	"github.com/MarcelArt/api-portfolio-marcel/utils"
	"github.com/gofiber/fiber/v2"
)

type IPortfolioHandler interface {
}

type PortfolioHandler struct {
	mService services.IMailService
}

func NewPortfolioHandler(mService services.IMailService) *PortfolioHandler {
	return &PortfolioHandler{
		mService: mService,
	}
}

// Send email message
// @Summary Send email message
// @Description Send email message
// @Tags Portfolio
// @Accept json
// @Produce json
// @Param User body models.PortfolioMessage true "Message"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /portfolio/email [post]
func (h *PortfolioHandler) SendMessage(c *fiber.Ctx) error {
	var message models.PortfolioMessage
	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	body := fmt.Sprintf(`
		From: %s &lt;%s&gt;
		<br/>
		<p>Message: %s</p>
	`, message.Name, message.Email, message.Message)

	err := h.mService.SendMail(services.Mailer{
		To:      []string{config.Env.SMTPEmail},
		Subject: message.Subject,
		Body:    body,
	})
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	return c.Status(fiber.StatusOK).JSON(models.NewJSONResponse(nil, "Message sent successfully"))
}
