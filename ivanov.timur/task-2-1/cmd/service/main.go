package main

import (
	"fmt"
)

const (
	cMinTemp = 15
	cMaxTemp = 30
)

type badTempError string

func (e badTempError) Error() string {
	return string(e)
}

type tempReg struct {
	minTemp    int
	maxTemp    int
	errBadTemp badTempError
}

func newTempReg(minTemp int, maxTemp int) *tempReg {
	return &tempReg{
		minTemp:    minTemp,
		maxTemp:    maxTemp,
		errBadTemp: badTempError("-1"),
	}
}

func (reg *tempReg) setTemp(pointer string, temp int) (int, error) {
	switch pointer {
	case ">=":
		if temp > reg.minTemp {
			reg.minTemp = temp
		}
	case "<=":
		if temp < reg.maxTemp {
			reg.maxTemp = temp
		}
	default:
		fmt.Println("Wrong input")
	}

	if reg.minTemp > reg.maxTemp {
		return 0, reg.errBadTemp
	}

	return reg.minTemp, nil
}

func main() {
	var depCount int

	_, err := fmt.Scan(&depCount)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	for range depCount {
		var workCount int

		_, err := fmt.Scan(&workCount)
		if err != nil {
			fmt.Println("Error:", err)

			return
		}

		regulator := newTempReg(cMinTemp, cMaxTemp)

		for range workCount {
			var (
				pointer string
				temp    int
			)

			_, err := fmt.Scan(&pointer, &temp)
			if err != nil {
				fmt.Println("Error:", err)

				return
			}

			out, err := regulator.setTemp(pointer, temp)
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(out)
		}
	}
}
