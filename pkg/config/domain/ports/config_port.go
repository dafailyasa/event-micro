package ports

import (
	"os"

	"github.com/dafailyasa/event-micro/pkg/config/domain/models"
)

type ConfigApplication interface {
	Config() error
	GetConfig() (*models.Config, error)
}

type ConfigRepository interface {
	GetConfigFile() (*os.File, error)
}
