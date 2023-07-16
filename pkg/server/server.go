package server

import (
	"github.com/dafailyasa/event-micro/pkg/config/domain/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	eventHdl "github.com/dafailyasa/event-micro/internal/event/domain/ports"
)

type Server struct {
	configurator ports.ConfigApplication
	event        eventHdl.EventHandlers
}

func NewServer(config ports.ConfigApplication, eventHdl eventHdl.EventHandlers) *Server {
	return &Server{
		configurator: config,
		event:        eventHdl,
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

	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/metrics", monitor.New(monitor.Config{Title: cfg.App.Name + " Metrics"}))
	prefix := app.Group(cfg.App.BaseURL)

	event := prefix.Group("event")
	event.Post("", s.event.Create)
	event.Get("/:id", s.event.FindById)

	err = app.Listen(":" + port)
	if err != nil {
		return err
	}
	return nil
}
