package main

import (
	"container/heap"
	"fmt"

	"github.com/alekseev.ii/task-2-2/cmd/internal/intHeap"
)

func main() {
	var (
		numberOfDishes  int
		currentRate     int
		targetRatePlace int
	)

	ratesHeap := &intHeap.IntHeap{}

	_, err := fmt.Scan(&numberOfDishes)
	if err != nil {
		fmt.Println(err)

		return
	}

	for range numberOfDishes {
		_, err := fmt.Scan(&currentRate)
		if err != nil {
			fmt.Println(err)

			return
		}

		heap.Push(ratesHeap, currentRate)
	}

	_, err = fmt.Scan(&targetRatePlace)
	if err != nil {
		fmt.Println(err)

		return
	}

	for range targetRatePlace - 1 {
		heap.Pop(ratesHeap)
	}

	fmt.Println(intHeap.ToInt(heap.Pop(ratesHeap)))
}
