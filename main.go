package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
)

func main() {
	fmt.Println("Launching app . . .")
	a := app.New()
	w := a.NewWindow("Groupie Tracker")
	w.ShowAndRun()
}
