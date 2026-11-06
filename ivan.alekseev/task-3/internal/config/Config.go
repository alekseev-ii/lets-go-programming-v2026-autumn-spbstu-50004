package config

import (
	"fmt"
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
		return cfg, fmt.Errorf("open config file: %w", err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("parse config file: %w", err)
	}

	return cfg, nil
}

func (c *Config) OpenInputFile() (*os.File, error) {
	iFile, err := os.Open(c.InputFile)
	if err != nil {
		return iFile, fmt.Errorf("open input file: %w", err)
	}

	return iFile, nil
}

func (c *Config) OpenOutputFile() (*os.File, error) {
	oFile, err := os.Open(c.OutputFile)
	if err != nil {
		return oFile, fmt.Errorf("open output file: %w", err)
	}

	return oFile, nil
}
