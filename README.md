# Window – Composite Pattern in Go

Objektorientierte Umsetzung eines UML-Klassendiagramms mit dem Composite Pattern.
Ein `Window` besteht aus beliebig vielen `Item`s (mindestens einem) und kann
andere `Window`s enthalten.

Das Projekt wurde unter Verwendung der kostenfreien Version von Claude (genauer Sonnet 5.5 medium) erstellt.
Der zur Erstellung verwendete Prompt befindet sich in der Datei `prompt.md`.

## Struktur

```
./                   # Projektordner
├── README.md
├── prompt.md
└── window/
    ├── go.mod
    ├── item.go      # Interface Item
    ├── glass.go     # Leaf: Glass
    ├── frame.go     # Leaf: Frame
    ├── window.go    # Composite: Window
    └── main.go      # Demo
```

## Klassen

| Typ      | Rolle     | Beschreibung                                         |
|----------|-----------|------------------------------------------------------|
| `Item`   | Interface | Definiert `Price() float64` und `Display() string`   |
| `Glass`  | Leaf      | Hat `name` und `price`                               |
| `Frame`  | Leaf      | Hat `name` und `price`                               |
| `Window` | Composite | Enthält 1..* `Item`s, auch verschachtelte `Window`s  |

Go kennt keine Klassenvererbung. `Item` ist deshalb ein Interface, das
`Glass`, `Frame` und `Window` implizit durch ihre Methoden erfüllen.

## Multiplizität 1..*

`NewWindow(first Item, rest ...Item)` verlangt mindestens ein Item als
Pflichtparameter. Ein leeres `Window` lässt sich dadurch schon beim
Kompilieren nicht erzeugen.

## Verhalten

- `Window.Price()` summiert rekursiv die Preise aller enthaltenen Items.
- `Window.Display()` gibt das Fenster und seine Bestandteile eingerückt als
  mehrzeiligen String zurück. Jede Zeile enthält den Preis des Bestandteils (Glass, Frame) bzw.
  die Summe der Bestandteile (Window).

## Starten

Im Projektordner:

```
go -C window run .
```

Die Option `-C` wechselt vor dem Ausführen in den angegebenen Ordner
(ab Go 1.20).

Alternativ direkt im Ordner `window/`:

```
go run .
```

Erwartete Ausgabe der Demo:

```
Window (205.50 EUR)
  Frame Holzrahmen gross (80.00 EUR)
  Glass Isolierglas (60.00 EUR)
  Window (65.50 EUR)
    Frame Holzrahmen klein (40.00 EUR)
    Glass Klarglas klein (25.50 EUR)
Gesamtpreis: 205.50 EUR
```

## Voraussetzungen

- Go 1.21 oder neuer
- Keine externen Bibliotheken
