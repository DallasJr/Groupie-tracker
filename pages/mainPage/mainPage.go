package mainPage

import (
	"groupie-tracker/.idea/filtre" // Importez le package filtre ici
	"groupie-tracker/pages/artistPage"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func LoadPage(myWindow fyne.Window) {

	titleLabel := canvas.NewText("Groupie Tracker", color.White)
	titleLabel.TextSize = 50
	titleContainer := container.NewCenter(titleLabel)

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search here")

	searchButton := widget.NewButtonWithIcon("", theme.MailForwardIcon(), func() {
		artistPage.LoadPage(myWindow)
	})

	// Créez les filtres
	filters := filtre.NewFilters()
	filtersUI := filtre.CreateFiltersUI(filters)
	myWindow.SetContent(filtersUI) // Définissez le contenu de la fenêtre avec les filtres

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
}
