package main

import (
	"GoSERG/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
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

	ebiten.SetTPS(g.cyclesPerSecList[g.chosenCyclesPerSec])

	g.counterPrev = g.counter
	g.bigCounterPrev = g.bigCounter

	if !g.pause {
		g.counter += g.s.tempo
		if int(g.counter) >= 120 {
			g.bigCounter += 1
			g.counter = 0
		}
	}

	if int(g.counterPrev) != int(g.counter) && int(g.counter)%speeds[g.s.herbsSpawnRate] == 0 {
		createHerbs(g, g.s.herbsPerSpawn)
	}

	if !g.pause {
		for i := 0; i < len(g.herbs); i++ {
			g.herbs[i].age += 1
		}

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

		for i := 0; i < len(g.carnivores); i++ {
			g.carnivores[i].action()
			g.carnivores[i].age += 1
		}
		for i := 0; i < len(g.herbivores); i++ {
			g.herbivores[i].action()
			g.herbivores[i].age += 1
		}
	}

	for i := 0; i < len(g.carnivores); i++ {
		g.carnivores[i].move()
	}

	for i := 0; i < len(g.herbivores); i++ {
		g.herbivores[i].move()
	}

	if g.reset == true {
		g.reset = false
		g.resetGame()
	}

	if int(g.counterPrev) != int(g.counter) {
		g.updateAnimalsDataHerbi(&g.herbivoresMeanSpeed, &g.herbivoresMeanSpeeds, 0, &g.herbivoresSpeeds)
		g.updateAnimalsDataHerbi(&g.herbivoresMeanBowelLength, &g.herbivoresMeanBowelLengths, 1, &g.herbivoresBowelLengths)
		g.updateAnimalsDataHerbi(&g.herbivoresMeanFatLimit, &g.herbivoresMeanFatLimits, 2, &g.herbivoresFatLimits)
		g.updateAnimalsDataHerbi(&g.herbivoresMeanLegsLength, &g.herbivoresMeanLegsLengths, 3, &g.herbivoresLegsLengths)
		g.updateAnimalsDataCarni(&g.carnivoresMeanSpeed, &g.carnivoresMeanSpeeds, 0, &g.carnivoresSpeeds)
		g.updateAnimalsDataCarni(&g.carnivoresMeanBowelLength, &g.carnivoresMeanBowelLengths, 1, &g.carnivoresBowelLengths)
		g.updateAnimalsDataCarni(&g.carnivoresMeanFatLimit, &g.carnivoresMeanFatLimits, 2, &g.carnivoresFatLimits)
		g.updateAnimalsDataCarni(&g.carnivoresMeanLegsLength, &g.carnivoresMeanLegsLengths, 3, &g.carnivoresLegsLengths)

	}
	return nil
}

// FIXME: it iterates over animals many times, make it a big single loop
func (g *game) updateAnimalsDataHerbi(
	meanVal *float64,
	meanValues *[]float64,
	gene int,
	values *[]int,
) {
	if len(*meanValues) >= 160 {
		*meanValues = (*meanValues)[1:]
	}
	var vals []int
	for _, h := range g.herbivores {
		vals = append(vals, h.dna[gene])
	}
	*values = vals
	if len(g.herbivores) > 0 {
		*meanVal = sumToFloat64(vals) / float64(len(g.herbivores))
		*meanValues = append(*meanValues, *meanVal)
	} else {
		if len(*meanValues) > 160 {
			*meanValues = (*meanValues)[1:]
		}
	}
}

func (g *game) updateAnimalsDataCarni(
	meanVal *float64,
	meanValues *[]float64,
	gene int,
	values *[]int,
) {
	if len(*meanValues) >= 160 {
		*meanValues = (*meanValues)[1:]
	}
	var vals []int
	for _, c := range g.carnivores {
		vals = append(vals, c.dna[gene])
	}
	*values = vals
	if len(g.carnivores) > 0 {
		*meanVal = sumToFloat64(vals) / float64(len(g.carnivores))
		*meanValues = append(*meanValues, *meanVal)
	} else {
		if len(*meanValues) > 160 {
			*meanValues = (*meanValues)[1:]
		}
	}
}

func sumToFloat64(vals []int) float64 {
	var sum int
	for _, v := range vals {
		sum += v
	}
	return float64(sum)
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{Y: 239})

	sc.drawAnimation(screen, g)
	sc.drawLogo(screen)
	sc.drawGrid(screen)
	sc.drawMainUILines(screen)
	sc.drawChartsBg(screen)
	sc.drawQuantitiesChart(screen)
	sc.plotQuantities(screen, g)
	sc.drawHistoricQuantitiesChart(screen)
	sc.plotHistoricQuantities(screen, g)
	sc.drawRightPanel(screen, g)
	sc.plotRightPanel(screen, g)
	sc.drawCountersIcons(screen)
	sc.drawSettings(screen, g)
	sc.drawButtons(screen)
	sc.drawHerbs(screen, g)
	sc.drawHerbivores(screen, g)
	sc.drawCarnivores(screen, g)

}
