package main

import (
	"GoSERG/fonts"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"strconv"
)

var (
	pressStart2P     font.Face
	mPlus1pRegular11 font.Face
	mPlus1pRegular12 font.Face
	mPlus1pRegular13 font.Face
	mPlus1pRegular14 font.Face
	mPlus1pRegular18 font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	pressStart2P, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	tt2, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mPlus1pRegular11, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    11,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mPlus1pRegular12, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mPlus1pRegular13, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    13,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mPlus1pRegular14, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    14,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mPlus1pRegular18, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    18,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *game) Layout(_, _ int) (int, int) {
	return 1061, 670
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{Y: 239})
	text.Draw(screen, "SERG", pressStart2P, 35, 72, color.Gray{Y: 80})
	text.Draw(screen, "bsski 2023", mPlus1pRegular11, 77, 83, color.Gray{Y: 200})

	squareSize := float32(10)
	y := float32(19)
	x := float32(222)

	vector.DrawFilledRect(screen, x, y, 41*squareSize-1, 41*squareSize-1, color.Gray{Y: 210}, false)

	for i := float32(1); i < 41; i++ {
		vector.StrokeLine(screen, x, y+i*squareSize, x+41*squareSize-1, y+i*squareSize, float32(1), color.Gray{Y: 239}, false)
	}
	for j := float32(1); j < 41; j++ {
		vector.StrokeLine(screen, x+j*squareSize, y, x+j*squareSize, y+41*squareSize-1, float32(1), color.Gray{Y: 239}, false)
	}

	// Charts backgrounds.
	vector.DrawFilledRect(screen, 29, 124, 161, 300, color.White, false)
	vector.DrawFilledRect(screen, 860, 44, 176, 6, color.Gray{Y: 210}, false)
	vector.DrawFilledRect(screen, 850, 50, 196, 609, color.Gray{Y: 210}, false)
	vector.DrawFilledRect(screen, 37, 459, 801, 191, color.White, false)

	// Main interface lines.
	vector.StrokeLine(screen, 12, 12, 12, 657, float32(1), color.Gray{Y: 210}, false)
	vector.StrokeLine(screen, 198, 12, 198, 436, float32(1), color.Gray{Y: 210}, false)
	vector.StrokeLine(screen, 656, 12, 656, 436, float32(1), color.Gray{Y: 210}, false)
	vector.StrokeLine(screen, 848, 12, 848, 657, float32(1), color.Gray{Y: 210}, false)
	vector.StrokeLine(screen, 1049, 12, 1049, 646, float32(1), color.Gray{Y: 210}, false)

	// Quantities chart.
	for i := float32(0); i < 30; i++ {
		if i == 18 {
			vector.StrokeLine(screen, 29, 243, 190, 243, 1, color.RGBA{R: uint8(178), G: uint8(34), B: uint8(34), A: uint8(255)}, false)
			continue
		}
		vector.StrokeLine(screen, 29, 423-10*i, 190, 423-10*i, 1, color.Gray{Y: 210}, false)
	}
	vector.StrokeLine(screen, 29, 424, 29, 123, 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, 29, 424, 190, 424, 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, 28, 124, 190, 124, 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, 190, 424, 190, 123, 1, color.Gray{Y: 130}, false)
	for i := 0; i < 15; i++ {
		if i == 10 {
			text.Draw(screen, "1.0", mPlus1pRegular11, 11, 426-20*10, color.Gray{Y: 50})
			continue
		}
		text.Draw(screen, "."+strconv.Itoa(i%10), mPlus1pRegular11, 17, 426-20*i, color.Gray{Y: 50})
	}
	text.Draw(screen, "k", mPlus1pRegular14, 19, 128, color.Gray{Y: 50})

	// All quantity history chart.
	for i := float32(0); i < 19; i++ {
		if i == 18 {
			vector.StrokeLine(screen, 37, 469, 837, 469, 1, color.RGBA{R: uint8(178), G: uint8(34), B: uint8(34), A: uint8(255)}, false)
			continue
		}
		vector.StrokeLine(screen, 37, 649-10*i, 837, 649-10*i, 1, color.Gray{Y: 210}, false)
	}
	vector.StrokeLine(screen, 37, 650, 37, 458, 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, 37, 650, 837, 650, 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, 37, 459, 837, 458, 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, 838, 650, 838, 458, 1, color.Gray{Y: 130}, false)

	type ChartData struct {
		Label  string
		X      int
		Y      int
		Offset int
	}

	var chartsInfo = map[int]map[string]interface{}{
		0: {
			"type": "trait",
			"charts": []ChartData{
				{"SPEED", 870, 75, 60},
				{"BOWEL LENGTH", 870, 225, 37},
				{"FAT LIMIT", 870, 375, 53},
				{"LEGS LENGTH", 870, 525, 44},
			},
		},
		1: {
			"title": "HERBIVORES",
			"type":  "distribution",
			"charts": []ChartData{
				{"SPEED DISTRIBUTION", 874, 88, 25},
				{"BOWEL LENGTH DISTRIBUTION", 874, 238, -2},
				{"FAT LIMIT DISTRIBUTION", 874, 388, 16},
				{"LEGS LENGTH DISTRIBUTION", 874, 538, 5},
			},
		},
		2: {
			"title": "CARNIVORES",
			"type":  "distribution",
			"charts": []ChartData{
				{"SPEED DISTRIBUTION", 874, 88, 25},
				{"BOWEL LENGTH DISTRIBUTION", 874, 238, -2},
				{"FAT LIMIT DISTRIBUTION", 874, 388, 16},
				{"LEGS LENGTH DISTRIBUTION", 874, 538, 5},
			},
		},
	}

	chartInfo := chartsInfo[g.rightPanelButtonClicked]

	if title, ok := chartInfo["title"].(string); ok {
		text.Draw(screen, title, mPlus1pRegular11, 913, 58, color.RGBA{R: 50, G: 50, B: 50, A: 255})
	}

	for _, chartData := range chartInfo["charts"].([]ChartData) {
		label := chartData.Label
		x := chartData.X
		y := chartData.Y
		offset := chartData.Offset

		chartType := chartInfo["type"].(string)
		switch chartType {
		case "trait":
			drawTraitChartBg(screen, x, y)
		case "distribution":
			drawDistributionChartBg(screen, x, y)
		}
		text.Draw(screen, label, mPlus1pRegular11, x+offset, y-4, color.Gray{Y: 20})
	}

	// Herbs icon.
	vector.DrawFilledCircle(screen, 721, 18, 3, color.RGBA{R: 34, G: 139, B: 34, A: 255}, false)
	// Herbivores icon.
	vector.DrawFilledRect(screen, 694-1, 28-1, 11, 11, color.Gray{Y: 45}, false)
	vector.DrawFilledRect(screen, 694, 28, 9, 9, color.RGBA{R: 0, G: 128, B: 96, A: 255}, false)
	// Carnivores icon.
	vector.DrawFilledRect(screen, 694-1, 43-1, 11, 11, color.Gray{Y: 45}, false)
	vector.DrawFilledRect(screen, 694, 43, 9, 9, color.RGBA{R: 255, G: 77, B: 77, A: 255}, false)

	// Settings text.
	type TextData struct {
		Text     string
		Position [2]int
	}

	text.Draw(screen, "HERBS: "+strconv.Itoa(len(g.herbs)), mPlus1pRegular12, 727, 22, color.Gray{Y: 50})
	text.Draw(screen, "HERBIVORES: "+strconv.Itoa(len(g.herbivores)), mPlus1pRegular12, 707, 37, color.Gray{Y: 50})
	text.Draw(screen, "CARNIVORES: "+strconv.Itoa(len(g.carnivores)), mPlus1pRegular12, 707, 52, color.Gray{Y: 50})

	textData := []TextData{
		{"SETTINGS", [2]int{659, 69}},
		{"- Cycles per sec: " + strconv.Itoa(g.cyclesPerSec), [2]int{663, 85}},
		{"- Tempo: " + fmt.Sprintf("%.2f", g.s.tempo), [2]int{663, 105}},
		{"- Mutation %: " + strconv.Itoa(int(g.s.mutationChance*100)), [2]int{663, 125}},
		{"HERBS", [2]int{660, 148}},
		{"- Start. number: " + strconv.Itoa(g.s.herbsStartingNr), [2]int{663, 165}},
		{"- Energy: " + strconv.Itoa(g.s.herbsEnergy), [2]int{663, 185}},
		{"- Per spawn: " + strconv.Itoa(g.s.herbsPerSpawn), [2]int{663, 205}},
		{"- Spawn rate: " + strconv.Itoa(g.s.herbsSpawnRate), [2]int{663, 225}},
		{"HERBIVORES", [2]int{660, 248}},
		{"- Start. number: " + strconv.Itoa(g.s.herbivoresStartingNr), [2]int{663, 265}},
		{"- Spawn energy: " + strconv.Itoa(g.s.herbivoresSpawnEnergy), [2]int{663, 285}},
		{"- Breed. level: " + strconv.Itoa(g.s.herbivoresBreedLevel), [2]int{663, 305}},
		{"- Move cost: " + strconv.Itoa(g.s.herbivoresMoveCost), [2]int{663, 325}},
		{"CARNIVORES", [2]int{660, 348}},
		{"- Start. number: " + strconv.Itoa(g.s.carnivoresStartingNr), [2]int{663, 365}},
		{"- Spawn energy: " + strconv.Itoa(g.s.carnivoresSpawnEnergy), [2]int{663, 385}},
		{"- Breed. level: " + strconv.Itoa(g.s.carnivoresBreedLevel), [2]int{663, 405}},
		{"- Move cost: " + strconv.Itoa(g.s.carnivoresMoveCost), [2]int{663, 425}},
	}

	for _, td := range textData {
		text.Draw(screen, td.Text, mPlus1pRegular12, td.Position[0], td.Position[1], color.RGBA{R: 50, G: 50, B: 50, A: 255})
	}

	text.Draw(screen, "NUMBER", mPlus1pRegular12, 85, 435, color.Gray{Y: 50})
}
