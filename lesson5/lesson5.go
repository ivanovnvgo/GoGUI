// Hyperlink
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	url2 "net/url"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson5.GUI.Hyperlink")
	w.Resize(fyne.NewSize(400, 600))
	url, err := url2.Parse("https://bolid.ru/")
	if err != nil {
		log.Println("Bad URL")
	}
	link := widget.NewHyperlink("Bolid", url)
	w.SetContent(container.NewVBox(
		link,
	))
	w.ShowAndRun()
}
