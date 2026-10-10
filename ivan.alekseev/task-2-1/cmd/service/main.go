package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid department number")
		return
	}

	for i := 0; i < n; i++ {
		var k int
		_, err = fmt.Scan(&k)
		if err != nil {
			fmt.Println("Invalid employee number")
			return
		}
		minT, maxT := 15, 30
		for j := 0; j < k; j++ {
			var (
				dest string
				val  int
			)
			_, err = fmt.Scan(&dest, &val)
			if err != nil {
				fmt.Println("Bad temperature input")
				return
			}
			if dest == "<=" {
				maxT = min(maxT, val)
			} else if dest == ">=" {
				minT = max(minT, val)
			} else {
				fmt.Println("Bad temperature input")
				return
			}
			if maxT < minT {
				fmt.Println(-1)
			} else {
				fmt.Println(minT)
			}
		}
	}
}
