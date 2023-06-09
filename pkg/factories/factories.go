package factories

import (
	configApp "github.com/dafailyasa/event-micro/pkg/config/application"
	configRepo "github.com/dafailyasa/event-micro/pkg/infrastructure/repositories"
	loggerApp "github.com/dafailyasa/event-micro/pkg/logger/application"
	loggerRepo "github.com/dafailyasa/event-micro/pkg/logger/infrastructure/repositories"
	valApp "github.com/dafailyasa/event-micro/pkg/validator/application"
)

const MongoClientTimeout = 10

type Factory struct {
	//Variables
	logFilePath    string
	configFilePath string
	//Packages
	validator    *valApp.Validator
	configurator *configApp.ConfigService
	logger       loggerApp.Logger
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
		return app
	}
	return f.validator
}

func (f *Factory) InitializeLogger() *loggerApp.Logger {
	if f.configurator == nil {
		validator := f.InitalizeValidator()
		path := f.configFilePath

		repo := loggerRepo.NewCsvFile(path)
		app := loggerApp.NewLogger(repo, validator)

		f.logger = *app
	}

	return &f.logger
}

func (f *Factory) InitializeConfigurator() *configApp.ConfigService {
	if f.configurator == nil {
		validator := f.InitalizeValidator()
		logger := f.InitializeLogger()
		path := f.configFilePath

		repo := configRepo.NewJSONRepository(path)
		app := configApp.NewConfigService(repo, validator, logger)
		err := app.Config()
		if err != nil {
			panic(err)
		}
		f.configurator = app
		return app
	}
	return f.configurator
}
