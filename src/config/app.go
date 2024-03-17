package config

import (
	"github.com/abdiltegar/image-processing/src/http/controller"
	"github.com/abdiltegar/image-processing/src/http/route"
	"github.com/abdiltegar/image-processing/src/repository"
	"github.com/abdiltegar/image-processing/src/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	App    *fiber.App
	Config *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	//setup repositories
	imageRepository := repository.NewImageRepository()

	//setup use cases
	imageUseCase := usecase.NewImageUseCase(imageRepository)

	//setup controllers
	imageController := controller.NewImageController(imageUseCase)

	routeConfig := route.RouteConfig{
		App:             config.App,
		ImageController: imageController,
	}
	routeConfig.Setup()
}
