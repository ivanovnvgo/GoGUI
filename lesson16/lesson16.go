// Creation of the main application menu, child menus and buttons
package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson16.GUI.Creation of the main application menu, child menus and buttons")
	w.Resize(fyne.NewSize(400, 400))
	// Structural diagram: main menu -> child menu -> buttons, start creating the application menu from the end
	// Create buttons:
	fileItem1 := fyne.NewMenuItem("New file", func() {
		os.Create("created.txt")
	})
	fileItem2 := fyne.NewMenuItem("Safe", func() {
		fmt.Println("File is saved!")
	})
	// Create child menu:
	menu1 := fyne.NewMenu("File", fileItem1, fileItem2)

	actionItem1 := fyne.NewMenuItem("Hello", func() {
		fmt.Println("Hello from menu")
	})
	actionItem2 := fyne.NewMenuItem("Bye", func() {
		fmt.Println("Bye")
	})
	actionItem3 := fyne.NewMenuItem("Good night", func() {
		fmt.Println("Good night")
	})
	menu2 := fyne.NewMenu("Actions", actionItem1, actionItem2, actionItem3)

	// Create main menu
	mainMenu := fyne.NewMainMenu(menu1, menu2)
	w.SetMainMenu(mainMenu)
	w.ShowAndRun()
}
