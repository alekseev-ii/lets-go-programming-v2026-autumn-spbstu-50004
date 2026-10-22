package main

import "fmt"

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
		return -1
	}
	return c.low
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

		currentConditioner := conditioner{15, 30}

		for range employeeCount {
			var (
				dest string
				val  int
			)

			_, err = fmt.Scan(&dest, &val)
			if err != nil {
				fmt.Println("Bad temperature input")

				return
			}

			switch dest {
			case "<=":
				currentConditioner.setHigh(val)
			case ">=":
				currentConditioner.setLow(val)
			default:
				fmt.Println("Bad temperature input")

				return
			}

			fmt.Println(currentConditioner.getTemp())
		}
	}
}
