package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	value     int
	listIndex int
}

type Heap []*Element

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*Element))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists [][]int) []int {
	k := len(lists)
	h := &Heap{}
	heap.Init(h)

	for i := 0; i < k; i++ {
		if len(lists[i]) > 0 {
			element := &Element{value: lists[i][0], listIndex: i}
			heap.Push(h, element)
		}
	}

	result := []int{}
	for h.Len() > 0 {
		element := heap.Pop(h).(*Element)
		result = append(result, element.value)

		listIndex := element.listIndex
		lists[listIndex] = lists[listIndex][1:]
		if len(lists[listIndex]) > 0 {
			element = &Element{value: lists[listIndex][0], listIndex: listIndex}
			heap.Push(h, element)
		}
	}

	return result
}

func main() {
	lists := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	result := mergeKLists(lists)
	fmt.Println(result) // Output: [1 1 2 3 4 4 5 6]
}
