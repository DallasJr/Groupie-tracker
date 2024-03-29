package main

import (
	"fmt"
	"groupie-tracker/pages/mainPage"

	"fyne.io/fyne/v2/app"
)

func main() {

	fmt.Println("Launching app . . .")

	//load les Artists depuis l'API
	//structs.Load()
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")
	mainPage.LoadPage(myWindow)
	myWindow.ShowAndRun()
}
