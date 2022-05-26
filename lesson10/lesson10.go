// Set the app to dark theme
// fyne package -os linux -icon icon.png

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson10.GUI.Set the app to dark theme")
	w.Resize(fyne.NewSize(500, 500))
	btn := widget.NewButton("Change Theme", func() {
		a.Settings().SetTheme(theme.LightTheme())
	})
	a.Settings().SetTheme(theme.DarkTheme())
	w.SetContent(btn)
	w.ShowAndRun()
}
