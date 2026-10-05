package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
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
	NumCode    string  `xml:"NumCode"`
	CharCode   string  `xml:"CharCode"`
	Value      string  `xml:"Value"`
	ValueFloat float64 `xml:"-"`
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

func parseValue(s string) (float64, error) {
	normalized := strings.ReplaceAll(s, ",", ".")
	return strconv.ParseFloat(normalized, 64)
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

	for i := range valCurs.Valutes {
		sVal := valCurs.Valutes[i].Value
		val, err := parseValue(sVal)
		if err != nil {
			panic(fmt.Sprintf("error parsing valute value (%s): %s", sVal, err))
		}
		valCurs.Valutes[i].ValueFloat = val
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].ValueFloat > valCurs.Valutes[j].ValueFloat
	})

	for _, val := range valCurs.Valutes {
		fmt.Println(val.NumCode, val.CharCode, val.Value)
	}
}
