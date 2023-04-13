package main

import "fmt"

func countingSort(arr []int, maxVal int) []int {
	counts := make([]int, maxVal+1)
	sorted := make([]int, len(arr))

	for _, val := range arr {
		counts[val]++
	}

	for i := 1; i <= maxVal; i++ {
		counts[i] += counts[i-1]
	}

	for _, val := range arr {
		sorted[counts[val]-1] = val
		counts[val]--
	}

	return sorted
}

func main() {
	arr := []int{4, 2, 1, 5, 3}
	maxVal := 5
	sorted := countingSort(arr, maxVal)
	fmt.Println("Sorted array:", sorted)
}
