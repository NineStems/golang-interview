package main

import "fmt"

// Singleton реализация структуры-одиночки.
type Singleton struct{}

// singleton экземпляр одиночки.
var singleton *Singleton

// newSingleton конструктор.
func newSingleton() *Singleton {
	if singleton != nil {
		return singleton
	}
	singleton = &Singleton{}
	return singleton
}

func main() {
	single := newSingleton()
	fmt.Println(single)
	nextSingle := newSingleton()
	fmt.Println(nextSingle, single, single == nextSingle)
}
