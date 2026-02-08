package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 1. Create the app instance
	myApp := app.New()

	// 2. Create a window
	myWindow := myApp.NewWindow("Fyne Simple App")

	// 3. Define your widgets (labels, buttons, etc.)
	myLabel := widget.NewLabel("Hello, Fyne!")
	myButton := widget.NewButton("Click Me", func() {
		myLabel.SetText("You clicked the button!")
	})

	// 4. Set the content and layout
	// container.NewVBox stacks items vertically
	myWindow.SetContent(container.NewVBox(
		myLabel,
		myButton,
	))

	// 5. Show and run (this blocks until the window is closed)
	myWindow.ShowAndRun()
}
