package main

import "fmt"

func main() {
	var a, b int
	var operator string
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&operator)
	switch operator {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(float64(a) / float64(b))
		}
	}
}
