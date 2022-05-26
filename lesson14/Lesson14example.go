// Example lesson4 "Multiline input, work with multiline fields. MultilineEntry"
package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("Example working code")
	w.Resize(fyne.NewSize(250, 250))
	// MultilineEntry
	label1 := widget.NewLabel("Notes")
	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("Enter here ...")
	btn := widget.NewButton("Write", func() {
		file, err := os.Create("info.txt")
		if err != nil {
			fmt.Println("Errors read file", err)
			os.Exit(1)
		}
		defer file.Close()
		file.WriteString(entry1.Text)
	})
	cont := container.NewVBox(label1, entry1, btn)
	w.SetContent(cont)
	w.ShowAndRun()
}
