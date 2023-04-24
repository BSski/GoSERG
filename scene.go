// Functions that draw entire GUI.
package main

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"strconv"
)

type scene struct {
	// Remove filling entire screen, fill only the areas that are needed.
	// It will enable you to draw only the changing parts.
	seabed *ebiten.Image
}

var sc = scene{
	seabed: getSeabedSpr(),
}

func getSeabedSpr() *ebiten.Image {
	var err error
	seabedReader := bytes.NewReader(spr.seabedBytes)
	seabedSpr, _, err := ebitenutil.NewImageFromReader(seabedReader)
	if err != nil {
		log.Fatal(err)
	}
	return seabedSpr
}

// Animation to prevent Windows from hanging the window when paused.
// Useful in approximating lag.
func (sc *scene) drawAnimation(screen *ebiten.Image, g *game) {
	text.Draw(screen, string(g.animation[g.animationCounter]), openSansRegular14, 1051, 663, color.Gray{Y: 50})
	if g.animationCounter == len(g.animation)-1 {
		g.animationCounter = 0
	}
	g.animationCounter += 1
}

func (sc *scene) drawLogo(screen *ebiten.Image) {
	text.Draw(screen, "SERG", pressStart2P, 35, 50, color.Gray{Y: 80})
	text.Draw(screen, "bsski 2023", openSansRegular11, 77, 61, color.Gray{Y: 200})
}

func (sc *scene) drawBoard(screen *ebiten.Image, g *game) {
	squareSize := float32(10)
	y := float32(21)
	x := float32(224)

	vector.DrawFilledRect(screen, x-1, y-1, float32(g.boardSize)*squareSize+1, float32(g.boardSize)*squareSize+1, color.Gray{Y: 80}, false)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(sc.seabed, options)

	for i := 0; i < g.boardSize; i++ {
		for j := 0; j < g.boardSize; j++ {
			tileColor := g.boardTilesType[i+1][j+1].color
			vector.DrawFilledRect(
				screen,
				x+float32(j)*squareSize,
				y+float32(i)*squareSize,
				squareSize-1,
				squareSize-1,
				tileColor,
				false,
			)
		}
	}
}

func (sc *scene) drawMainUILines(screen *ebiten.Image) {
	vector.StrokeLine(screen, 12, 12, 12, 657, float32(1), color.Gray{Y: 210}, false)
	vector.StrokeLine(screen, 198, 12, 198, 436, float32(1), color.Gray{Y: 210}, false)
	vector.StrokeLine(screen, 656, 12, 656, 436, float32(1), color.Gray{Y: 210}, false)
	vector.StrokeLine(screen, 848, 12, 848, 657, float32(1), color.Gray{Y: 210}, false)
	vector.StrokeLine(screen, 1049, 12, 1049, 646, float32(1), color.Gray{Y: 210}, false)
}

func (sc *scene) drawChartsBg(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 860, 44, 176, 6, color.Gray{Y: 210}, false)
	vector.DrawFilledRect(screen, 850, 50, 196, 609, color.Gray{Y: 210}, false)
	vector.DrawFilledRect(screen, 37, 469, 801, 180, color.White, false)
}

func (sc *scene) drawHistoricQuantitiesChart(screen *ebiten.Image) {
	for i := float32(0); i < 10; i++ {
		vector.StrokeLine(screen, 37, 649-20*i, 837, 649-20*i, 1, color.Gray{Y: 210}, false)
	}
	vector.StrokeLine(screen, 37, 650, 37, 467, 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, 37, 650, 837, 650, 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, 37, 468, 837, 468, 1, color.Gray{Y: 130}, false)
	vector.StrokeLine(screen, 838, 650, 838, 467, 1, color.Gray{Y: 130}, false)

	// Y axis labels.
	text.Draw(screen, fmt.Sprintf("%d", 0), openSansRegular12, 24, 653, color.Gray{Y: 50})
	for i := 1; i < 10; i++ {
		text.Draw(screen, fmt.Sprintf("%d", i*100), openSansRegular12, 14, 653-20*i, color.Gray{Y: 50})
	}

	text.Draw(screen, "Quantity", openSansRegular12, 412, 664, color.Gray{Y: 50})
}

func (sc *scene) plotHistoricQuantities(screen *ebiten.Image, g *game) {
	xHerbiChart := 36
	for i := 0; i < len(g.d.herbivoresQuantities); i++ {
		if len(g.d.herbivoresQuantities) > 800 {
			if i%(int(float64(len(g.d.herbivoresQuantities))/float64(800))+1) != 0 {
				continue
			}
			xHerbiChart += 1
			vector.DrawFilledRect(
				screen,
				float32(xHerbiChart),
				float32(648-int(g.d.herbivoresQuantities[i]/5)),
				2,
				2,
				color.RGBA{R: 0, G: 230, B: 115, A: 255},
				false,
			)
		} else {
			vector.DrawFilledRect(
				screen,
				float32(37+i),
				float32(648-int(g.d.herbivoresQuantities[i]/5)),
				2,
				2,
				color.RGBA{R: 0, G: 230, B: 115, A: 255},
				false,
			)
		}
	}
	xCarniChart := 36
	for i := 0; i < len(g.d.carnivoresQuantities); i++ {
		if len(g.d.carnivoresQuantities) > 800 {
			if i%(int(float64(len(g.d.carnivoresQuantities))/float64(800))+1) != 0 {
				continue
			}
			xCarniChart += 1
			vector.DrawFilledRect(
				screen,
				float32(xCarniChart),
				float32(648-int(g.d.carnivoresQuantities[i]/5)),
				2,
				2,
				color.RGBA{R: 255, G: 77, B: 77, A: 255},
				false,
			)
		} else {
			vector.DrawFilledRect(
				screen,
				float32(37+i),
				float32(648-int(g.d.carnivoresQuantities[i]/5)),
				2,
				2,
				color.RGBA{R: 255, G: 77, B: 77, A: 255},
				false,
			)
		}
	}
}

func (sc *scene) drawRightPanel(screen *ebiten.Image, g *game) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(860, 11)
	screen.DrawImage(g.rightPanelSprites[g.rightPanelOption], options)

	chartInfo := g.c.chartsInfo[g.rightPanelOption]

	if title, ok := chartInfo["title"].(string); ok {
		text.Draw(screen, title, openSansRegular13, 909, 58, color.RGBA{R: 50, G: 50, B: 50, A: 255})
	}

	for _, chData := range chartInfo["charts"].([]chartData) {
		label := chData.Label
		x := chData.X
		y := chData.Y
		offset := chData.Offset

		chartType := chartInfo["type"].(string)
		switch chartType {
		case "trait":
			drawTraitChartBg(screen, x, y)
		case "distribution":
			drawDistributionChartBg(screen, x, y)
			if g.rightPanelOption == 1 {
				// Herbivores icon.
				vector.DrawFilledRect(screen, 896-1, 49-1, 11, 11, color.Gray{Y: 45}, false)
				vector.DrawFilledRect(screen, 896, 49, 9, 9, color.RGBA{R: 0, G: 128, B: 96, A: 255}, false)
			} else if g.rightPanelOption == 2 {
				// Carnivores icon.
				vector.DrawFilledRect(screen, 896-1, 49-1, 11, 11, color.Gray{Y: 45}, false)
				vector.DrawFilledRect(screen, 896, 49, 9, 9, color.RGBA{R: 255, G: 77, B: 77, A: 255}, false)
			}

		}
		text.Draw(screen, label, openSansRegular11, x+offset, y-4, color.Gray{Y: 20})
	}
}

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
		text.Draw(screen, strconv.Itoa(i), openSansRegular11, chartX-11, chartY+4+15*(7-i), color.Gray{Y: 50})
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

	text.Draw(screen, "0", openSansRegular11, chartX-9, chartY+104, color.Gray{Y: 50})
	text.Draw(screen, "50", openSansRegular11, chartX-16, chartY+53, color.Gray{Y: 50})
	text.Draw(screen, "100", openSansRegular11, chartX-23, chartY+5, color.Gray{Y: 50})

	for i := 0; i < 8; i++ {
		text.Draw(screen, strconv.Itoa(i), openSansRegular11, chartX+7+20*i, chartY+111, color.Gray{Y: 50})
	}
}

func (sc *scene) plotRightPanel(screen *ebiten.Image, g *game) {
	switch g.rightPanelOption {
	case 0:
		herbivoreData := [][]float64{
			g.d.herbivoresMeanSpeeds,
			g.d.herbivoresMeanBowelLengths,
			g.d.herbivoresMeanFatLimits,
			g.d.herbivoresMeanLegsLengths,
		}
		carnivoreData := [][]float64{
			g.d.carnivoresMeanSpeeds,
			g.d.carnivoresMeanBowelLengths,
			g.d.carnivoresMeanFatLimits,
			g.d.carnivoresMeanLegsLengths,
		}
		startY := []int{180, 330, 480, 630}
		drawCharts(screen, herbivoreData, startY, color.RGBA{R: 0, G: 255, B: 85, A: 255})
		drawCharts(screen, carnivoreData, startY, color.RGBA{R: 255, G: 112, B: 77, A: 255})
	case 1:
		data := []struct {
			values [8]int
			y      int
		}{
			{g.d.herbivoresSpeeds, 88},
			{g.d.herbivoresBowelLengths, 238},
			{g.d.herbivoresFatLimits, 388},
			{g.d.herbivoresLegsLengths, 538},
		}
		for _, d := range data {
			drawDistributionBars(screen, 873, d.y, d.values, len(g.herbivores), color.RGBA{R: 0, G: 255, B: 85, A: 255})
		}
	case 2:
		data := []struct {
			values [8]int
			y      int
		}{
			{g.d.carnivoresSpeeds, 88},
			{g.d.carnivoresBowelLengths, 238},
			{g.d.carnivoresFatLimits, 388},
			{g.d.carnivoresLegsLengths, 538},
		}
		for _, d := range data {
			drawDistributionBars(screen, 873, d.y, d.values, len(g.carnivores), color.RGBA{R: 255, G: 112, B: 77, A: 255})
		}
	}
}

func (sc *scene) drawCounters(screen *ebiten.Image, g *game) {
	text.Draw(screen, "Year: "+strconv.Itoa(g.timeYear), openSansRegular12, 320, 448, color.Gray{Y: 50})
	text.Draw(screen, "Month: "+strconv.Itoa(g.timeMonth), openSansRegular12, 406, 448, color.Gray{Y: 50})
	text.Draw(screen, "Day: "+strconv.Itoa(g.timeDay), openSansRegular12, 500, 448, color.Gray{Y: 50})

	// Herbs icon.
	drawSingleHerb(screen, 714, 13, herb0Spr)
	// Herbivores icon.
	vector.DrawFilledRect(screen, 694-1, 28-1, 11, 11, color.Gray{Y: 45}, false)
	vector.DrawFilledRect(screen, 694, 28, 9, 9, color.RGBA{R: 0, G: 128, B: 96, A: 255}, false)
	// Carnivores icon.
	vector.DrawFilledRect(screen, 694-1, 43-1, 11, 11, color.Gray{Y: 45}, false)
	vector.DrawFilledRect(screen, 694, 43, 9, 9, color.RGBA{R: 255, G: 77, B: 77, A: 255}, false)

	text.Draw(screen, "HERBS: "+strconv.Itoa(len(g.herbs)), openSansRegular12, 727, 22, color.Gray{Y: 50})
	text.Draw(screen, "HERBIVORES: "+strconv.Itoa(len(g.herbivores)), openSansRegular12, 707, 37, color.Gray{Y: 50})
	text.Draw(screen, "CARNIVORES: "+strconv.Itoa(len(g.carnivores)), openSansRegular12, 707, 52, color.Gray{Y: 50})
}

func (sc *scene) drawSettings(screen *ebiten.Image, g *game) {
	// Herbs icon.
	drawSingleHerb(screen, 661, 139, herb0Spr)
	// Herbivores icon.
	vector.DrawFilledRect(screen, 661-1, 239-1, 11, 11, color.Gray{Y: 45}, false)
	vector.DrawFilledRect(screen, 661, 239, 9, 9, color.RGBA{R: 0, G: 128, B: 96, A: 255}, false)
	// Carnivores icon.
	vector.DrawFilledRect(screen, 661-1, 339-1, 11, 11, color.Gray{Y: 45}, false)
	vector.DrawFilledRect(screen, 661, 339, 9, 9, color.RGBA{R: 255, G: 77, B: 77, A: 255}, false)

	textData := []struct {
		text     string
		position [2]int
	}{
		{"SETTINGS", [2]int{659, 89}},
		{"- Simulation speed: " + strconv.Itoa(g.chosenGameSpeed), [2]int{663, 105}},
		{"- Mutation % chance: " + strconv.Itoa(int(g.s.mutationChance)), [2]int{663, 125}},
		{"HERBS", [2]int{674, 148}},
		{"- Starting number: " + strconv.Itoa(g.s.herbsStartingNr), [2]int{663, 165}},
		{"- Energy: " + strconv.Itoa(g.s.herbsEnergy), [2]int{663, 185}},
		{"- Per spawn: " + strconv.Itoa(g.s.herbsPerSpawn), [2]int{663, 205}},
		{"- Spawn rate: " + strconv.Itoa(g.s.herbsSpawnRate/2+1), [2]int{663, 225}},
		{"HERBIVORES", [2]int{674, 248}},
		{"- Starting number: " + strconv.Itoa(g.s.herbivoresStartingNr), [2]int{663, 265}},
		{"- Spawn energy: " + strconv.Itoa(g.s.herbivoresSpawnEnergy), [2]int{663, 285}},
		{"- Breeding level: " + strconv.Itoa(g.s.herbivoresBreedLevel), [2]int{663, 305}},
		{"- Move cost: " + strconv.Itoa(g.s.herbivoresMoveCost), [2]int{663, 325}},
		{"CARNIVORES", [2]int{674, 348}},
		{"- Starting number: " + strconv.Itoa(g.s.carnivoresStartingNr), [2]int{663, 365}},
		{"- Spawn energy: " + strconv.Itoa(g.s.carnivoresSpawnEnergy), [2]int{663, 385}},
		{"- Breeding level: " + strconv.Itoa(g.s.carnivoresBreedLevel), [2]int{663, 405}},
		{"- Move cost: " + strconv.Itoa(g.s.carnivoresMoveCost), [2]int{663, 425}},
	}

	for _, td := range textData {
		text.Draw(screen, td.text, openSansRegular12, td.position[0], td.position[1], color.RGBA{R: 50, G: 50, B: 50, A: 255})
	}
}

func (sc *scene) drawButtons(screen *ebiten.Image) {
	for b := range buttons {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(buttons[b].x, buttons[b].y)
		screen.DrawImage(buttons[b].sprites[buttons[b].state], options)
	}
}

func (sc *scene) drawHerbs(screen *ebiten.Image, g *game) {
	for i := 0; i < len(g.herbs); i++ {
		h := g.herbs[i]
		drawSingleHerb(screen, h.g.board[h.y][h.x][0], h.g.board[h.y][h.x][1], h.spr)
	}
}

func (sc *scene) drawHerbivores(screen *ebiten.Image, g *game) {
	for i := 0; i < len(g.herbivores); i++ {
		g.herbivores[i].draw(screen)
	}
}

func (sc *scene) drawCarnivores(screen *ebiten.Image, g *game) {
	for i := 0; i < len(g.carnivores); i++ {
		g.carnivores[i].draw(screen)
	}
}
