package artistPage

import (
	"groupie-tracker/.idea/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func LoadPage(myWindow fyne.Window) {
	m := structs.NewMap("north_carolina-usa", 4)
	m.AddMarker(m.GetLat(), m.GetLong(), "394e70", "wave-square")
	img := m.GetImg()
	fyneImg := canvas.NewImageFromImage(img)
	content := container.NewMax(fyneImg)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 500))
}
