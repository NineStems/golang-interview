package main

import (
	"fmt"
)

type Cloneable interface {
	Clone() Cloneable
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Clone() Cloneable {
	return &Person{
		Name: p.Name + "_clone",
		Age:  p.Age,
	}
}

func main() {
	// создаем оригинальный объект
	original := &Person{
		Name: "John",
		Age:  30,
	}

	// клонируем объект
	clone := original.Clone()

	fmt.Println("Original:", original)
	fmt.Println("Clone:", clone)

}
