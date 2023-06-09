package models

import customErr "github.com/dafailyasa/event-micro/pkg/custom-error"

type JWT struct {
	Secret    string `json:"secret"`
	ExpiresIn int    `json:"days_of_expiration"`
}

func (j *JWT) Validate() error {
	if j.Secret == "" {
		return customErr.ErrSecretRequired
	}
	if j.ExpiresIn == 0 {
		return customErr.ErrExpiresInRequired
	}
	return nil
}
