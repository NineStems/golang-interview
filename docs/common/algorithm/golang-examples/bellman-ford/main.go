package main

import (
	"fmt"
	"math"
)

type Edge struct {
	from   int
	to     int
	weight float64
}

func BellmanFord(edges []Edge, start int, numVertices int) ([]float64, []int) {
	distances := make([]float64, numVertices)
	previous := make([]int, numVertices)

	for i := 0; i < numVertices; i++ {
		distances[i] = math.Inf(1)
		previous[i] = -1
	}

	distances[start] = 0

	for i := 0; i < numVertices-1; i++ {
		for _, edge := range edges {
			if distances[edge.from]+edge.weight < distances[edge.to] {
				distances[edge.to] = distances[edge.from] + edge.weight
				previous[edge.to] = edge.from
			}
		}
	}

	// Check for negative cycles
	for _, edge := range edges {
		if distances[edge.from]+edge.weight < distances[edge.to] {
			panic("Negative cycle detected!")
		}
	}

	return distances, previous
}

func main() {
	edges := []Edge{
		{0, 1, 4},
		{0, 2, 5},
		{1, 2, -2},
		{1, 3, 6},
		{2, 3, 1},
		{2, 4, 4},
		{3, 4, 2},
	}

	start := 0
	numVertices := 5

	distances, previous := BellmanFord(edges, start, numVertices)

	fmt.Println("Distances:", distances)
	fmt.Println("Previous:", previous)
}
