package main

import "fmt"

// AbstractFactory определяет методы для создания объектов.
type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

// ConcreteFactory1 реализация AbstractFactory.
type ConcreteFactory1 struct{}

func (cf1 *ConcreteFactory1) CreateProductA() ProductA {
	return &ConcreteProductA1{}
}

func (cf1 *ConcreteFactory1) CreateProductB() ProductB {
	return &ConcreteProductB1{}
}

// ConcreteFactory2 реализация AbstractFactory.
type ConcreteFactory2 struct{}

func (cf2 *ConcreteFactory2) CreateProductA() ProductA {
	return &ConcreteProductA2{}
}

func (cf2 *ConcreteFactory2) CreateProductB() ProductB {
	return &ConcreteProductB2{}
}

// ProductA определяет методы, которые должен реализовать создаваемый объект типа A.
type ProductA interface {
	GetName() string
}

// ConcreteProductA1 Реализация ProductA.
type ConcreteProductA1 struct{}

func (cpa1 *ConcreteProductA1) GetName() string {
	return "Concrete Product A1"
}

// ConcreteProductA2 Реализация ProductA.
type ConcreteProductA2 struct{}

func (cpa2 *ConcreteProductA2) GetName() string {
	return "Concrete Product A2"
}

// ProductB определяет методы, которые должен реализовать создаваемый объект типа B.
type ProductB interface {
	GetName() string
}

// ConcreteProductB1 Реализация ProductB.
type ConcreteProductB1 struct{}

func (cpb1 *ConcreteProductB1) GetName() string {
	return "Concrete Product B1"
}

// ConcreteProductB2 Реализация ProductB.
type ConcreteProductB2 struct{}

func (cpb2 *ConcreteProductB2) GetName() string {
	return "Concrete Product B2"
}

func main() {
	// Создаем новый экземпляр ConcreteFactory1.
	cf1 := &ConcreteFactory1{}

	// Создаем объекты типа A и B с помощью методов CreateProductA() и CreateProductB().
	productA1 := cf1.CreateProductA()
	productB1 := cf1.CreateProductB()

	// Выводим имена созданных объектов.
	fmt.Println(productA1.GetName())
	fmt.Println(productB1.GetName())

	// Создаем новый экземпляр ConcreteFactory2.
	cf2 := &ConcreteFactory2{}

	// Создаем объекты типа A и B с помощью методов CreateProductA() и CreateProductB().
	productA2 := cf2.CreateProductA()
	productB2 := cf2.CreateProductB()

	// Выводим имена созданных объектов.
	fmt.Println(productA2.GetName())
	fmt.Println(productB2.GetName())
}
