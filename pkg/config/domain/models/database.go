package models

import customErr "github.com/dafailyasa/event-micro/pkg/custom-error"

type Database struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	PoolMin int64  `json:"pooMin"`
	PoolMax int64  `json:"poolMax"`
	Timeout int64  `json:"timeout"`
}

func (d *Database) Validate() error {
	if d.Name == "" {
		return customErr.ErrDatabaseNameRequired
	}
	if d.URL == "" {
		return customErr.ErrDatabaseURLRequired
	}
	return nil
}
