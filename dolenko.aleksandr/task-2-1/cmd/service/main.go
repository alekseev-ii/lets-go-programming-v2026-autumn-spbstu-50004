package main

import (
	"errors"
	"fmt"
)

const (
	MaxTemp = 30
	MinTemp = 15
	lte     = "<="
	gte     = ">="
	NoTemp  = -1
)

var (
	ErrDepartmentNumber = errors.New("invalid department number")
	ErrEmployeesNumber  = errors.New("invalid employees number")
	ErrTemperatureInput = errors.New("invalid temperature input")
)

type conditioner struct {
	low  int
	high int
}

func newConditioner() *conditioner {
	return &conditioner{high: MaxTemp, low: MinTemp}
}

func (c *conditioner) regulate(operator string, temp int) int {
	switch operator {
	case lte:
		if temp < c.high {
			c.high = temp
		}
	case gte:
		if temp > c.low {
			c.low = temp
		}
	}

	if c.low > c.high {
		return NoTemp
	}

	return c.low
}

func main() {
	var departamentsCount int

	_, err := fmt.Scan(&departamentsCount)
	if err != nil {
		fmt.Println(ErrDepartmentNumber)

		return
	}

	for range departamentsCount {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println(ErrEmployeesNumber)

			return
		}

		curCond := newConditioner()

		for range employeesCount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println(ErrTemperatureInput)

				return
			}

			fmt.Println(curCond.regulate(operator, temp))
		}
	}
}
