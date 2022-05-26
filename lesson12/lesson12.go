// Set icon in app
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson12.GUI.Set icon in app")
	w.Resize(fyne.NewSize(400, 400))
	// SetIcon(ресурс)
	ic, err := fyne.LoadResourceFromPath("icon.png")
	if err != nil {
		log.Fatal(err)
	}
	w.SetIcon(ic)
	w.SetContent(widget.NewLabel("Hi!!!"))
	w.ShowAndRun()
}
