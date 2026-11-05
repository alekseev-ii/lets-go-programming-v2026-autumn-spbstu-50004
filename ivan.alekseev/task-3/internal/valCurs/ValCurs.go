package valCurs

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

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
		return nil, fmt.Errorf("unknown charset: %s", charset)
	}
}

func NewValCurs(sourceFile *os.File) (ValCurs, error) {
	decoder := xml.NewDecoder(sourceFile)
	decoder.CharsetReader = CharsetReader
	var valCurs ValCurs
	err := decoder.Decode(&valCurs)
	return valCurs, err
}

func (vc *ValCurs) ParseValues() error {
	for i := range vc.Valutes {
		err := vc.Valutes[i].ParseValue()
		if err != nil {
			return err
		}
	}
	return nil
}

func (vc *ValCurs) SortByGreaterValue() {
	sort.Slice(vc.Valutes, func(i, j int) bool {
		return vc.Valutes[i].ValueFloat > vc.Valutes[j].ValueFloat
	})
}
