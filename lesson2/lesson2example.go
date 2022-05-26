// Example code.Buttons, label and entry.Calculator
package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson2.GUI.Example code.Calculator")

	label1 := widget.NewLabel("Enter first number")
	entry1 := widget.NewEntry()

	label2 := widget.NewLabel("Enter second number")
	entry2 := widget.NewEntry()

	answer := widget.NewLabel("") // Write the answer here

	btn := widget.NewButton("count", func() {
		n1, err1 := strconv.ParseFloat(entry1.Text, 64)
		n2, err2 := strconv.ParseFloat(entry2.Text, 64)
		if err1 != nil || err2 != nil {
			answer.SetText("Input Error")
		} else {
			sum := n1 + n2
			sub := n1 - n2
			mul := n1 * n2
			div := n1 / n2
			answer.SetText(fmt.Sprintf("(+) %f\n(-) %f\n(*) %f\n(/) %f\n", sum, sub, mul, div))
		}
	})

	// Output per page
	w.SetContent(container.NewVBox(
		label1,
		entry1,
		label2,
		entry2,
		answer,
		btn,
	))
	w.ShowAndRun()
}
