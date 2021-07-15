package main

import (
	"math/rand"

	"mood/graphics"
	"mood/moods"
)

var (
	positiveMoods = []string{"positive1", "positive2", "positive3"}
	negativeMoods = []string{"negtive1", "negative2"}
	depth         = 10
)

func main() {
	moodScales := make([][]moods.Mood, 0, len(positiveMoods)+len(negativeMoods))
	for _, mood := range positiveMoods {
		moodScales = append(moodScales, moods.CreateMoodScale(mood, uint8(rand.Intn(150)), 0, 0, depth))
	}
	for _, mood := range negativeMoods {
		moodScales = append(moodScales, moods.CreateMoodScale(mood, 0, 0, uint8(rand.Intn(150)), depth))
	}
	graphics.SetupUi(moodScales)
}
