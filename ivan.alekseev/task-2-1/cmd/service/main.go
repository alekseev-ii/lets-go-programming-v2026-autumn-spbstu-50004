package main

import "fmt"

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
		minTemperature, maxTemperature := 15, 30
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
				maxTemperature = min(maxTemperature, val)
			case ">=":
				minTemperature = max(minTemperature, val)
			default:
				fmt.Println("Bad temperature input")

				return
			}

			if maxTemperature < minTemperature {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemperature)
			}
		}
	}
}
