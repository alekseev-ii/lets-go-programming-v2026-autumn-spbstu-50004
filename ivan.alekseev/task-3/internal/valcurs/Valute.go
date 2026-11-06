package valcurs

import (
	"fmt"
	"strconv"
	"strings"
)

type Valute struct {
	NumCode    string  `json:"num_code" xml:"NumCode"`
	CharCode   string  `json:"char_code" xml:"CharCode"`
	Value      string  `json:"-" xml:"Value"`
	ValueFloat float64 `json:"value" xml:"-"`
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
