package main

import (
	"fmt"

	"github.com/abdiltegar/image-processing/src/config"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		App:    app,
		Config: viperConfig,
	})

	host := fmt.Sprintf("%s:%s", viperConfig.Get("SERVER_HOST"), viperConfig.Get("SERVER_PORT"))
	err := app.Listen(host)
	if err != nil {
		panic(err)
	}
}
