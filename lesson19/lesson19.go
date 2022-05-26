// HSplit, VSplit
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson19.GUI.HSplit, VSplit")
	w.Resize(fyne.NewSize(400, 400))
	label1 := widget.NewLabel("Text of label1")
	btn1 := widget.NewButton("Button1", func() {})
	label2 := widget.NewLabel("Text of label2")
	btn2 := widget.NewButton("Button2", func() {})
	contVbox1 := container.NewVBox(label1, btn1)
	contVbox2 := container.NewVBox(label2, btn2)
	//HSplit
	//containerHSplit := container.NewHSplit(contVbox1, contVbox2)
	//VSplit
	containerVSplit := container.NewVSplit(contVbox1, contVbox2)
	contVVbox := container.NewVBox(containerVSplit)
	w.SetContent(contVVbox)
	w.ShowAndRun()
}
