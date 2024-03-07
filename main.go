package main

import (
	"fmt"
	"groupie-tracker/pages/mainPage"
	"groupie-tracker/search"

	"fyne.io/fyne/v2/app"
)

func main() {
	search.Load()
	fmt.Println("Launching app . . .")
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")
	mainPage.LoadPage(myWindow)
	myWindow.ShowAndRun()
}
