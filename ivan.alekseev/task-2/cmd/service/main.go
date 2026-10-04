package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"encoding/xml"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(charset) {
	case "windows-1251", "cp1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	case "utf-8", "ascii":
		return input, nil
	default:
		return nil, fmt.Errorf("unknown charset: %s", charset)
	}
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

	f, err := os.Open(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("error opening input file: %s", err))
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	decoder.CharsetReader = charsetReader
	var valCurs ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(fmt.Sprintf("error parsing config file: %s", err))
	}

	fmt.Println(len(valCurs.Valutes))
	fmt.Println(valCurs.Valutes[0].NumCode)
	fmt.Println(valCurs.Valutes[0].CharCode)
	fmt.Println(valCurs.Valutes[0].Value)
}
