package models

type Config struct {
	App      *App      `json:"app"`
	Database *Database `json:"database"`
	Server   *Server   `json:"server"`
}

func (c *Config) Validate() error {
	if err := c.Database.Validate(); err != nil {
		return err
	}

	if err := c.App.Validate(); err != nil {
		return err
	}

	if err := c.Server.Validate(); err != nil {
		return err
	}

	return nil
}
