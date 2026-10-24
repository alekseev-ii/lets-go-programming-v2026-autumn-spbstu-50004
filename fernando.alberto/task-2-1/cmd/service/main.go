package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	cMinTemp = 15
	cMaxTemp = 30
)

var (
	errNoTemp      = errors.New("no suitable temperature")
	errBadOperator = errors.New("wrong operator")
)

type tempRegulator struct {
	minTemp int
	maxTemp int
}

func newTempRegulator(minTemp, maxTemp int) *tempRegulator {
	return &tempRegulator{minTemp: minTemp, maxTemp: maxTemp}
}

func (r *tempRegulator) setTemp(operator string, temp int) (int, error) {
	switch operator {
	case ">=":
		if temp > r.minTemp {
			r.minTemp = temp
		}
	case "<=":
		if temp < r.maxTemp {
			r.maxTemp = temp
		}
	default:
		return 0, errBadOperator
	}

	if r.minTemp > r.maxTemp {
		return 0, errNoTemp
	}

	return r.minTemp, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer func() {
		_ = writer.Flush()
	}()

	var depCount int
	if _, err := fmt.Fscan(reader, &depCount); err != nil {
		_, _ = fmt.Fprintln(writer, "Error:", err)

		return
	}

	for range depCount {
		var workCount int
		if _, err := fmt.Fscan(reader, &workCount); err != nil {
			_, _ = fmt.Fprintln(writer, "Error:", err)

			return
		}

		regulator := newTempRegulator(cMinTemp, cMaxTemp)

		for range workCount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Fscan(reader, &operator, &temp); err != nil {
				_, _ = fmt.Fprintln(writer, "Error:", err)

				return
			}

			result, err := regulator.setTemp(operator, temp)
			if errors.Is(err, errBadOperator) {
				_, _ = fmt.Fprintln(writer, "Wrong input")

				return
			}

			if err != nil {
				_, _ = fmt.Fprintln(writer, -1)

				continue
			}

			_, _ = fmt.Fprintln(writer, result)
		}
	}
}
