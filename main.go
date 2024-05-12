package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var a fyne.App

func main() {
	a = app.New()
	w := a.NewWindow("AES-256 CBC SHA256")

	w.SetContent(render())
	w.SetMaster()

	w.Show()
	a.Run()
}
