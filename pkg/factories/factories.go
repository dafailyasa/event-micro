package factories

import (
	configApp "github.com/dafailyasa/event-micro/pkg/config/application"
	valApp "github.com/dafailyasa/event-micro/pkg/validators/application"
)

const MongoClientTimeout = 10

type Factory struct {
	//Variables
	logFilePath    string
	configFilePath string
	//Packages
	validator    *valApp.Validator
	configurator *configApp.ConfigService
	logger       *loggerApp.Logger
	//	dbClient     *mongo.Client
}

func NewFactory(logFilePath string, configFilePath string) *Factory {
	return &Factory{
		logFilePath:    logFilePath,
		configFilePath: configFilePath,
	}
}

func (f *Factory) InitalizeValidator() *valApp.Validator {
	if f.validator == nil {
		app := valApp.NewValidator()
		f.validator = app
	}
	return f.validator
}

func (f *Factory) InitializeConfigurator() {

}
