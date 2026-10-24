package main

import "fmt"

const (
	defaultLow  = 15
	defaultHigh = 30
	signGte     = ">="
	signLte     = "<="
)

type climateControl struct {
	low  int
	high int
}

func newClimateControl() climateControl {
	return climateControl{low: defaultLow, high: defaultHigh}
}

func (c *climateControl) update(sign string, value int) int {
	switch sign {
	case signGte:
		c.low = max(c.low, value)
	case signLte:
		c.high = min(c.high, value)
	}

	if c.low > c.high {
		return -1
	}

	return c.low
}

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Error reading number of departments:", err)

		return
	}

	for range departments {
		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Println("Error reading number of employees:", err)

			return
		}

		ctrl := newClimateControl()

		for range employees {
			var (
				sign  string
				value int
			)

			if _, err := fmt.Scan(&sign, &value); err != nil {
				fmt.Println("Error reading temperature parametrs:", err)

				return
			}

			fmt.Println(ctrl.update(sign, value))
		}
	}
}
