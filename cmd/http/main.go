package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dafailyasa/event-micro/pkg/factories"
	"github.com/dafailyasa/event-micro/pkg/server"
)

func main() {
	_, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Failed to load asia jakarta timezone")
		os.Exit(1)
	}
	factory := factories.NewFactory(
		"./logs/log.csv",
		"./config/env.json",
	)

	factory.InitalizeValidator()
	configurator := factory.InitializeConfigurator()
	factory.InitializeLogger()
	factory.InitializeMongoDB()

	userHandlers := factory.BuildEventHandlers()

	server := server.NewServer(configurator, userHandlers)

	config, err := configurator.GetConfig()
	if err != nil {
		panic(err)
	}

	err = server.Run(config.Server.Port, config.App.Name)
	if err != nil {
		panic(err)
	}
}
