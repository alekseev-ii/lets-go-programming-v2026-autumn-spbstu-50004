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
		panic("no such file or directory " + *configPath)
	}

	inputFile, err := cfg.OpenInputFile()

	if err != nil {
		panic("no such file or directory " + cfg.InputFile)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			panic("error closing file " + cfg.InputFile)
		}
	}(inputFile)

	valCurses, err := valcurs.NewValCurs(inputFile)

	if err != nil {
		fmt.Printf("Error parsing ValCurs: %s\n", err)

		return
	}

	err = valCurses.ParseValues()

	if err != nil {
		fmt.Printf("Error parsing values: %s\n", err)

		return
	}

	valCurses.SortByGreaterValue()

	err = valCurses.Save(cfg.OutputFile)

	if err != nil {
		fmt.Printf("Error saving result: %s\n", err)

		return
	}

}
