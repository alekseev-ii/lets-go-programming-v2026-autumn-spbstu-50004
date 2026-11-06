package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ivan.alekseev/task-3/internal/config"
	"github.com/ivan.alekseev/task-3/internal/valcurs"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")

	flag.Parse()

	cfg, err := config.NewConfig(*configPath)
	if err != nil {
		panic(fmt.Errorf("load config: %w", err))
	}

	inputFile, err := cfg.OpenInputFile()
	if err != nil {
		panic(fmt.Errorf("open input: %w", err))
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			panic(fmt.Errorf("close input: %w", err))
		}
	}(inputFile)

	valCurses, err := valcurs.NewValCurs(inputFile)
	if err != nil {
		panic(fmt.Errorf("read valcurs: %w", err))
	}

	err = valCurses.ParseValues()
	if err != nil {
		panic(fmt.Errorf("parse valcurs: %w", err))
	}

	valCurses.SortByGreaterValue()

	err = valCurses.Save(cfg.OutputFile)
	if err != nil {
		panic(fmt.Errorf("save valcurs: %w", err))
	}
}
