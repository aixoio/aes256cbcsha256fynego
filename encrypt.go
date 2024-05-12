package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/aixoio/aesbuddy"
)

func encryptWindow(w fyne.Window) fyne.CanvasObject {
	backbtn := widget.NewButtonWithIcon("Back to menu", theme.NavigateBackIcon(), func() { w.SetContent(render(w)) })
	filePathTxt := widget.NewLabel("No file selected")
	path := ""
	pwdWid := widget.NewPasswordEntry()

	pwdWid.SetPlaceHolder("Password")

	encBtn := widget.NewButton("Encrypt", func() {
		dat, err := os.ReadFile(path)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		aes_key := sha256_to_bytes([]byte(pwdWid.Text))

		enced, err := aesbuddy.AesCBCEncrypt(aes_key, dat)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		dialog.ShowFileSave(func(uc fyne.URIWriteCloser, err error) {
			if uc.URI() == nil {
				return
			}
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			_, err = uc.Write(enced)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			uc.Close()

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
		container.NewGridWithColumns(
			3,
			backbtn,
			widget.NewLabel("Encrypt"),
			widget.NewLabel(""),
		),
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
						if uc.URI() == nil {
							return
						}
						if err != nil {
							dialog.ShowError(err, w)
							return
						}
						path = uc.URI().Path()
						filePathTxt.SetText(uc.URI().Name())
						if len(pwdWid.Text) != 0 && path != "" {
							encBtn.Enable()
						} else {
							encBtn.Disable()
						}

						uc.Close()
					}, w)
				}),
			),
			widget.NewLabel("Password"),
			pwdWid,
		),
	)

}
