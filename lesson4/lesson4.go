// My First Application.CheckGroup.RadioGroup
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"net/url"
)

func main() {
	// Create an application object
	a := app.New()
	// Create a new window
	w := a.NewWindow("lesson4.GUI.My First Application")
	w.Resize(fyne.NewSize(300, 500))
	title := widget.NewLabel("Checkout")
	nameLabel := widget.NewLabel("Your name: ")
	name := widget.NewEntry()
	foodLabel := widget.NewLabel("Choose food to order")
	food := widget.NewCheckGroup([]string{"Pizza", "Cake", "Nuggets", "Burger", "Cola"}, func(s []string) {})
	genderLabel := widget.NewLabel("What's your gender: ")
	gender := widget.NewRadioGroup([]string{"male", "female"}, func(s string) {})
	result := widget.NewLabel("")
	btn := widget.NewButton("Checkout", func() {
		userName := name.Text
		foodArray := food.Selected
		genderValue := gender.Selected
		// First, write the single values
		result.SetText(userName + "\n" + genderValue + "\n")
		// Then write the values from the slice
		for _, i := range foodArray {
			result.SetText(result.Text + i + "\n")
		}
	})
	// Create a hyperLink
	url, err := url.Parse("https://pkg.go.dev/")
	if err != nil {
		log.Fatal("hyperlink unreadable")
	}
	link := widget.NewHyperlink("Go Golang!", url)
	// Output per page
	w.SetContent(container.NewVBox(
		title,
		nameLabel,
		name,
		foodLabel,
		food,
		genderLabel,
		gender,
		btn,
		result,
		link,
	))

	// Launching the application
	w.ShowAndRun()
}
