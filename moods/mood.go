package moods

import (
	"image/color"
)

type Mood struct {
	Name     string
	Depth    int
	ColorON  color.NRGBA
	ColorOFF color.NRGBA
}

func CreateMoodScale(name string, r, g, b uint8, n int) []Mood {
	res := make([]Mood, n)
	for i := 0; i < n; i++ {
		res[i] = Mood{Name: name, Depth: i + 1,
			ColorON:  color.NRGBA{r, g, b, 50 + uint8(205*(i+1)/(n+1))},
			ColorOFF: color.NRGBA{r, g, b, 25 + uint8(100*(i+1)/(n+1))},
		}
	}
	return res
}
