package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func newContent() *fyne.Container {
	toolbarAddIcon := widget.NewToolbarAction(theme.ContentAddIcon(), func() {
		// TODO: ...
	})
	toolbarAddIcon.ToolbarObject().Hide()
	//toolbarAddIcon.Disable()

	appBarContainer := widget.NewToolbar(
		NewToolbarTitle("PicoW LED"),
		widget.NewToolbarSpacer(),
		NewToolbarSeparator(),
		toolbarAddIcon,
	)

	devices := widget.NewAccordionItem(
		"Devices",
		container.NewCenter(widget.NewLabel("Devices Page")),
	)

	groups := widget.NewAccordionItem(
		"Groups",
		container.NewCenter(widget.NewLabel("Groups Page")),
	)

	scenes := widget.NewAccordionItem(
		"Scenes",
		container.NewCenter(widget.NewLabel("Scenes Page")),
	)

	settings := widget.NewAccordionItem(
		"Settings",
		container.NewCenter(widget.NewLabel("Settings Page")),
	)

	accordion := widget.NewAccordion(
		devices,
		groups,
		scenes,
		settings,
	)
	accordion.Open(3)

	pagesStackContainer := container.NewStack(accordion)

	pagesContainer := container.NewBorder(
		container.NewVBox(appBarContainer, widget.NewSeparator()),
		nil,
		nil,
		nil,
		pagesStackContainer,
	)

	content := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		pagesContainer,
	)

	return content
}

type ToolbarTitle struct {
	widget.Label
	Title string
}

func NewToolbarTitle(title string) *ToolbarTitle {
	return &ToolbarTitle{
		Title: title,
	}
}

func (t *ToolbarTitle) ToolbarObject() fyne.CanvasObject {
	t.SetText(t.Title)
	return t
}

type ToolbarSeparator struct {
	widget.Separator
}

func NewToolbarSeparator() *ToolbarSeparator {
	return &ToolbarSeparator{}
}

func (t *ToolbarSeparator) ToolbarObject() fyne.CanvasObject {
	return t
}
