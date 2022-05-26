// Arrangement on the grid of geometric shapes, setting their colors and borders
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Lesson8.GUI.Arrangement on the grid of geometric shapes, setting their colors and borders")
	w.Resize(fyne.NewSize(500, 500))
	circle1 := canvas.NewCircle(color.NRGBA{R: 70, G: 8, B: 163, A: 240})
	circle1.StrokeColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	circle1.StrokeWidth = 3
	rectangle1 := canvas.NewRectangle(color.NRGBA{R: 17, G: 94, B: 6, A: 255})
	rectangle1.StrokeColor = color.NRGBA{R: 11, G: 23, B: 9}
	rectangle1.StrokeWidth = 3
	line1 := canvas.NewLine(color.NRGBA{R: 195, G: 0, B: 0, A: 255})
	line1.StrokeWidth = 5
	icon1 := widget.NewIcon(theme.FileIcon())
	btn := widget.NewButton("Press me", func() {
		rectangle1.Hide()
		circle1.StrokeWidth = 10
		line1.StrokeWidth = 10
		line1.StrokeColor = color.NRGBA{R: 44, G: 24, B: 52, A: 255}
	})
	content1 := container.NewGridWithColumns(2, circle1, rectangle1)
	content2 := container.NewGridWithColumns(2, line1, icon1)
	contentAll := container.NewGridWithRows(3, content1, content2, btn)
	w.SetContent(
		contentAll,
	)
	w.ShowAndRun()
}
