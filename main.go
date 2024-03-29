package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"groupie-tracker/core"
	"groupie-tracker/pages"
)

const resX = 1280
const resY = 720

func main() {

	fmt.Println("Launching app . . .")

	//load les Artists depuis l'API
	core.Load()
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")
	pages.LoadMainPage(myWindow)
	myWindow.Resize(fyne.NewSize(resX, resY))
	myWindow.ShowAndRun()
}
