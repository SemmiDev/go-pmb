package routes

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AddHealthRouter(router fiber.Router) {
	router.Get("/health", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.Status(http.StatusOK).SendString("<b>I'm</b> Healthy")
	})
}
