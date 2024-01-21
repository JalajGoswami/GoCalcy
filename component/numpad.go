package component

import (
	"example/handler"

	"fyne.io/fyne/v2/widget"
)

func NewNumBtn(n string) *widget.Button {
	btn := widget.NewButton(n, handler.HandleNumberPressed(n))
	return btn
}
