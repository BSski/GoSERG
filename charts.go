package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func drawCharts(screen *ebiten.Image, data [][]float64, startY []int, color color.RGBA) {
	for i := 0; i < len(data); i++ {
		for v := range data[i] {
			vector.DrawFilledRect(screen, float32(870+v), float32(startY[i]-int(data[i][v]*15)), 2, 2, color, false)
		}
	}
}

func drawDistributionBars(screen *ebiten.Image, x, y int, traitValues [8]int, animalsLen int, color color.RGBA) {
	for k := 0; k < 8; k++ {
		height := int(float64(traitValues[k]) * 100.0 / float64(animalsLen))
		if height < 1 {
			height = 1
		}

		vector.DrawFilledRect(
			screen,
			float32(x)+4.0+20.0*float32(k),
			float32(y+100-int(height)),
			16,
			float32(height),
			color,
			false,
		)
	}
}
