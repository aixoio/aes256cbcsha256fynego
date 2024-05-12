package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func render() fyne.CanvasObject {
	return container.NewBorder(
		widget.NewLabel("AES 256-Bit CBC SHA256 Util"),
		nil,
		nil,
		nil,
		container.NewGridWithColumns(
			2,
			widget.NewButton("Encrypt", func() {
				w := a.NewWindow("Encrypt")
				w.SetContent(encryptWindow())
				w.Show()
			}),
			widget.NewButton("Decrypt", func() {
				w := a.NewWindow("Decrypt")
				w.SetContent(decryptWindow())
				w.Show()
			}),
		),
	)
}

func encryptWindow() fyne.CanvasObject {
	return widget.NewLabel("Encrypt")
}

func decryptWindow() fyne.CanvasObject {
	return widget.NewLabel("Decrypt")
}
