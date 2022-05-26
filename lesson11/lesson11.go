// Add images by links
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"log"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson11.GUI.Add images by links")
	w.Resize(fyne.NewSize(800, 500))
	a.Settings().SetTheme(theme.LightTheme())
	res, err := fyne.LoadResourceFromURLString("https://bitok.blog/uploads/images/Golang.png")
	if err != nil {
		log.Fatal(err)
	}
	img := canvas.NewImageFromResource(res)
	w.SetContent(img)
	w.ShowAndRun()
}
