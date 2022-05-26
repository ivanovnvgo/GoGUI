// Resizing and positioning on a window
package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("lesson15.GUI.Resizing and positioning on a window")
	// Set the application window size
	w.Resize(fyne.NewSize(400, 400))

	/*
		// Create new button
		btn := widget.NewButton("Click me", func() { fmt.Println("Hi, Nick") })
		// Set button size
		btn.Resize(fyne.NewSize(200, 70))
		// Set the position of the button on the application window
		btn.Move(fyne.NewPos(100, 3))
		// Create a new data entry field
		entry := widget.NewEntry()
		// Set entry size
		entry.Resize(fyne.NewSize(300, 40))
		// Set the position of the entry on the application window
		entry.Move(fyne.NewPos(50, 93))
		// Create label
		label := widget.NewLabel("TEXT")
		// label haven't Resize
		// Set the position of the label on the window
		label.Move(fyne.NewPos(170, 153))
		// To change the size of the button, you need to create a container
		cont := container.NewWithoutLayout(entry, btn, label)
	*/

	label := widget.NewLabel("Регистрация")
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Move(fyne.NewPos(130, 3))

	entryNameNewUser := widget.NewEntry()
	entryNameNewUser.Resize(fyne.NewSize(300, 40))
	entryNameNewUser.Move(fyne.NewPos(50, 46))
	entryNameNewUser.SetPlaceHolder("Имя пользователя")

	entryPasswordNewUser := widget.NewPasswordEntry() // Hidden password entry
	entryPasswordNewUser.Resize(fyne.NewSize(300, 40))
	entryPasswordNewUser.Move(fyne.NewPos(50, 86))
	entryPasswordNewUser.SetPlaceHolder("Пароль")

	entryMailNewUser := widget.NewEntry()
	entryMailNewUser.Resize(fyne.NewSize(300, 40))
	entryMailNewUser.Move(fyne.NewPos(50, 126))
	entryMailNewUser.SetPlaceHolder("Mail")

	entryInfoNewUser := widget.NewMultiLineEntry()
	entryInfoNewUser.Resize(fyne.NewSize(300, 100))
	entryInfoNewUser.Move(fyne.NewPos(50, 166))
	entryInfoNewUser.SetPlaceHolder("Дополнительная информация")

	btnSubmit := widget.NewButton("Submit", func() {
		file, err := os.Create("user.txt")
		if err != nil {
			fmt.Println("Error create file", err)
			os.Exit(1)
		}
		defer file.Close()

		userName := entryNameNewUser.Text
		userPassword := entryPasswordNewUser.Text
		userMail := entryMailNewUser.Text
		userInfo := entryInfoNewUser.Text
		//fmt.Printf("New User Registration: %s\nPassword: %s\nMail: %v\nInfo: \n%s\n", userName, userPassword, userMail, userInfo)
		data := fmt.Sprintf(
			"user name: %s\n"+
				"user password: %s\n"+
				"user mail: %v\n"+
				"user info: %s\n",
			userName, userPassword, userMail, userInfo)
		file.WriteString(data)
	})
	btnSubmit.Resize(fyne.NewSize(200, 40))
	btnSubmit.Move(fyne.NewPos(100, 270))

	cont := container.NewWithoutLayout(
		label,
		entryNameNewUser,
		entryPasswordNewUser,
		entryMailNewUser,
		entryInfoNewUser,
		btnSubmit,
	)
	w.SetContent(cont)
	w.ShowAndRun()
}
