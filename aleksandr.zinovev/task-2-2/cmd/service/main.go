package main

import (
	"container/heap"
	"fmt"

	"aleksandr.zinovev/task-2-2/internal/intheap"
)

func main() {
	var (
		dishesCount int
		targetIndex int
		preferences []int
	)

	_, err := fmt.Scan(&dishesCount)
	if err != nil {
		fmt.Println("Invalid dishes count")

		return
	}

	preferences = make([]int, dishesCount)
	for index := range dishesCount {
		_, err = fmt.Scan(&preferences[index])
		if err != nil {
			fmt.Println("Invalid preferences")

			return
		}
	}

	_, err = fmt.Scan(&targetIndex)
	if err != nil {
		fmt.Println("Invalid target index")

		return
	}

	minHeap := &intheap.IntHeap{}
	heap.Init(minHeap)

	for _, rating := range preferences {
		if minHeap.Len() < targetIndex {
			heap.Push(minHeap, rating)
		} else {
			smallestItem, ok := heap.Pop(minHeap).(int)
			if !ok {
				return
			}

			if rating > smallestItem {
				heap.Push(minHeap, rating)
			} else {
				heap.Push(minHeap, smallestItem)
			}
		}
	}

	if minHeap.Len() > 0 {
		result, ok := heap.Pop(minHeap).(int)
		if !ok {
			return
		}

		fmt.Println(result)
	}
}
