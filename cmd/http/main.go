package main

import (
	"github.com/dafailyasa/event-micro/pkg/factories"
	"github.com/dafailyasa/event-micro/pkg/server"
)

func main() {
	factory := factories.NewFactory(
		"logs/log.csv",
		"./config/env.json",
	)

	factory.InitalizeValidator()
	configurator := factory.InitializeConfigurator()
	factory.InitializeLogger()
	factory.InitializeMongoDB()

	server := server.NewServer(configurator)

	config, err := configurator.GetConfig()
	if err != nil {
		panic(err)
	}

	err = server.Run(config.Server.Port, config.App.Name)
	if err != nil {
		panic(err)
	}
}
