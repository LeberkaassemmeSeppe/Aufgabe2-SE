Implementiere objektorientiert in Go das folgende UML-Klassendiagramm unter Verwendung des Composite Pattern:

**Klassen**
- `Item` definiert `Price() float64` und `Display() string`.
- `Window`, `Glass` und `Frame` sind Unterklassen von `Item`.
- `Glass` und `Frame` sind Leafs mit den Attributen `name` (string) und `price` (float64).
- `Window` ist das Composite und enthält mindestens ein `Item` (1..*). Diese Multiplizität muss in `Window` erzwungen werden. Windows dürfen verschachtelt werden.

**Verhalten**
- `Window.Price()` berechnet rekursiv den Gesamtpreis der enthaltenen Items.
- `Window.Display()` gibt das Fenster und seine Bestandteile eingerückt als
  mehrzeiligen String zurück. Jede Zeile enthält den Preis des Bestandteils (Glass, Frame) bzw.
  die Summe der Bestandteile (Window).

**Umsetzung in Go**
- Bilde `Item` als Interface ab (Go kennt keine Klassenvererbung).
- Halte den Code einfach und nachvollziehbar.
- Ergänze eine kleine `main()`-Funktion mit einem verschachtelten Window als Demo.
- Verzichte auf Zusatzfunktionen und externe Bibliotheken.

Erstelle zusätzlich eine Readme, die das Programm beschreibt.
