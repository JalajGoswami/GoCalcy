package store

import (
	"fyne.io/fyne/v2/data/binding"
)

type store struct {
	CurrentInput    binding.String
	PreviousOperand binding.Float
	Operator        uint
}

var GlobalStore = &store{
	CurrentInput: binding.NewString(),
	PreviousOperand: binding.NewFloat(),
}
