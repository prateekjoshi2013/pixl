package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"github.com/prateekjoshi2013/pixl/apptype"
	"github.com/prateekjoshi2013/pixl/swatch"
	"github.com/prateekjoshi2013/pixl/ui"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	appInit.PixlWindow.ShowAndRun()
}
