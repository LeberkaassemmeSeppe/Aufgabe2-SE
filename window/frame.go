package main

import "fmt"

// Frame ist ein Leaf.
type Frame struct {
	name  string
	price float64
}

func NewFrame(name string, price float64) *Frame {
	return &Frame{name: name, price: price}
}

func (f *Frame) Price() float64 { return f.price }

func (f *Frame) Display() string {
	return fmt.Sprintf("Frame %s (%.2f EUR)", f.name, f.price)
}
