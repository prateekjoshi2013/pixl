package swatch

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/prateekjoshi2013/pixl/apptype"
)

/*
	type Widget interface {
		CanvasObject

		// CreateRenderer returns a new WidgetRenderer for this widget.
		// This should not be called by regular code, it is used internally to render a widget.
		CreateRenderer() WidgetRenderer
	}

*/

type Swatch struct {
	// extending BaseWidget provides helpers that handle basic widget behaviours.
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh() // to update the color on screen
}

func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		clickHandler: clickHandler,
		SwatchIndex:  swatchIndex,
	}

	// Adds default BaseWidget functionality by extending widget.BaseWidget in Swatch struct
	swatch.ExtendBaseWidget(swatch)
	return swatch
}

// This should not be called by regular code, it is used internally to render a widget.
func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  swatch,
	}
}
