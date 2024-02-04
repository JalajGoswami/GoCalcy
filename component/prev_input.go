package component

import (
	"example/store"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
)

func NewPrevInputText() *canvas.Text {
	txt := getPreviousOperand()
	heading := canvas.NewText(txt, color.White)
	heading.Alignment = fyne.TextAlignTrailing
	heading.TextSize = 15

	prevValue := store.GlobalStore.PreviousOperand
	prevValue.AddListener(binding.NewDataListener(func() {
		heading.Text = getPreviousOperand()
		// heading.Refresh()
	}))

	return heading
}

func getPreviousOperand() string {
	prevValue := store.GlobalStore.PreviousOperand
	num, err := prevValue.Get()
	if err != nil {
		store.ErrorChannel <- err
		return ""
	}
	if num == 0 {
		return ""
	}
	return strconv.FormatFloat(num, 'g', -1, 64)
}
