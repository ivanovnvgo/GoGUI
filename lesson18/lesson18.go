// Drop down menu.Accordion
package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson18.GUI.Drop down menu.Accordion")
	w.Resize(fyne.NewSize(400, 400))
	// Create Accordion
	text := "The plugin adds a random text generator, capable of creating witty texts in different genres." +
		"Created text can be inserted newly at the caret, or replace a selection. " +
		"The dummy text generator is added to the main menu, tools menu and into the generate... " +
		"popup (Alt+Insert).  Used from the main menu, text is generated in the previously selected genre. " +
		"Used from the Tools menu or Generate popup, the plugin allows to choose from several genres: " +
		"Culinary Inspirations Esoteric Wisdom Fake Latin"
	textLabel1 := widget.NewLabel(text)
	textLabel1.Wrapping = fyne.TextWrapBreak // Create automatic text wrapping by lines to fit the application window

	item1 := widget.NewAccordionItem(
		"List1",
		textLabel1,
	)
	item2 := widget.NewAccordionItem(
		"Button",
		widget.NewButton("Say Hello",
			func() {
				fmt.Println("Hello!")
			},
		),
	)
	item3 := widget.NewAccordionItem(
		"Charters",
		widget.NewAccordion(
			widget.NewAccordionItem(
				"Charter1",
				widget.NewLabel("Text1"),
			),
			widget.NewAccordionItem(
				"Charter2",
				widget.NewLabel("Text2"),
			),
			widget.NewAccordionItem(
				"Charter3",
				widget.NewLabel("Text2"),
			),
			widget.NewAccordionItem(
				"Charter4",
				widget.NewLabel("Text2"),
			),
		),
	)
	accordion := widget.NewAccordion(item1, item2, item3)
	w.SetContent(accordion)
	w.ShowAndRun()
	a.Run()
}
