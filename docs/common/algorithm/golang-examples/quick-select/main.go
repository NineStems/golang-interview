package main

import "fmt"

func partition(arr []int, left, right, pivotIndex int) int {
	pivotValue := arr[pivotIndex]
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
	storeIndex := left
	for i := left; i < right; i++ {
		if arr[i] < pivotValue {
			arr[i], arr[storeIndex] = arr[storeIndex], arr[i]
			storeIndex++
		}
	}
	arr[right], arr[storeIndex] = arr[storeIndex], arr[right]
	return storeIndex
}

func quickselect(arr []int, left, right, k int) int {
	if left == right {
		return arr[left]
	}
	pivotIndex := (left + right) / 2
	pivotIndex = partition(arr, left, right, pivotIndex)
	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return quickselect(arr, left, pivotIndex-1, k)
	} else {
		return quickselect(arr, pivotIndex+1, right, k)
	}
}

func main() {
	arr := []int{3, 7, 1, 2, 8, 4, 5}
	k := 3
	fmt.Println("Original array:", arr)
	fmt.Printf("%d-th smallest element is %d\n", k, quickselect(arr, 0, len(arr)-1, k-1))
}
