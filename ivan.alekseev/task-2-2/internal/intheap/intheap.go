package intheap

import "errors"

var ErrNotInteger = errors.New("not an integer")

func ToInt(x any) (int, error) {
	v, ok := x.(int)
	if !ok {
		return 0, ErrNotInteger
	}

	return v, nil
}

type IntHeap []int

func (h *IntHeap) Len() int { return len(*h) }

func (h *IntHeap) inRange(i, j int) bool {
	return (i >= 0 || i < len(*h)) || (j >= 0 || j < len(*h))
}

func (h *IntHeap) Less(i, j int) bool {
	if !h.inRange(i, j) {
		return false
	}

	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if !h.inRange(i, j) {
		return
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	v, err := ToInt(x)
	if err != nil {
		return
	}

	*h = append(*h, v)
}

func (h *IntHeap) Pop() any {
	if len(*h) == 0 {
		return nil
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
