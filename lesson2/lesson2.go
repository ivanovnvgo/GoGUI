// Buttons, label and entry
package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson2.GUI.Buttons, label and entry")
	w.Resize(fyne.NewSize(400, 600))
	label := widget.NewLabel("Text label")
	entry := widget.NewEntry()
	btn := widget.NewButton("Button", func() {
		data := entry.Text
		fmt.Println(data)
		label.SetText(data)
	})
	w.SetContent(container.NewVBox(
		label,
		entry,
		btn,
	))
	w.ShowAndRun()
}
