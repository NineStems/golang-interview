package main

import "fmt"

type Component interface {
	Operation() string
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

type Decorator struct {
	component Component
}

func (d *Decorator) Operation() string {
	if d.component != nil {
		return d.component.Operation()
	}
	return ""
}

type ConcreteDecorator1 struct {
	Decorator
}

func (c *ConcreteDecorator1) Operation() string {
	return "ConcreteDecorator1(" + c.Decorator.Operation() + ")"
}

type ConcreteDecorator2 struct {
	Decorator
}

func (c *ConcreteDecorator2) Operation() string {
	return "ConcreteDecorator2(" + c.Decorator.Operation() + ")"
}

func main() {
	component := &ConcreteComponent{}

	decorator1 := &ConcreteDecorator1{}
	decorator1.component = component

	decorator2 := &ConcreteDecorator2{}
	decorator2.component = decorator1

	result := decorator2.Operation()
	fmt.Println(result)
}
