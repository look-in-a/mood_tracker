package graphics

import (
	"mood/moods"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type tappableIcon struct {
	widget.Icon
	tappedResource, untappedResource fyne.Resource
	tapped                           bool

	mood  *moods.Mood
	onTap func(*tappableIcon)
}

func NewTappableIcon(tappedResource, untappedResource fyne.Resource, mood *moods.Mood, onTap func(*tappableIcon)) *tappableIcon {
	icon := &tappableIcon{mood: mood, onTap: onTap}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(untappedResource)
	icon.tappedResource, icon.untappedResource = tappedResource, untappedResource
	return icon
}

var CurrentMoods = map[string]int{}

func (t *tappableIcon) Tapped(_ *fyne.PointEvent) {
	t.tapped = !t.tapped
	if t.tapped {
		CurrentMoods[t.mood.Name] = t.mood.Depth
		t.SetResource(t.tappedResource)
		t.onTap(t)
	} else {
		t.SetResource(t.untappedResource)
		delete(CurrentMoods, t.mood.Name)
	}
}

func (t *tappableIcon) TappedSecondary(_ *fyne.PointEvent) {
}

func (t *tappableIcon) Refresh() {
	t.tapped = false
	t.SetResource(t.untappedResource)
}
