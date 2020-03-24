package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/cvila84/jirop/pkg/fyne/icon"
	customWidget "github.com/cvila84/jirop/pkg/fyne/widget"
	"os/exec"
)

func main() {
	application := app.New()

	w := application.NewWindow("JIRA")
	w.SetIcon(icon.Jira)

	parameter := customWidget.NewOnShotEntry(func(text string) bool {
		err := exec.Command("C:\\Program Files\\Mozilla Firefox\\firefox.exe", "https://jira.gemalto.com/browse/AOTAC-"+text).Run()
		if err != nil {
			panic(err)
		}
		return true
	})

	label := widget.NewLabel("  AOTAC   -")
	labelContainer := fyne.NewContainerWithLayout(layout.NewCenterLayout())
	labelContainer.AddObject(label)

	container := fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(2))
	container.AddObject(labelContainer)
	container.AddObject(parameter)
	w.SetContent(container)
	w.SetFixedSize(true)
	w.Resize(fyne.Size{Width: 160, Height: 40})
	w.CenterOnScreen()
	w.RequestFocus()
	w.ShowAndRun()
}
