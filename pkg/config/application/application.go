package application

import (
	"github.com/dafailyasa/event-micro/pkg/config/domain/models"
	"github.com/dafailyasa/event-micro/pkg/config/domain/ports"
	logger "github.com/dafailyasa/event-micro/pkg/logger/domain/ports"
	validator "github.com/dafailyasa/event-micro/pkg/validator/domain/ports"
)

type ConfigService struct {
	repository    ports.ConfigRepository
	configuration *models.Config
	validator     validator.ValidatorApplication
	logger        logger.LoggerApplication
}

var _ ports.ConfigApplication = (*ConfigService)(nil)

func NewConfigService(
	repository ports.ConfigRepository,
	val validator.ValidatorApplication,
	logger logger.LoggerApplication,
) *ConfigService {
	return &ConfigService{
		repository: repository,
		validator:  val,
		logger:     logger,
	}
}
