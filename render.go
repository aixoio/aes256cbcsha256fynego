package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func render(w fyne.Window) fyne.CanvasObject {
	return container.NewCenter(
		container.NewBorder(
			widget.NewLabel("AES 256-Bit CBC SHA256 Util"),
			nil,
			nil,
			nil,
			container.NewGridWithColumns(
				2,
				widget.NewButton("Encrypt", func() {
					w.SetContent(encryptWindow(w))
				}),
				widget.NewButton("Decrypt", func() {
					w.SetContent(decryptWindow(w))
				}),
			),
		),
	)
}
