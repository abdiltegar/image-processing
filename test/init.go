package test

import (
	"github.com/abdiltegar/image-processing/src/config"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var app *fiber.App

var viperConfig *viper.Viper

func init() {
	viperConfig = config.NewViper()
	app = config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		App:    app,
		Config: viperConfig,
	})
}
