package main

import (
	"container/heap"
	"fmt"

	"github.com/alekseev.ii/task-2-2/internal/intheap"
)

func main() {
	var (
		numberOfDishes  int
		currentRate     int
		targetRatePlace int
	)

	ratesHeap := &intheap.IntHeap{}

	if _, err := fmt.Scan(&numberOfDishes); err != nil {
		fmt.Println(err)

		return
	}

	for range numberOfDishes {
		if _, err := fmt.Scan(&currentRate); err != nil {
			fmt.Println(err)

			return
		}

		heap.Push(ratesHeap, currentRate)
	}

	if _, err := fmt.Scan(&targetRatePlace); err != nil {
		fmt.Println(err)

		return
	}

	if targetRatePlace < 1 || targetRatePlace > numberOfDishes {
		fmt.Println("Not enough dishes")

		return
	}

	for range targetRatePlace - 1 {
		heap.Pop(ratesHeap)
	}

	value, err := intheap.ToInt(heap.Pop(ratesHeap))
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(value)
}
