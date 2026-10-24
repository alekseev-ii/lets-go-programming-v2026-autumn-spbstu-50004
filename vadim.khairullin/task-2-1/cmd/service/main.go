package main

import (
	"errors"
	"fmt"
)

var (
	errEmployees   = errors.New("wrong number of employees")
	errTemperature = errors.New("wrong format of temperature")
	errSign        = errors.New("wrong format of sign")
)

const (
	MIN = 15
	MAX = 30
)

type TemperatureRegime struct {
	minimum int
	maximum int
}

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil || departments < 1 || departments > 1000 {
		fmt.Println("Error: wrong number of Departments")

		return
	}

	for range departments {
		regime := TemperatureRegime{minimum: MIN, maximum: MAX}

		err = regime.setTemperatureByEmployees()
		if err != nil {
			fmt.Println(err)

			return
		}
	}
}

func (r *TemperatureRegime) setTemperatureByEmployees() error {
	var employees int

	_, err := fmt.Scan(&employees)
	if err != nil || employees < 1 || employees > 1000 {
		return errEmployees
	}

	for range employees {
		r.minimum, r.maximum, err = r.defineTemperatureRegime(r.minimum, r.maximum)
		if err != nil {
			return err
		}

		if r.minimum <= r.maximum {
			fmt.Println(r.minimum)
		} else {
			fmt.Println(-1)
		}
	}

	return nil
}

func (r *TemperatureRegime) defineTemperatureRegime(minTemp, maxTemp int) (int, int, error) {
	var (
		sign string
		temp int
	)

	_, err := fmt.Scan(&sign, &temp)
	if err != nil {
		return 0, 0, errTemperature
	}

	if temp < MIN || temp > MAX {
		return 0, 0, errTemperature
	}

	switch sign {
	case ">=":
		if minTemp < temp {
			minTemp = temp
		}
	case "<=":
		if maxTemp > temp {
			maxTemp = temp
		}
	default:
		return 0, 0, errSign
	}

	return minTemp, maxTemp, nil
}
