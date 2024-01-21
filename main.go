package main

import (
	"example/handler"
	"example/helper"
	"example/screen"
	"example/store"
	"log"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	icon, err := helper.LoadResourceFromPath("./assets/icon.png")
	if err != nil {
		log.Fatalf("Static Asset Error: %v", err)
	}
	a.SetIcon(icon)
	defer a.Run()

	go func() {
		for {
			err := <-store.ErrorChannel
			handler.HandleError(err, a)
		}
	}()

	screen.NewHomeScreen(a)
}
