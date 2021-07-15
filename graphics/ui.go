package graphics

import (
	"log"
	"mood/moods"
	"mood/sqlite"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func SetupUi(moodScales [][]moods.Mood) {
	a := app.New()
	w := a.NewWindow("How are you?")
	w.SetFixedSize(true)
	a.Settings().SetTheme(theme.LightTheme())

	hBoxes := make([]fyne.CanvasObject, len(moodScales)+1)
	for i := 0; i < len(moodScales); i++ {
		hBoxes[i] = CreateMoodRow(moodScales[i])
	}
	commit := widget.NewButton("commit", func() {
		sqlite.InsertMoods(CurrentMoods)
		log.Println("commited")
		CurrentMoods = map[string]int{}
		for _, box := range hBoxes {
			box.Refresh()
		}
	})

	history := widget.NewButton("show history", func() {
		sqlite.PrintMoods()
	})

	hBoxes[len(moodScales)] = commit

	vbox := container.NewVBox(hBoxes...)

	tabs := container.NewAppTabs(
		container.NewTabItem("now", vbox),
		container.NewTabItem("history", history),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)
	w.ShowAndRun()
}
