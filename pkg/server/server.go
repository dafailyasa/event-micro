package server

import (
	"github.com/dafailyasa/event-micro/pkg/config/domain/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Server struct {
	configurator ports.ConfigApplication
}

func NewServer(config ports.ConfigApplication) *Server {
	return &Server{
		configurator: config,
	}
}

func (s *Server) Run(port string, appName string) error {
	cfg, err := s.configurator.GetConfig()
	if err != nil {
		return err
	}

	app := fiber.New(fiber.Config{
		AppName:           appName,
		ReduceMemoryUsage: true,
		EnablePrintRoutes: true,
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: cfg.App.Name + " Metrics"}))

	err = app.Listen(":" + port)
	if err != nil {
		return err
	}
	return nil
}
