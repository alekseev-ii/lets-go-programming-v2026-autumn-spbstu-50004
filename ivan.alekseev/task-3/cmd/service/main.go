package main

import (
	"flag"
	"fmt"

	"github.com/ivan.alekseev/task-3/internal/config"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")

	flag.Parse()

	cfg, err := config.NewConfig(*configPath)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(cfg.InputFile)
	fmt.Println(cfg.OutputFile)
}
