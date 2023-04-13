package main

import "fmt"

func twoPointer(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 9
	result := twoPointer(arr, target)
	fmt.Println("Two pointer result:", result)
}
