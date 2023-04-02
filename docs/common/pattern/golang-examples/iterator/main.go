package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Aggregate interface {
	GetIterator() Iterator
}

type NameRepository struct {
	Names []string
}

func (n *NameRepository) GetIterator() Iterator {
	return &NameIterator{Names: n.Names, Index: 0}
}

type NameIterator struct {
	Names []string
	Index int
}

func (n *NameIterator) HasNext() bool {
	if n.Index < len(n.Names) {
		return true
	}
	return false
}

func (n *NameIterator) Next() interface{} {
	if n.HasNext() {
		value := n.Names[n.Index]
		n.Index++
		return value
	}
	return nil
}

func main() {
	names := []string{"John", "Doe", "Jane", "Smith"}
	nameRepo := &NameRepository{Names: names}
	iterator := nameRepo.GetIterator()

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
