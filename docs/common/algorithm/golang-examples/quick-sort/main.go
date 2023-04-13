package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr) // выбор опорного элемента не обязательно случайность
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	return arr
}

func main() {
	arr := []int{9, 7, 5, 3, 1, 2, 4, 6, 8}
	fmt.Println("Unsorted Array:", arr)
	sortedArr := quickSort(arr)
	fmt.Println("Sorted Array:", sortedArr)
}
