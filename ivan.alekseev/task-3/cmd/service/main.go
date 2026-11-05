package main

import (
	"flag"
	"fmt"

	"github.com/ivan.alekseev/task-3/internal/config"
	"github.com/ivan.alekseev/task-3/internal/valCurs"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")

	flag.Parse()

	cfg, err := config.NewConfig(*configPath)
	if err != nil {
		fmt.Println(err)

		return
	}

	inputFile, err := cfg.OpenInputFile()
	if err != nil {
		fmt.Println(err)

		return
	}
	defer inputFile.Close()

	valCurses, err := valCurs.NewValCurs(inputFile)
	if err != nil {
		fmt.Println(err)

		return
	}

	err = valCurses.ParseValues()
	if err != nil {
		fmt.Println(err)

		return
	}

	valCurses.SortByGreaterValue()
	err = valCurses.Save(cfg.OutputFile)
	if err != nil {
		fmt.Println(err)

		return
	}

}
