package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func NewConfig(configPath string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(configPath)

	if err != nil {

		return cfg, err
	}

	err = yaml.Unmarshal(data, &cfg)

	return cfg, err
}

func (c *Config) OpenInputFile() (*os.File, error) {

	return os.Open(c.InputFile)
}

func (c *Config) OpenOutputFile() (*os.File, error) {

	return os.Open(c.OutputFile)
}
