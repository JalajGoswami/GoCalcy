package component

import (
	"example/handler"
	"example/store"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type numericalEntry struct {
	widget.Entry
}

func NewEntry() *numericalEntry {
	e := &numericalEntry{Entry: widget.Entry{Wrapping: fyne.TextTruncate}}
	e.ExtendBaseWidget(e)
	return e
}

func (e *numericalEntry) TypedRune(r rune) {
	isValid := handler.ValidateNumberInput(r)
	if isValid {
		e.Entry.TypedRune(r)
	}
}

func NewNumberInput() *numericalEntry {
	bindedText := store.GlobalStore.CurrentInput
	inp := NewEntry()
	inp.Bind(bindedText)
	return inp
}
