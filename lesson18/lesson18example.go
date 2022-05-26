// Example use accordion and Card
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson18.GUI.Example use accordion and Card")
	w.Resize(fyne.NewSize(400, 400))
	card := widget.NewCard(
		"OOO XXX",
		"About us",
		widget.NewAccordion(
			widget.NewAccordionItem(
				"Paragraph1",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Subparagraph1.1",
						widget.NewLabel("Text from subparagraph1.1"),
					),
					widget.NewAccordionItem(
						"Subparagraph1.2",
						widget.NewLabel("Text from subparagraph1.2"),
					),
					widget.NewAccordionItem(
						"Subparagraph1.3",
						widget.NewLabel("Text from subparagraph1.3"),
					),
				),
			),
			widget.NewAccordionItem(
				"Paragraph2",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Subparagraph2.1",
						widget.NewLabel("Text from subparagraph2.1"),
					),
					widget.NewAccordionItem(
						"Subparagraph2.2",
						widget.NewLabel("Text from subparagraph2.2"),
					),
					widget.NewAccordionItem(
						"Subparagraph2.3",
						widget.NewLabel("Text from subparagraph2.3"),
					),
				),
			),
			widget.NewAccordionItem(
				"Paragraph3",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Subparagraph3.1",
						widget.NewLabel("Text from subparagraph3.1"),
					),
					widget.NewAccordionItem(
						"Subparagraph3.2",
						widget.NewLabel("Text from subparagraph3.2"),
					),
					widget.NewAccordionItem(
						"Subparagraph3.3",
						widget.NewLabel("Text from subparagraph3.3"),
					),
				),
			),
		),
	)
	w.SetContent(card)
	w.ShowAndRun()
}
