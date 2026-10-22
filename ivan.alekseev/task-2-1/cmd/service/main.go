package main

import (
	"errors"
	"fmt"
)

const (
	minTemperature = 15
	maxTemperature = 30
	noTemperature  = -1
)

var errBadTemperatureInput = errors.New("bad temperature input")

type conditioner struct {
	low  int
	high int
}

func (c *conditioner) setHigh(temp int) {
	c.high = min(c.high, temp)
}

func (c *conditioner) setLow(temp int) {
	c.low = max(c.low, temp)
}

func (c *conditioner) getTemp() int {
	if c.low > c.high {
		return noTemperature
	}

	return c.low
}

func (c *conditioner) parseCommand() error {
	var (
		dest string
		temp int
	)

	_, err := fmt.Scan(&dest, &temp)
	if err != nil {

		return fmt.Errorf("%w: %w", errBadTemperatureInput, err)
	}

	switch dest {
	case "<=":
		c.setHigh(temp)
	case ">=":
		c.setLow(temp)
	default:
		return errBadTemperatureInput
	}

	return nil
}

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Invalid department number")

		return
	}

	for range departmentCount {
		var employeeCount int

		_, err = fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Invalid employee number")

			return
		}

		currentConditioner := conditioner{minTemperature, maxTemperature}

		for range employeeCount {
			err = currentConditioner.parseCommand()
			if err != nil {
				fmt.Println(err.Error())

				return
			}

			fmt.Println(currentConditioner.getTemp())
		}
	}
}
