package models

import customErr "github.com/dafailyasa/event-micro/pkg/custom-error"

type Server struct {
	Port string `json:"port"`
}

func (s *Server) Validate() error {
	if s.Port == "" {
		return customErr.ErrServerPortRequired
	}
	return nil
}
