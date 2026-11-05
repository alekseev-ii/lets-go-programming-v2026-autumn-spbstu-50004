package valCurs

import (
	"fmt"
	"strconv"
	"strings"
)

type Valute struct {
	NumCode    string  `xml:"NumCode" json:"num_code"`
	CharCode   string  `xml:"CharCode" json:"char_code"`
	Value      string  `xml:"Value" json:"-"`
	ValueFloat float64 `xml:"-" json:"value"`
}

func ParseValue(sValue string) (float64, error) {
	normalized := strings.ReplaceAll(sValue, ",", ".")
	return strconv.ParseFloat(normalized, 64)
}

func (v *Valute) ParseValue() error {
	vFloat64, err := ParseValue(v.Value)
	if err != nil {
		return fmt.Errorf("error parsing value (%s): %w", v.Value, err)
	}
	v.ValueFloat = vFloat64
	return nil
}
