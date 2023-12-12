package ui

import (
	"fyne.io/fyne/v2"
	"github.com/prateekjoshi2013/pixl/apptype"
	"github.com/prateekjoshi2013/pixl/swatch"
)

type AppInit struct {
	// this is the application window where app is rendered
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
