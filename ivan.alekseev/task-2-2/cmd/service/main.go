package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	var (
		numberOfDishes  int
		currentRate     int
		targetRatePlace int
	)
	ratesHeap := &IntHeap{}
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
	fmt.Println(heap.Pop(ratesHeap).(int))
}
