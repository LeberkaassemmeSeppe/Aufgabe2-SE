package main

// Item ist die gemeinsame Schnittstelle für Leafs und Composite.
type Item interface {
	Price() float64
	Display() string
}
