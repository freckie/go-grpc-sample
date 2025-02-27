package config

import (
	"fmt"
	"strings"
)

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

func DefaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		Level:  "info",
		Format: "text",
	}
}

func (c *LoggerConfig) Evaluate() error {
	c.Level = strings.ToLower(c.Level)
	switch c.Level {
	case "debug":
	case "info":
	case "warn":
	case "error":
	default:
		return fmt.Errorf("log level [%s] is not supported", c.Level)
	}

	c.Format = strings.ToLower(c.Format)
	switch c.Format {
	case "json":
	case "text":
	default:
		return fmt.Errorf("log format [%s] is not supported", c.Level)
	}
	return nil
}
