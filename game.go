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

	buttons map[string]*button
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

	buttons = getBtns()
}

func (g *game) Layout(_, _ int) (int, int) {
	return 1061, 670
}

// TODO: big problems with buttons: no debounce and also some are not working and throwing errors.
func (g *game) Update() error {
	processEvents(g)

	// TODO: maybe do this and delete all pause checks?
	//if g.pause {
	//	return nil
	//}

	// TODO: set g.clock.tick(g.cycles_per_sec) here.
	g.counterPrev = g.counter
	g.bigCounterPrev = g.bigCounter

	if !g.pause {
		g.counter += g.s.tempo
		if int(g.counter) >= 120 {
			g.bigCounter += 1
			g.counter = 0
		}
	}

	g.counterForFps += 1
	if g.counterForFps >= 120 {
		g.counterForFps = 0
	}

	if int(g.counterPrev) != int(g.counter) && int(g.counter)%speed[g.s.herbsSpawnRate] == 0 {
		createHerbs(g, g.s.herbsPerSpawn)
	}

	if !g.pause {
		// Check if any herbivore or carnivore starved.
		for i := 0; i < len(g.carnivores); i++ {
			if g.carnivores[i].energy <= 0 {
				g.carnivores[i].starve()
			}
		}
		for i := 0; i < len(g.herbivores); i++ {
			if g.herbivores[i].energy <= 0 {
				g.herbivores[i].starve()
			}
		}

		// Breed or eat.
		for i := 0; i < len(g.carnivores); i++ {
			g.carnivores[i].action()
			g.carnivores[i].age += 1
		}
		for i := 0; i < len(g.herbivores); i++ {
			g.herbivores[i].action()
			g.herbivores[i].age += 1
		}
	}

	// Move carnivores.
	for i := 0; i < len(g.carnivores); i++ {
		g.carnivores[i].move()
	}

	// Move herbivores.
	for i := 0; i < len(g.herbivores); i++ {
		g.herbivores[i].move()
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{Y: 239})

	// TODO: czy ten if jest potrzebny?
	//if int(g.counterForFps)%g.cyclesPerSecDividers[g.chosenCyclesPerSec] == 0 {
	//
	//}

	// Animation to prevent Windows from hanging the window when paused.
	// Useful in approximating lag.
	text.Draw(screen, string(g.animation[g.animationCounter]), mPlus1pRegular14, 1048, 648, color.Gray{Y: 50})
	if g.animationCounter == len(g.animation)-1 {
		g.animationCounter = 0
	}
	g.animationCounter += 1

	text.Draw(screen, "SERG", pressStart2P, 35, 60, color.Gray{Y: 80})
	text.Draw(screen, "bsski 2023", mPlus1pRegular11, 77, 71, color.Gray{Y: 200})

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

	chartInfo := chartsInfo[g.rightPanelOption]

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

	for b := range buttons {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(buttons[b].x, buttons[b].y)
		screen.DrawImage(buttons[b].sprites[buttons[b].state], options)
	}

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(860, 11)
	screen.DrawImage(g.rightPanelSprites[g.rightPanelOption], options)

	xHerbiChart := 36
	for i := 0; i < len(g.herbivoresTotalQuantities); i++ {
		if len(g.herbivoresTotalQuantities) > 800 {
			if i%(int(len(g.herbivoresTotalQuantities)/800)+1) == 0 {
				xHerbiChart++
				vector.DrawFilledRect(screen, float32(xHerbiChart), float32(648-int(g.herbivoresTotalQuantities[i]/5)), 2, 2, color.RGBA{R: 0, G: 230, B: 115, A: 255}, false)
			}
		} else {
			vector.DrawFilledRect(screen, float32(37+i), float32(648-int(g.herbivoresTotalQuantities[i]/5)), 2, 2, color.RGBA{R: 0, G: 230, B: 115, A: 255}, false)
		}
	}
	xCarniChart := 36
	for i := 0; i < len(g.carnivoresTotalQuantities); i++ {
		if len(g.carnivoresTotalQuantities) > 800 {
			if i%(int(len(g.carnivoresTotalQuantities)/800)+1) == 0 {
				xCarniChart++
				vector.DrawFilledRect(screen, float32(xCarniChart), float32(648-int(g.carnivoresTotalQuantities[i]/5)), 2, 2, color.RGBA{R: 255, G: 77, B: 77, A: 255}, false)
			}
		} else {
			vector.DrawFilledRect(screen, float32(37+i), float32(648-int(g.carnivoresTotalQuantities[i]/5)), 2, 2, color.RGBA{R: 255, G: 77, B: 77, A: 255}, false)
		}
	}

	// Drawing charts.
	// Quantity of herbivores and carnivores charts.
	for i := 0; i < len(g.herbivoresQuantities); i++ {
		vector.DrawFilledRect(screen, float32(29+i), float32(422-int(g.herbivoresQuantities[i]/5)), 2, 2, color.RGBA{R: 0, G: 230, B: 115, A: 255}, false)
	}
	for i := 0; i < len(g.carnivoresQuantities); i++ {
		vector.DrawFilledRect(screen, float32(29+i), float32(422-int(g.carnivoresQuantities[i]/5)), 2, 2, color.RGBA{R: 255, G: 77, B: 77, A: 255}, false)
	}

	switch g.rightPanelOption {
	case 0:
		herbivoreData := [][]float64{
			g.herbivoresMeanSpeeds,
			g.herbivoresMeanBowelLengths,
			g.herbivoresMeanFatLimits,
			g.herbivoresMeanLegsLengths,
		}
		carnivoreData := [][]float64{
			g.carnivoresMeanSpeeds,
			g.carnivoresMeanBowelLengths,
			g.carnivoresMeanFatLimits,
			g.carnivoresMeanLegsLengths,
		}
		startY := []int{180, 330, 480, 630}
		drawCharts(screen, herbivoreData, startY, color.RGBA{R: 0, G: 255, B: 85, A: 255})
		drawCharts(screen, carnivoreData, startY, color.RGBA{R: 255, G: 112, B: 77, A: 255})
	case 1:
		data := []struct {
			values []int
			y      int
		}{
			{g.herbivoresSpeeds, 89},
			{g.herbivoresBowelLengths, 239},
			{g.herbivoresFatLimits, 389},
			{g.herbivoresLegsLengths, 539},
		}
		for _, d := range data {
			drawDistributionBars(screen, 873, d.y, d.values, len(d.values), color.RGBA{R: 0, G: 255, B: 85, A: 255})
		}
	case 2:
		data := []struct {
			values []int
			y      int
		}{
			{g.carnivoresSpeeds, 89},
			{g.carnivoresBowelLengths, 239},
			{g.carnivoresFatLimits, 389},
			{g.carnivoresLegsLengths, 539},
		}
		for _, d := range data {
			drawDistributionBars(screen, 873, d.y, d.values, len(d.values), color.RGBA{R: 255, G: 112, B: 77, A: 255})
		}
	}

	// Draw all herbs and increment their age.
	for i := 0; i < len(g.herbs); i++ {
		g.herbs[i].draw(screen)
		if !g.pause {
			g.herbs[i].age += 1
		}
	}

	// TODO: change carni/herbi colors basing on their energy level.
	// Draw all carnivores and increment their age.
	for i := 0; i < len(g.carnivores); i++ {
		g.carnivores[i].draw(screen)
		if !g.pause {
			g.carnivores[i].age += 1
		}
	}

	// Draw all herbivores and increment their age.
	for i := 0; i < len(g.herbivores); i++ {
		g.herbivores[i].draw(screen)
		if !g.pause {
			g.herbivores[i].age += 1
		}
	}
	vector.DrawFilledRect(screen, 694, 43, 9, 9, color.RGBA{R: 255, G: 77, B: 77, A: 255}, false)
}
