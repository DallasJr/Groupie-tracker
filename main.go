package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"groupie-tracker/core"
	"groupie-tracker/pages"
)

func main() {

	fmt.Println("Launching app . . .")

	//load les Artists depuis l'API
	core.Load()
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")
	pages.LoadMainPage(myWindow)
	myWindow.ShowAndRun()
}
