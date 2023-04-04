package main

import (
	"fmt"
	"math"
)

type Graph struct {
	nodes []int
	edges map[int]map[int]int
}

func NewGraph() *Graph {
	return &Graph{
		nodes: []int{},
		edges: make(map[int]map[int]int),
	}
}

func (g *Graph) AddNode(node int) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) AddEdge(startNode, endNode, weight int) {
	if _, ok := g.edges[startNode]; !ok {
		g.edges[startNode] = make(map[int]int)
	}
	g.edges[startNode][endNode] = weight
}

func (g *Graph) Dijkstra(start int) map[int]int {
	distances := make(map[int]int)
	for _, node := range g.nodes {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0

	visited := make(map[int]bool)
	for len(visited) < len(g.nodes) {
		node := g.minDistance(distances, visited)
		visited[node] = true

		for neighbor, weight := range g.edges[node] {
			newDistance := distances[node] + weight
			if newDistance < distances[neighbor] {
				distances[neighbor] = newDistance
			}
		}
	}

	return distances
}

func (g *Graph) minDistance(distances map[int]int, visited map[int]bool) int {
	min := math.MaxInt32
	var minNode int
	for node, distance := range distances {
		if distance < min && !visited[node] {
			min = distance
			minNode = node
		}
	}
	return minNode
}

func main() {
	graph := NewGraph()

	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)
	graph.AddNode(4)
	graph.AddNode(5)

	graph.AddEdge(1, 2, 10)
	graph.AddEdge(1, 4, 30)
	graph.AddEdge(1, 5, 100)
	graph.AddEdge(2, 3, 50)
	graph.AddEdge(3, 5, 10)
	graph.AddEdge(4, 3, 20)
	graph.AddEdge(4, 5, 60)

	distances := graph.Dijkstra(1)

	fmt.Println(distances)
}
