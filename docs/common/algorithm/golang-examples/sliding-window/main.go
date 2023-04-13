package main

import "fmt"

func slidingWindow(arr []int, k int) []int {
	n := len(arr)
	if k > n {
		return nil
	}

	result := make([]int, n-k+1)
	windowSum := 0

	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	result[0] = windowSum

	for i := k; i < n; i++ {
		windowSum += arr[i] - arr[i-k]
		result[i-k+1] = windowSum
	}

	return result
}

func main() {
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := slidingWindow(arr, k)
	fmt.Println("Sliding window result:", result)
}
