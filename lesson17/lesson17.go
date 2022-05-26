// Submenus nested in child menus
package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson17.GUI.Submenus nested in child menus")
	w.Resize(fyne.NewSize(400, 400))
	// Create child menus that will have nested menus
	item1 := fyne.NewMenuItem("Actions", nil)
	item2 := fyne.NewMenuItem("Say", nil)
	menu1 := fyne.NewMenu("Menu1", item1, item2)
	item1.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Print", func() {
			fmt.Println("Print")
		}),
		fyne.NewMenuItem("Open", func() {
			fmt.Println("Open")
		}),
		fyne.NewMenuItem("Save", func() {
			fmt.Println("Save")
		}),
		fyne.NewMenuItem("Save all", func() {
			fmt.Println("Save all")
		}),
	)
	item2.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Hi", func() {
			fmt.Println("Hi")
		}),
		fyne.NewMenuItem("Hello", func() {
			fmt.Println("Hello")
		}),
		fyne.NewMenuItem("Bye", func() {
			fmt.Println("Bye")
		}),
	)
	mainMenu := fyne.NewMainMenu(menu1)
	w.SetMainMenu(mainMenu)
	w.ShowAndRun()
}
