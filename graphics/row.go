package graphics

import (
	"mood/moods"

	"fyne.io/fyne/v2"
)

func CreateMoodRow(moodScale []moods.Mood) *fyne.Container {
	//TODO assert not empty
	selector := radioSelector{}

	n := len(moodScale)
	mood := moodScale[0].Name

	buttons := make([]fyne.CanvasObject, n+1)
	buttons[0] = CreateText(mood)

	for i := 0; i < n; i++ {
		onIcon, offIcon := createIcons(moodScale[i])
		resourceOff := fyne.NewStaticResource(mood, onIcon)
		resourceOn := fyne.NewStaticResource(mood, offIcon)
		icon := NewTappableIcon(resourceOff, resourceOn, &moodScale[i], selector.OnItemTap)
		selector.Add(icon)
		buttons[i+1] = icon
	}
	return fyne.NewContainerWithLayout(NewHorizontal(70), buttons...)
}
