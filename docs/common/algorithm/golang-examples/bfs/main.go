package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) BFS() []int {
	result := []int{}
	queue := []*Node{n}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return result
}

func main() {
	root := &Node{
		value: 1,
		left:  &Node{value: 2, left: &Node{value: 4}, right: &Node{value: 5}},
		right: &Node{value: 3, left: &Node{value: 6}, right: &Node{value: 7}},
	}
	fmt.Println(root.BFS()) // Output: [1 2 3 4 5 6 7]
}
