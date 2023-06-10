package main

import (
	"github.com/dafailyasa/event-micro/pkg/factories"
)

func main() {
	factory := factories.NewFactory(
		"logs/log.csv",
		"./config/env.json",
	)

	factory.InitalizeValidator()
	_ = factory.InitializeConfigurator()
	factory.InitializeLogger()
	factory.InitializeMongoDB()

}
