package models

import (
	"time"

	customErr "github.com/dafailyasa/event-micro/pkg/custom-error"
)

type Log struct {
	ID        string    `json:"id"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func (l Log) Validate() error {
	if l.Level == "" {
		return customErr.ErrLogLevelRequired
	}
	if l.Message == "" {
		return customErr.ErrLogMessageRequired
	}
	if l.CreatedAt.IsZero() {
		return customErr.ErrCreateAtRequired
	}
	return nil
}
