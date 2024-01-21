package screen

import (
	"example/component"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewHomeScreen(a fyne.App) {
	w := a.NewWindow("Calculator")
	w.Resize(fyne.Size{Width: 300, Height: 400})
	w.SetFixedSize(true)

	heading := component.NewHeading()
	numInp := component.NewNumberInput()
	numGrid := container.NewGridWithColumns(3,
		component.NewNumBtn("9"),
		component.NewNumBtn("8"),
		component.NewNumBtn("7"),
		component.NewNumBtn("6"),
		component.NewNumBtn("5"),
		component.NewNumBtn("4"),
		component.NewNumBtn("3"),
		component.NewNumBtn("2"),
		component.NewNumBtn("1"),
	)

	vBox := container.NewVBox(heading, numInp, numGrid)
	w.SetContent(vBox)
	w.Canvas().Focus(numInp)

	w.Show()
}
