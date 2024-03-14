package main

import (
	"fmt"
	"groupie-tracker/.idea/filtre"
	"groupie-tracker/pages/mainPage"
	"groupie-tracker/search"

	"fyne.io/fyne/v2/app"
)

func main() {
	search.Load()
	filtre.CreateFiltersUI(filtre.NewFilters())
	fmt.Println("Launching app . . .")
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")
	mainPage.LoadPage(myWindow)
	myWindow.ShowAndRun()
}
