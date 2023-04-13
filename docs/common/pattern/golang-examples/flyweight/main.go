package main

import "fmt"

type Flyweight interface {
	Operation() string
}

type concreteFlyweight struct {
	sharedState string
}

func (f *concreteFlyweight) Operation() string {
	return f.sharedState
}

type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}
}

func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := f.flyweights[key]; ok {
		return flyweight
	}
	flyweight := &concreteFlyweight{sharedState: key}
	f.flyweights[key] = flyweight
	return flyweight
}

// Пример использования
func main() {
	factory := NewFlyweightFactory()
	flyweight1 := factory.GetFlyweight("key1")
	flyweight2 := factory.GetFlyweight("key2")
	flyweight3 := factory.GetFlyweight("key1")

	fmt.Println(flyweight1.Operation()) // Output: key1
	fmt.Println(flyweight2.Operation()) // Output: key2
	fmt.Println(flyweight3.Operation()) // Output: key1

	fmt.Println(flyweight1 == flyweight3) // Output: true
}
