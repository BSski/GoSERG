package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"strconv"
)

// Function drawing trait mean value chart's background.
func drawTraitChartBg(screen *ebiten.Image, chartX, chartY int) {
	vector.DrawFilledRect(screen, float32(chartX), float32(chartY), 162, 105, color.White, false)

	for i := 0; i < 7; i++ {
		vector.StrokeLine(screen, float32(chartX), float32(chartY+15*i), float32(chartX+161), float32(chartY+15*i), 1, color.Gray{Y: 210}, false)
	}

	vector.StrokeLine(screen, float32(chartX), float32(chartY+106), float32(chartX), float32(chartY), 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, float32(chartX-1), float32(chartY+106), float32(chartX+161), float32(chartY+106), 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, float32(chartX-1), float32(chartY), float32(chartX+161), float32(chartY), 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, float32(chartX+161+1), float32(chartY-1), float32(chartX+161+1), float32(chartY+106), 1, color.Gray{Y: 130}, false)

	for i := 0; i < 8; i++ {
		text.Draw(screen, strconv.Itoa(i), mPlus1pRegular11, chartX-11, chartY+4+15*(7-i), color.Gray{Y: 50})
	}
}

func drawDistributionChartBg(screen *ebiten.Image, chartX, chartY int) {
	vector.DrawFilledRect(screen, float32(chartX), float32(chartY), 162, 100, color.White, false)

	for i := 0; i < 11; i++ {
		vector.StrokeLine(screen, float32(chartX), float32(chartY+10*(10-i)), float32(chartX+161), float32(chartY+10*(10-i)), 1, color.Gray{Y: 210}, false)
	}
	vector.StrokeLine(screen, float32(chartX), float32(chartY+101), float32(chartX), float32(chartY), 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, float32(chartX), float32(chartY+101), float32(chartX+161), float32(chartY+101), 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, float32(chartX-1), float32(chartY), float32(chartX+161+1), float32(chartY), 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, float32(chartX+161+1), float32(chartY+101), float32(chartX+161+1), float32(chartY), 1, color.Gray{Y: 130}, false)

	text.Draw(screen, "0", mPlus1pRegular11, chartX-9, chartY+104, color.Gray{Y: 50})
	text.Draw(screen, "50", mPlus1pRegular11, chartX-16, chartY+53, color.Gray{Y: 50})
	text.Draw(screen, "100", mPlus1pRegular11, chartX-23, chartY+5, color.Gray{Y: 50})

	for i := 0; i < 8; i++ {
		text.Draw(screen, strconv.Itoa(i), mPlus1pRegular11, chartX+7+20*i, chartY+110, color.Gray{Y: 50})
	}
}

func drawCharts(screen *ebiten.Image, data [][]float64, startY []int, color color.RGBA) {
	for i := 0; i < len(data); i++ {
		for v := range data[i] {
			vector.DrawFilledRect(screen, float32(870+v), float32(startY[i]-int(data[i][v]*15)), 2, 2, color, false)
		}
	}
}

func drawDistributionBars(screen *ebiten.Image, x, y int, traitValues []int, animalsLen int, color color.RGBA) {
	traitValCounter := make(map[int]int)
	for _, value := range traitValues {
		traitValCounter[value]++
	}

	for k := 0; k < 8; k++ {
		if count, ok := traitValCounter[k]; ok {
			height := float64(count) * 100.0 / float64(animalsLen)
			if height < 1 {
				height = 1
			}

			vector.DrawFilledRect(screen, float32(x)+4.0+20.0*float32(k), float32(y+100-int(height)), 16, float32(height), color, false)
		}
	}
}
