package main

import (
	"fmt"
	"strings"
)

type Component interface {
	Add(c Component)
	Remove(c Component)
	Display(depth int)
}

type Leaf struct {
	name string
}

func (l *Leaf) Add(c Component) {
	fmt.Println("Cannot add to a leaf")
}

func (l *Leaf) Remove(c Component) {
	fmt.Println("Cannot remove from a leaf")
}

func (l *Leaf) Display(depth int) {
	fmt.Printf("%s%s\n", strings.Repeat("-", depth), l.name)
}

type Composite struct {
	name     string
	children []Component
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Remove(component Component) {
	for i, child := range c.children {
		if child == component {
			c.children = append(c.children[:i], c.children[i+1:]...)
			return
		}
	}
}

func (c *Composite) Display(depth int) {
	fmt.Printf("%s%s\n", strings.Repeat("-", depth), c.name)
	for _, child := range c.children {
		child.Display(depth + 2)
	}
}

func main() {
	root := &Composite{name: "root"}
	root.Add(&Leaf{name: "leaf1"})
	root.Add(&Leaf{name: "leaf2"})

	comp := &Composite{name: "composite1"}
	comp.Add(&Leaf{name: "leaf3"})
	comp.Add(&Leaf{name: "leaf4"})

	root.Add(comp)

	root.Display(1)
}
