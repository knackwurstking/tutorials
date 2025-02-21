package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type customTheme struct {
	background customColor
	foreground customColor
	shadow     customColor
	separator  customColor
	button     customColor
}

func newCustomTheme() fyne.Theme {
	return &customTheme{
		background: &customColorBackground{},
		foreground: &customColorForeground{},
		shadow:     &customColorShadow{},
		separator:  &customColorShadow{},
		button:     &customColorButton{},
	}
}

func (c *customTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s) // TODO: Add monospaced fonts
	}

	if s.Bold {
		if s.Italic {
			return resourceRecMonoCasualBoldItalic1085Ttf
		}

		return resourceRecMonoCasualRegular1085Ttf
	}

	if s.Italic {
		return resourceRecMonoCasualItalic1085Ttf
	}

	return resourceRecMonoCasualRegular1085Ttf
}

func (c *customTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		c.background.SetVariant(v)
		return c.background

	case theme.ColorNameForeground:
		c.foreground.SetVariant(v)
		return c.foreground

	case theme.ColorNameShadow:
		c.shadow.SetVariant(v)
		return c.shadow

	case theme.ColorNameScrollBar:
		return color.RGBA{0, 0, 0, 0}

	case theme.ColorNameSeparator:
		c.separator.SetVariant(v)
		return c.separator

	case theme.ColorNameButton:
		c.button.SetVariant(v)
		return c.button

	case theme.ColorNameDisabled:
		// TODO: Add disabled color for dark/light mode
		//return color.RGBA{255, 0, 0, 255}
		return theme.DefaultTheme().Color(n, v)

	default:
		log.Printf("[DEBUG][*customTheme.Color] Not handled color: %s (%d)", n, v)
		return theme.DefaultTheme().Color(n, v)
	}
}

func (c *customTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
	//icon := theme.DefaultTheme().Icon(n)
	//log.Printf("[DEBUG][*customTheme.Icon] %s: %+v", n, icon.Name())
	//return icon
}

func (c *customTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
	//size := theme.DefaultTheme().Size(n)
	//log.Printf("[DEBUG][*customTheme.Size] %s: %+v", n, size)
	//return size
}
