package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var a fyne.App

func main() {
	a = app.New()
	w := a.NewWindow("AES-256 CBC SHA256")

	w.SetContent(render(w))
	w.SetMaster()
	w.Resize(fyne.NewSize(612, 612))

	w.Show()
	a.Run()
}
