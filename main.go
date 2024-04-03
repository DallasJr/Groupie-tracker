package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"groupie-tracker/core"
	"groupie-tracker/pages"
	"groupie-tracker/structs"
)

const resX = 1280
const resY = 720

func main() {

	fmt.Println("Launching app . . .")

	myApp := app.New()
	structs.InitializeSpotify()

	// Load les Artists depuis l'API
	core.Load()

	myWindow := myApp.NewWindow("Groupie Tracker")

	// Récuperer le contenu de la page principale
	pages.LoadMainPage(myApp, myWindow)

	myWindow.Resize(fyne.NewSize(resX, resY))
	myWindow.ShowAndRun()
}
