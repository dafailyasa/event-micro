package models

import customErr "github.com/dafailyasa/event-micro/pkg/custom-error"

type App struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Host        string `json:"host"`
	BaseURL     string `json:"base_url"`
	Version     string `json:"version"`
}

func (a *App) Validate() error {
	if a.Name == "" {
		return customErr.ErrAppNameRequired
	}
	if a.Description == "" {
		return customErr.ErrAppDescriptionRequired
	}
	if a.Host == "" {
		return customErr.ErrAppHostRequired
	}
	if a.BaseURL == "" {
		return customErr.ErrAppBaseURLRequired
	}
	if a.Version == "" {
		return customErr.ErrAppVersionRequired
	}
	return nil
}
