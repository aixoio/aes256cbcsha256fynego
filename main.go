package main

import "fyne.io/fyne/v2/app"

func main() {
	a := app.New()
	w := a.NewWindow("AES-256 CBC SHA256")

	w.ShowAndRun()
}
