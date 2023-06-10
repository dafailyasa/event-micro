package application

import (
	"encoding/json"

	"github.com/dafailyasa/event-micro/pkg/config/domain/models"
)

func (c *ConfigService) Config() error {
	configuration := new(models.Config)
	file, err := c.repository.GetConfigFile()
	if err != nill {
		return err
	}
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&configuration); err != nil {
		return err
	}
	c.configuration = configuration
	return nil
}
