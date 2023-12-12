package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

/*

	SwatchRenderer Implements  
	
	type WidgetRenderer interface {
		
		// Destroy is for internal use.
		Destroy()
		
		// Layout is a hook that is called if the widget needs to be laid out.
		// This should never call Refresh.
		Layout(Size)
		
		// MinSize returns the minimum size of the widget that is rendered by this renderer.
		MinSize() Size
		
		// Objects returns all objects that should be drawn.
		Objects() []CanvasObject
		
		// Refresh is a hook that is called if the widget has updated and needs to be redrawn.
		// This might trigger a Layout.
		Refresh()
	}

*/





type SwatchRenderer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

func (renderer *SwatchRenderer) MinSize() fyne.Size {
	return renderer.square.MinSize()
}

func (renderer *SwatchRenderer) Layout(size fyne.Size) {
	renderer.objects[0].Resize(size)
}

func (renderer *SwatchRenderer) Refresh() {
	renderer.Layout(fyne.NewSize(20, 20))
	renderer.square.FillColor = renderer.parent.Color
	if renderer.parent.Selected {
		renderer.square.StrokeWidth = 3
		renderer.square.StrokeColor = color.NRGBA{255, 255, 255, 255} // white
		renderer.objects[0] = &renderer.square
	} else {
		renderer.square.StrokeWidth = 0
		renderer.objects[0] = &renderer.square
	}
	canvas.Refresh(renderer.parent)
}

func (renderer *SwatchRenderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}

func (renderer *SwatchRenderer) Destroy() {} // can be empty just required for implementation of
