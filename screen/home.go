package screen

import (
	"example/component"
	"example/constants"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewHomeScreen(a fyne.App) {
	w := a.NewWindow("Calculator")
	w.Resize(fyne.Size{Width: 300, Height: 400})
	w.SetFixedSize(true)

	heading := component.NewHeading()
	prevInput := component.NewPrevInputText()
	numInp := component.NewNumberInput()
	numGrid := container.NewGridWithColumns(4,
		component.NewNumBtn("C"),
		component.NewNumBtn("CE"),
		component.NewOperatorBtn(constants.Operator_Div),
		component.NewOperatorBtn(constants.Operator_Mul),
		component.NewNumBtn("7"),
		component.NewNumBtn("8"),
		component.NewNumBtn("9"),
		component.NewOperatorBtn(constants.Operator_Sub),
		component.NewNumBtn("4"),
		component.NewNumBtn("5"),
		component.NewNumBtn("6"),
		component.NewOperatorBtn(constants.Operator_Add),
		component.NewNumBtn("1"),
		component.NewNumBtn("2"),
		component.NewNumBtn("3"),
		component.NewOperatorBtn(constants.Operator_Eql),
	)

	vBox := container.NewVBox(heading, prevInput, numInp, numGrid)
	w.SetContent(vBox)
	w.Canvas().Focus(numInp)

	w.Show()
}
