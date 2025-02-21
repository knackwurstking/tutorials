package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Entry Widget Test")

	entry, label := createTestEntry()

	w.SetContent(container.NewVBox(
		layout.NewSpacer(),
		label, entry,
	))

	w.ShowAndRun()
}

func createTestEntry() (*widget.Entry, *widget.Label) {
	label := widget.NewLabel("Test Entry Field:")

	entry := widget.NewEntry()
	entry.PlaceHolder = "Try to type something here..."

	return entry, label
}
