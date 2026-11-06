package valcurs

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const (
	outputFilePerm = 0o600
	outputDirPerm  = 0o755
)

var errUnknownCharset = errors.New("unknown charset")

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

func CharsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(charset) {
	case "utf-8", "ascii":
		return input, nil
	case "windows-1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	default:
		return nil, errUnknownCharset
	}
}

func NewValCurs(sourceFile *os.File) (ValCurs, error) {
	decoder := xml.NewDecoder(sourceFile)

	decoder.CharsetReader = CharsetReader

	var valCurs ValCurs

	err := decoder.Decode(&valCurs)
	if err != nil {
		return valCurs, fmt.Errorf("decode file: %w", err)
	}

	return valCurs, nil
}

func (vc *ValCurs) ParseValues() error {
	for i := range vc.Valutes {
		err := vc.Valutes[i].ParseValue()
		if err != nil {
			return fmt.Errorf("parse value %s: %w", vc.Valutes[i].Value, err)
		}
	}

	return nil
}

func (vc *ValCurs) SortByGreaterValue() {
	sort.Slice(vc.Valutes, func(i, j int) bool {
		return vc.Valutes[i].ValueFloat > vc.Valutes[j].ValueFloat
	})
}

func (vc *ValCurs) Save(outputFile string) error {
	outDir := filepath.Dir(outputFile)

	err := os.MkdirAll(outDir, outputDirPerm)
	if err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	result, err := json.MarshalIndent(&vc.Valutes, "", "\t")
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	err = os.WriteFile(outputFile, result, outputFilePerm)
	if err != nil {
		return fmt.Errorf("write result: %w", err)
	}

	return nil
}
