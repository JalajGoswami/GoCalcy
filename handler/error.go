package handler

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HandleError(err error, app fyne.App) {
	timeStamp := time.Now().Format("2006-01-02 03:04:05 PM")
	fmt.Fprintln(os.Stderr, timeStamp, err)

	w := app.NewWindow("Unexpected Error")

	lb := widget.NewLabel(err.Error())
	btn := widget.NewButton("Ok", func() { app.Quit() })
	btn.Alignment = widget.ButtonAlignCenter

	vBox := container.NewVBox(lb, btn)
	w.SetContent(vBox)

	w.Show()
}
