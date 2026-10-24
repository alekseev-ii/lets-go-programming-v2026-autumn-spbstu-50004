package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	defaultMinTemp     = 15
	defaultMaxTemp     = 30
	expectedPartsCount = 2
)

type TemperatureController struct {
	minTemp int
	maxTemp int
}

func NewTemperatureController() *TemperatureController {
	return &TemperatureController{
		minTemp: defaultMinTemp,
		maxTemp: defaultMaxTemp,
	}
}

func (tc *TemperatureController) ApplyConstraint(operator string, desiredTemp int) {
	if operator == ">=" && desiredTemp > tc.minTemp {
		tc.minTemp = desiredTemp
	} else if operator == "<=" && desiredTemp < tc.maxTemp {
		tc.maxTemp = desiredTemp
	}
}

func (tc *TemperatureController) GetOptimal() int {
	if tc.minTemp <= tc.maxTemp {
		return tc.minTemp
	}

	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	departmentsCount, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	for range departmentsCount {
		if !scanner.Scan() {
			return
		}

		employeesCount, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		controller := NewTemperatureController()

		for range employeesCount {
			if !scanner.Scan() {
				return
			}

			inputParts := strings.Fields(scanner.Text())
			if len(inputParts) != expectedPartsCount {
				continue
			}

			operator := inputParts[0]
			desiredTemp, _ := strconv.Atoi(inputParts[1])

			controller.ApplyConstraint(operator, desiredTemp)
			fmt.Println(controller.GetOptimal())
		}
	}
}
