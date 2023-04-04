package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	merged := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			merged[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			merged[k] = left[i]
			i++
		} else if left[i] < right[j] {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
	}
	return merged
}

func main() {
	arr := []int{9, 7, 5, 3, 1, 2, 4, 6, 8}
	fmt.Println("Unsorted Array:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted Array:", sortedArr)
}
