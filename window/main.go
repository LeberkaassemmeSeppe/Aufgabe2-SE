package main

import "fmt"

func main() {
	innerWindow := NewWindow(
		NewFrame("Holzrahmen klein", 40.00),
		NewGlass("Klarglas klein", 25.50),
	)

	outerWindow := NewWindow(
		NewFrame("Holzrahmen gross", 80.00),
		NewGlass("Isolierglas", 60.00),
		innerWindow, // verschachteltes Window
	)

	fmt.Println(outerWindow.Display())
	fmt.Printf("Gesamtpreis: %.2f EUR\n", outerWindow.Price())
}
