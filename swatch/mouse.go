package swatch

import "fyne.io/fyne/v2/driver/desktop"

/*
	Swatch implements 

	type Mouseable interface {
		MouseDown(*MouseEvent)
		MouseUp(*MouseEvent)
	}
	
	Mouseable represents desktop mouse events that 
	can be sent to CanvasObjects
*/

func (swatch *Swatch) MouseDown(ev *desktop.MouseEvent) {
	swatch.clickHandler(swatch)
	swatch.Selected = true
	swatch.Refresh()
}

func (swatch *Swatch) MouseUp(ev *desktop.MouseEvent) {}
