package component

import (
	"example/constants"

	"fyne.io/fyne/v2/widget"
)

func NewOperatorBtn(op constants.OperatorName) *widget.Button {
	btn := widget.NewButton(string(op), func() {
		switch op {
		case constants.Operator_Add:
			handleAdd()
		case constants.Operator_Sub:
			handleSub()
		case constants.Operator_Mul:
			handleMul()
		case constants.Operator_Div:
			handleDiv()
		}
	})
	return btn
}

func handleAdd() {

}
func handleSub() {

}
func handleMul() {

}
func handleDiv() {

}
