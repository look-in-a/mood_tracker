package graphics

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type text struct{}

func (t *text) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(70, 70)
}

func (t *text) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	delta := float32(70 / len(objects))
	pos := fyne.NewPos(0, 70/2-delta)

	for _, o := range objects {
		o.Resize(fyne.NewSize(delta, delta))
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(delta, 0))
	}
}

func CreateText(mood string) fyne.CanvasObject {
	var labels []fyne.CanvasObject
	for _, m := range mood {
		labels = append(labels, widget.NewLabelWithStyle(string(m), fyne.TextAlignCenter, fyne.TextStyle{Bold: false}))
		//canvas.NewText(string(m), color.Black))
	}

	return fyne.NewContainerWithLayout(&text{}, labels...)
}
