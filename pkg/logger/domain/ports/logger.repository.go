package ports

import "github.com/dafailyasa/event-micro/pkg/logger/domain/models"

type LoggerRepository interface {
	Save(log *models.Log) error
}
