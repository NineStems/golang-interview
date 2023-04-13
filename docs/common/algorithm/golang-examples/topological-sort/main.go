package main

import "fmt"

func topoSort(graph map[int][]int) []int {
	inDegree := make(map[int]int)
	for node := range graph {
		inDegree[node] = 0
	}
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}
	queue := []int{}
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}
	result := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}

func main() {
	graph := map[int][]int{
		1: {2, 3},
		2: {4},
		3: {4, 5},
		4: {6},
		5: {6},
		6: {},
	}
	sorted := topoSort(graph)
	fmt.Println("Topologically sorted nodes:", sorted)
}
