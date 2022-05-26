// Changes color.Canvas, color, rgba
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson7.GUI.Changes color.Canvas, color, rgba")
	w.Resize(fyne.NewSize(300, 300))
	colorTextRGBA := color.NRGBA{R: 57, G: 129, B: 255, A: 100} // https://rgbacolorpicker.com
	text := canvas.NewText("Text string of color RGBA", colorTextRGBA)
	w.SetContent(text)
	w.ShowAndRun()
}
