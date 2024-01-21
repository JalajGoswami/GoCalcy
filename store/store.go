package store

import (
	"fyne.io/fyne/v2/data/binding"
)

const (
	Eql uint = iota
	Add
	Sub
	Mul
	Div
)

type store struct {
	CurrentInput    binding.String
	PreviousOperand float64
	Operator        uint
}

var GlobalStore = &store{
	CurrentInput: binding.NewString(),
}
