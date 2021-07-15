package graphics

import (
	"fyne.io/fyne/v2"
)

type horizontal struct {
	fyne.Size
}

func NewHorizontal(n int) *horizontal {
	return &horizontal{fyne.Size{float32(n), float32(n)}}
}

func (h *horizontal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w := float32(len(objects)) * h.Width
	return fyne.NewSize(w, h.Height)
}

func (h *horizontal) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	for _, o := range objects {
		o.Resize(h.Size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(h.Width, 0))
	}
}
