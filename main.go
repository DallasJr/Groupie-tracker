package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"groupie-tracker/pages/mainPage"
)

func main() {
	fmt.Println("Launching app . . .")
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")
	mainPage.LoadPage(myWindow)
	myWindow.ShowAndRun()
}
