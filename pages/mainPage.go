package mainPage

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func LoadPage() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker")

	titleLabel := canvas.NewText("Groupie Tracker", color.White)
	titleLabel.TextSize = 50
	titleContainer := container.NewCenter(titleLabel)

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search here")

	searchButton := widget.NewButtonWithIcon("", theme.MailForwardIcon(), func() {

	})

	searchEntry.Resize(fyne.NewSize(415, searchEntry.MinSize().Height))
	searchEntry.Move(fyne.NewPos(-40, 0))
	searchButton.Resize(fyne.NewSize(searchButton.MinSize().Width, searchButton.MinSize().Height))
	searchButton.Move(fyne.NewPos(380, 0))

	inputContainer := container.NewWithoutLayout(searchEntry, searchButton)

	content := container.NewCenter(container.NewVBox(
		titleContainer,
		inputContainer,
	))
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 500))
	myWindow.ShowAndRun()
}
