// Creating and working with a component Card, practical use
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson13.GUI.Creating and working with a component Card, practical use")
	w.Resize(fyne.NewSize(500, 500))
	// Component Card
	card1 := widget.NewCard(
		"Name component Card 1 (text)",
		"Description component Card 1 (text)",
		canvas.NewImageFromFile("1.jpg"), // You can insert a button, an inscription from the widget library
	)
	card2 := widget.NewCard(
		"Name component Card 2 (text)",
		"Description component Card 2 (text)",
		widget.NewButton("Card 2", func() {}),
	)
	box := container.NewHBox(card1, card2)
	w.SetContent(box)
	w.ShowAndRun()
}
