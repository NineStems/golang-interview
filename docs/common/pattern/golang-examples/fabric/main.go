package main

import "fmt"

// Creator определяет метод для создания объекта.
type Creator interface {
	Create() Product
}

// ConcreteCreator реализация.
type ConcreteCreator struct{}

// Create создает экземпляр одного из объектов.
func (cc *ConcreteCreator) Create(conn string) Product {
	switch conn {
	case "concrete":
		return &ConcreteProduct{}
	case "another":
		return &AnotherProduct{}
	}

	return nil
}

// Product определяет методы, которые должен реализовать создаваемый объект.
type Product interface {
	GetName() string
}

// ConcreteProduct первая версия реализация.
type ConcreteProduct struct{}

func (cp *ConcreteProduct) GetName() string {
	return "Concrete Product"
}

// AnotherProduct вторая версия реализация.
type AnotherProduct struct{}

func (ap *AnotherProduct) GetName() string {
	return "Another Product"
}

func main() {
	// Создаем новый экземпляр ConcreteCreator.
	cc := &ConcreteCreator{}

	// Создаем объект с помощью метода Create().
	cProduct := cc.Create("concrete")

	// Выводим имя созданного объекта.
	fmt.Println(cProduct.GetName())

	// Создаем другой объект под этим же интерфейсом с помощью метода Create().
	aProduct := cc.Create("another")

	// Выводим имя созданного объекта.
	fmt.Println(aProduct.GetName())
}
