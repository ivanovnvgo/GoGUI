// Multiline input, work with multiline fields. MultilineEntry
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson14.GUI.Multiline input, work with multiline fields. MultilineEntry")
	w.Resize(fyne.NewSize(250, 250))
	// MultilineEntry
	text := "gggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"
	entry := widget.NewMultiLineEntry()
	// entry.SetPlaceHolder("Enter here...")
	// entry.SetText("Lettering template for input")
	entry.SetText(text)
	//entry.Wrapping = fyne.TextWrapBreak
	entry.Wrapping = fyne.TextWrapOff
	w.SetContent(entry)
	w.ShowAndRun()
}
