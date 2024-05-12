package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/aixoio/aesbuddy"
)

func render() fyne.CanvasObject {
	return container.NewCenter(
		container.NewBorder(
			widget.NewLabel("AES 256-Bit CBC SHA256 Util"),
			nil,
			nil,
			nil,
			container.NewGridWithColumns(
				2,
				widget.NewButton("Encrypt", func() {
					w := a.NewWindow("Encrypt")
					w.SetContent(encryptWindow(w))
					w.Resize(fyne.NewSize(512, 128))
					w.Show()
				}),
				widget.NewButton("Decrypt", func() {
					w := a.NewWindow("Decrypt")
					w.SetContent(decryptWindow(w))
					w.Resize(fyne.NewSize(512, 128))
					w.Show()
				}),
			),
		),
	)
}

func encryptWindow(w fyne.Window) fyne.CanvasObject {
	filePathTxt := widget.NewLabel("No file selected")
	path := ""
	pwdWid := widget.NewPasswordEntry()

	pwdWid.SetPlaceHolder("Password")

	encBtn := widget.NewButton("Encrypt", func() {
		dat, err := os.ReadFile(path)
		if err != nil {
			dialog.ShowError(err, w)
			w.Close()
		}

		aes_key := sha256_to_bytes([]byte(pwdWid.Text))

		enced, err := aesbuddy.AesCBCEncrypt(aes_key, dat)
		if err != nil {
			dialog.ShowError(err, w)
			w.Close()
		}

		dialog.ShowFileSave(func(uc fyne.URIWriteCloser, err error) {
			if err != nil {
				w.Close()
			}

			_, err = uc.Write(enced)
			if err != nil {
				dialog.ShowError(err, w)
				w.Close()
			}

			w.Close()

		}, w)

	})

	encBtn.Disable()

	pwdWid.OnChanged = func(s string) {
		if len(s) != 0 && path != "" {
			encBtn.Enable()
		} else {
			encBtn.Disable()
		}
	}

	return container.NewBorder(
		widget.NewLabel("Encrypt"),
		encBtn,
		nil,
		nil,
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel("File"),
			container.NewGridWithColumns(
				2,
				filePathTxt,
				widget.NewButton("Select file", func() {
					dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
						if err != nil {
							w.Close()
						}
						path = uc.URI().Path()
						filePathTxt.SetText(uc.URI().Name())
						if len(pwdWid.Text) != 0 && path != "" {
							encBtn.Enable()
						} else {
							encBtn.Disable()
						}
					}, w)
				}),
			),
			widget.NewLabel("Password"),
			pwdWid,
		),
	)

}

func decryptWindow(w fyne.Window) fyne.CanvasObject {
	filePathTxt := widget.NewLabel("No file selected")
	path := ""
	pwdWid := widget.NewPasswordEntry()

	pwdWid.SetPlaceHolder("Password")

	decBtn := widget.NewButton("Decrypt", func() {
		dat, err := os.ReadFile(path)
		if err != nil {
			dialog.ShowError(err, w)
			w.Close()
		}

		aes_key := sha256_to_bytes([]byte(pwdWid.Text))

		deced, err := aesbuddy.AesCBCDecrypt(aes_key, dat)
		if err != nil {
			dialog.ShowError(err, w)
			w.Close()
		}

		dialog.ShowFileSave(func(uc fyne.URIWriteCloser, err error) {
			if err != nil {
				w.Close()
			}

			_, err = uc.Write(deced)
			if err != nil {
				dialog.ShowError(err, w)
				w.Close()
			}

			w.Close()

		}, w)

	})

	decBtn.Disable()

	pwdWid.OnChanged = func(s string) {
		if len(s) != 0 && path != "" {
			decBtn.Enable()
		} else {
			decBtn.Disable()
		}
	}

	return container.NewBorder(
		widget.NewLabel("Decrypt"),
		decBtn,
		nil,
		nil,
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel("File"),
			container.NewGridWithColumns(
				2,
				filePathTxt,
				widget.NewButton("Select file", func() {
					dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
						if err != nil {
							w.Close()
						}
						path = uc.URI().Path()
						filePathTxt.SetText(uc.URI().Name())
						if len(pwdWid.Text) != 0 && path != "" {
							decBtn.Enable()
						} else {
							decBtn.Disable()
						}
					}, w)
				}),
			),
			widget.NewLabel("Password"),
			pwdWid,
		),
	)
}
