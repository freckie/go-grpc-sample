package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"

	"frec.kr/tdoo/config"
)

type Config struct {
	Addr string `yaml:"addr"`

	DB     DbConfig            `yaml:"db"`
	Logger config.LoggerConfig `yaml:"logger"`
}

func LoadConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	v := &Config{}
	if err := yaml.NewDecoder(f).Decode(v); err != nil {
		return nil, err
	}
	return v, nil
}

func (c *Config) Evaluate() error {
	return errors.Join(
		c.DB.Evaluate(),
		c.Logger.Evaluate(),
	)
}
