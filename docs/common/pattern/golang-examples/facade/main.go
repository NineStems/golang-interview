package main

import "fmt"

type Subsystem1 struct{}

func (s *Subsystem1) Operation1() string {
	return "Subsystem1 operation"
}

type Subsystem2 struct{}

func (s *Subsystem2) Operation2() string {
	return "Subsystem2 operation"
}

type SubSys1Int interface {
	Operation1() string
}

type SubSys2Int interface {
	Operation2() string
}

type Facade struct {
	subsystem1 SubSys1Int
	subsystem2 SubSys2Int
}

func NewFacade() *Facade {
	return &Facade{
		subsystem1: &Subsystem1{},
		subsystem2: &Subsystem2{},
	}
}

func (f *Facade) Operation() string {
	result := "Facade initializes subsystems:\n"
	result += f.subsystem1.Operation1() + "\n"
	result += f.subsystem2.Operation2() + "\n"
	return result
}

// Пример использования
func main() {
	facade := NewFacade()
	result := facade.Operation()
	fmt.Println(result)
}
