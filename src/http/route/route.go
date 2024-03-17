package route

import (
	"os"

	"github.com/abdiltegar/image-processing/src/http/controller"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App             *fiber.App
	ImageController *controller.ImageController
}

func (c *RouteConfig) Setup() {
	routeV1 := c.App.Group("/api/v1")
	routeV1.Post("/convert", c.ImageController.Convert)
	routeV1.Post("/resize", c.ImageController.Resize)
	routeV1.Post("/compress", c.ImageController.Compress)

	c.App.Static("/result/", os.TempDir())
}
