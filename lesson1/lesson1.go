// Package fyne.First application
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson1.GUI.Package fyne.First application")
	w.Resize(fyne.NewSize(400, 400))
	label1 := widget.NewLabel("My first label")
	label2 := widget.NewLabel("My second label")
	w.SetContent(container.NewVBox(
		label1,
		label2,
	))
	w.ShowAndRun()
}
