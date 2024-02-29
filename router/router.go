package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gogit/gogit"
	"github.com/gogit/services"
)

func DefaultRouter(app *fiber.App, subject string, repositoryURL string, destinationDir string, serviceType string) {
	app.Post(subject, func(c *fiber.Ctx) error {
		var payload map[string]interface{}
		if err := c.BodyParser(&payload); err != nil {
			fmt.Println("Error parsing JSON payload:", err)
			return c.SendStatus(fiber.StatusBadRequest)
		}
		ref, ok := payload["ref"].(string)
		if !ok || ref != "refs/heads/main" {
			fmt.Println("Received event, but not the main branch push event")
			return c.SendStatus(fiber.StatusOK)
		}
		gogit.Gogit(repositoryURL, destinationDir)
		if serviceType == "go" {
			services.GoBuild(destinationDir, subject)
		}
		services.Service(subject)
		fmt.Println("Received push event to the main branch")
		return c.SendStatus(fiber.StatusOK)
	})
}
