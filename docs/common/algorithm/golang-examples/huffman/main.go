package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	char  rune
	freq  int
	left  *Node
	right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func buildHuffmanTree(freq map[rune]int) *Node {
	pq := make(PriorityQueue, len(freq))
	i := 0
	for c, f := range freq {
		pq[i] = &Node{c, f, nil, nil}
		i++
	}
	heap.Init(&pq)
	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)
		parent := &Node{0, left.freq + right.freq, left, right}
		heap.Push(&pq, parent)
	}
	return heap.Pop(&pq).(*Node)
}

func encode(root *Node, code map[rune]string, prefix string) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		code[root.char] = prefix
		return
	}
	encode(root.left, code, prefix+"0")
	encode(root.right, code, prefix+"1")
}

func compress(text string) (string, map[rune]string) {
	freq := make(map[rune]int)
	for _, c := range text {
		freq[c]++
	}
	root := buildHuffmanTree(freq)
	code := make(map[rune]string)
	encode(root, code, "")
	compressed := ""
	for _, c := range text {
		compressed += code[c]
	}
	return compressed, code
}

func decompress(compressed string, code map[rune]string) string {
	revCode := make(map[string]rune)
	for c, p := range code {
		revCode[p] = c
	}
	current := ""
	decompressed := ""
	for _, b := range compressed {
		current += string(b)
		if c, ok := revCode[current]; ok {
			decompressed += string(c)
			current = ""
		}
	}
	return decompressed
}

func main() {
	text := "hello world"
	compressed, code := compress(text)
	fmt.Println("Original text:", text)
	fmt.Println("Compressed text:", compressed)
	fmt.Println("Code table:", code)
	decompressed := decompress(compressed, code)
	fmt.Println("Decompressed text:", decompressed)
}
