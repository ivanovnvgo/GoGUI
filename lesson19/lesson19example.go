// Example code
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Lesson19.GUI.Example code")
	w.Resize(fyne.NewSize(400, 400))

	title := widget.NewLabel("Spelling")
	rightLabel := widget.NewLabel("right")
	wrongLabel := widget.NewLabel("wrong")
	rText := widget.NewLabel("The plugin adds a random text generator, capable of creating witty texts in different genres." +
		"Created text can be inserted newly at the caret, or replace a selection. " +
		"The dummy text generator is added to the main menu, tools menu and into the generate... " +
		"popup (Alt+Insert).  Used from the main menu, text is generated in the previously selected genre. " +
		"Used from the Tools menu or Generate popup, the plugin allows to choose from several genres: " +
		"Culinary Inspirations Esoteric Wisdom Fake Latin")
	rText.Wrapping = fyne.TextWrapBreak
	wText := widget.NewLabel("The plugin adds a random text generator, capable of creating witty texts in different genres." +
		"Created text can be inserted newly at the caret, or replace a selection. " +
		"The dummy text generator is added to the main menu, tools menu and into the generate... " +
		"popup (Alt+Insert).  Used from the main menu, text is generated in the previously selected genre. " +
		"Used from the Tools menu or Generate popup, the plugin allows to choose from several genres: " +
		"Culinary Inspirations Esoteric Wisdom Fake Latin")
	wText.Wrapping = fyne.TextWrapBreak
	rightBox := container.NewVBox(rightLabel, rText)
	wrongBox := container.NewVBox(wrongLabel, wText)
	split := container.NewHSplit(rightBox, wrongBox)
	cont := container.NewVBox(title, split)
	w.SetContent(cont)
	w.ShowAndRun()
}
