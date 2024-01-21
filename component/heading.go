package component

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func NewHeading() *canvas.Text {
	var heading = canvas.NewText("Simple Calculator", color.White)
	heading.Alignment = fyne.TextAlignCenter
	heading.TextSize = 15
	return heading
}
