package main

import "fmt"

type Graph struct {
	vertices int     // Количество вершин в графе
	edges    [][]int // Список смежных вершин для каждой вершины
}

// Функция для добавления ребра между двумя вершинами
func (g *Graph) addEdge(u int, v int) {
	g.edges[u] = append(g.edges[u], v)
}

// Рекурсивная функция для обхода графа в глубину
func dfs(g *Graph, visited []bool, v int) {
	// Помечаем текущую вершину как посещенную
	visited[v] = true
	fmt.Printf("%d ", v)

	// Рекурсивно обходим все смежные вершины
	for _, i := range g.edges[v] {
		if !visited[i] {
			dfs(g, visited, i)
		}
	}
}

func DFS(g *Graph, start int) {
	// Создаем слайс для хранения информации о посещении вершин
	visited := make([]bool, g.vertices)

	// Вызываем рекурсивную функцию для обхода графа в глубину
	dfs(g, visited, start)
}

func main() {
	// Создаем граф
	g := Graph{
		vertices: 5,
		edges:    make([][]int, 5),
	}

	// Добавляем ребра между вершинами
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	// Обходим граф в глубину, начиная с вершины 2
	fmt.Println("DFS traversal starting from vertex 2:")
	DFS(&g, 2)
}
