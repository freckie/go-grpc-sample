package config

import (
	"fmt"
	"strings"
)

type DbConfig struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

func (c *DbConfig) Evaluate() error {
	c.Driver = strings.ToLower(c.Driver)
	switch c.Driver {
	case "sqlite3":
	case "mysql":
	case "postgres":
	case "":
		return fmt.Errorf("db driver not specified")
	default:
		return fmt.Errorf("db driver type [%s] is not supported", c.Driver)
	}

	if c.Source == "" {
		return fmt.Errorf("db source not specified")
	}
	return nil
}
