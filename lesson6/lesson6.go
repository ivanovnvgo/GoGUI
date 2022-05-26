// Adding Images to the Application
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
	w := a.NewWindow("lesson6.GUI.Adding Images to the Application")
	w.Resize(fyne.NewSize(1200, 300))
	img1 := canvas.NewImageFromFile("1.jpg")
	img2 := canvas.NewImageFromFile("2.jpg")
	w.SetContent(container.NewGridWithColumns(3, img1, img2, widget.NewButton("Press", func() {})))

	w.ShowAndRun()
}
