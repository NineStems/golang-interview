package main

type Heap []int

func (h *Heap) Push(x int) {
	*h = append(*h, x)
	i := len(*h) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if (*h)[i] <= (*h)[parent] {
			break
		}
		(*h)[i], (*h)[parent] = (*h)[parent], (*h)[i]
		i = parent
	}
}

func (h *Heap) Pop() int {
	n := len(*h)
	x := (*h)[0]
	(*h)[0] = (*h)[n-1]
	*h = (*h)[:n-1]
	i := 0
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < len(*h) && (*h)[left] > (*h)[largest] {
			largest = left
		}
		if right < len(*h) && (*h)[right] > (*h)[largest] {
			largest = right
		}
		if largest == i {
			break
		}
		(*h)[i], (*h)[largest] = (*h)[largest], (*h)[i]
		i = largest
	}
	return x
}
