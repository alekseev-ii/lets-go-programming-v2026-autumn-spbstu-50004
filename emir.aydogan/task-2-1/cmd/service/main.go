package main

import (
	"errors"
	"fmt"
)

const (
	minAllowedTemp  = 15
	maxAllowedTemp  = 30
	unavailableTemp = -1

	greaterOrEqual = ">="
	lessOrEqual    = "<="
)

var errInvalidOperator = errors.New("invalid operator")

type temperatureRange struct {
	minTemp int
	maxTemp int
}

func newTemperatureRange() *temperatureRange {
	return &temperatureRange{
		minTemp: minAllowedTemp,
		maxTemp: maxAllowedTemp,
	}
}

func (r *temperatureRange) update(operator string, temp int) error {
	switch operator {
	case greaterOrEqual:
		if temp > r.minTemp {
			r.minTemp = temp
		}
	case lessOrEqual:
		if temp < r.maxTemp {
			r.maxTemp = temp
		}
	default:
		return fmt.Errorf("%w: %q", errInvalidOperator, operator)
	}

	return nil
}

func (r *temperatureRange) optimal() int {
	if r.minTemp > r.maxTemp {
		return unavailableTemp
	}

	return r.minTemp
}

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Error: invalid departments count:", err)

		return
	}

	for range departmentsCount {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Error: invalid employees count:", err)

			return
		}

		tempRange := newTemperatureRange()

		for range employeesCount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println("Error: invalid temperature request:", err)

				return
			}

			if err := tempRange.update(operator, temp); err != nil {
				fmt.Println("Error:", err)

				return
			}

			fmt.Println(tempRange.optimal())
		}
	}
}
