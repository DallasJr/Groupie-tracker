package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"groupie-tracker/filtre"
	"groupie-tracker/pages/mainPage"
	"groupie-tracker/structs"
)

func main() {
	structs.Load()
	filtre.CreateFiltersUI(filtre.NewFilters())
	fmt.Println("Launching app . . .")
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")
	mainPage.LoadPage(myWindow)
	myWindow.ShowAndRun()
}
