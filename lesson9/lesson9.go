//Positioning elements on the window using VBox and HBox
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson9.GUI.Positioning elements on the window using VBox and HBox")
	w.Resize(fyne.NewSize(500, 500))
	btn1 := widget.NewButton("Button1", func() {})
	btn2 := widget.NewButton("Button2", func() {})
	btn3 := widget.NewButton("Button3", func() {})
	btn4 := widget.NewButton("Button4", func() {})
	label1 := widget.NewLabel("Text1")
	label2 := widget.NewLabel("Text2")
	label3 := widget.NewLabel("Text3")
	label4 := widget.NewLabel("Text4")
	labelBox := container.NewHBox(
		label1,
		label2,
		label3,
		label4,
	)
	btnBox := container.NewHBox(
		btn1,
		btn2,
		btn3,
		btn4,
	)
	content := container.NewVBox(labelBox, btnBox)
	w.SetContent(content)
	w.ShowAndRun()
}
