package main

import "fmt"

// Glass ist ein Leaf.
type Glass struct {
	name  string
	price float64
}

func NewGlass(name string, price float64) *Glass {
	return &Glass{name: name, price: price}
}

func (g *Glass) Price() float64 { return g.price }

func (g *Glass) Display() string {
	return fmt.Sprintf("Glass %s (%.2f EUR)", g.name, g.price)
}
