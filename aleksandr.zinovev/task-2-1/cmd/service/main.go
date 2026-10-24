package main

import "fmt"

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
)

type conditioner struct {
	min int
	max int
}

func newConditioner() *conditioner {
	return &conditioner{
		min: defaultMinTemp,
		max: defaultMaxTemp,
	}
}

func (c *conditioner) regulate(sign string, targetTemp int) int {
	if sign == ">=" && targetTemp > c.min {
		c.min = targetTemp
	} else if sign == "<=" && targetTemp < c.max {
		c.max = targetTemp
	}

	if c.min > c.max {
		return -1
	}

	return c.min
}

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Invalid departments number")

		return
	}

	for range departmentsCount {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Invalid employees number")

			return
		}

		cond := newConditioner()

		for range employeesCount {
			var (
				sign       string
				targetTemp int
			)

			if _, err := fmt.Scan(&sign, &targetTemp); err != nil {
				fmt.Println("Invalid temperature input")

				return
			}

			fmt.Println(cond.regulate(sign, targetTemp))
		}
	}
}
