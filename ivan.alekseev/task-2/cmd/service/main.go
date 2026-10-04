package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")

	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic(fmt.Sprintf("error reading config file: %s", err))
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(fmt.Sprintf("error parsing config file: %s", err))
	}

	fmt.Println(cfg.InputFile)
	fmt.Println(cfg.OutputFile)
}
