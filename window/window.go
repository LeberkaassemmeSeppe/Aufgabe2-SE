package main

import (
	"fmt"
	"strings"
)

// Window ist das Composite und enthält mindestens ein Item (1..*).
type Window struct {
	items []Item
}

// NewWindow erzwingt die Multiplizität 1..*:
// Der erste Parameter ist Pflicht, weitere Items sind optional.
// Ein Window ohne Item lässt sich daher gar nicht erst erzeugen.
func NewWindow(first Item, rest ...Item) *Window {
	items := append([]Item{first}, rest...)
	return &Window{items: items}
}

// Price summiert rekursiv die Preise aller enthaltenen Items.
func (w *Window) Price() float64 {
	total := 0.0
	for _, item := range w.items {
		total += item.Price()
	}
	return total
}

// Display gibt das Fenster und seine Bestandteile eingerückt aus.
func (w *Window) Display() string {
	lines := []string{fmt.Sprintf("Window (%.2f EUR)", w.Price())}
	for _, item := range w.items {
		// Jede Zeile des Kindes (auch bei verschachtelten Windows) wird eingerückt.
		for _, line := range strings.Split(item.Display(), "\n") {
			lines = append(lines, "  "+line)
		}
	}
	return strings.Join(lines, "\n")
}
