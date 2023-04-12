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

	buttons = getBtns()
}

func (g *game) Layout(_, _ int) (int, int) {
	return 1061, 670
}

func (g *game) Update() error {
	processEvents(g)

	if g.reset == true {
		g.reset = false
		g.resetGame()
	}

	g.cyclesPerSec = g.cyclesPerSecList[g.chosenCyclesPerSec]
	ebiten.SetTPS(g.cyclesPerSec)

	if g.pause {
		return nil
	}

	g.counterPrev = g.counter
	g.counter += g.s.tempo
	if int(g.counter) >= 120 {
		g.counter = 0
	}
	g.totalCyclesCounter += 1

	if int(g.counterPrev) != int(g.counter) && int(g.counter)%speeds[g.s.herbsSpawnRate] == 0 {
		spawnHerbs(g, g.s.herbsPerSpawn)
	}

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

	for i := 0; i < len(g.carnivores); i++ {
		g.carnivores[i].move()
	}

	for i := 0; i < len(g.herbivores); i++ {
		g.herbivores[i].move()
	}

	if int(g.counterPrev) != int(g.counter) {
		if len(g.herbivoresQuantities) >= 160 {
			g.herbivoresQuantities = (g.herbivoresQuantities)[1:]
		}
		g.herbivoresQuantities = append(g.herbivoresQuantities, len(g.herbivores))
		g.herbivoresTotalQuantities = append(g.herbivoresTotalQuantities, len(g.herbivores))

		if len(g.carnivoresQuantities) >= 160 {
			g.carnivoresQuantities = (g.carnivoresQuantities)[1:]
		}
		g.carnivoresQuantities = append(g.carnivoresQuantities, len(g.carnivores))
		g.carnivoresTotalQuantities = append(g.carnivoresTotalQuantities, len(g.carnivores))

		g.updateAnimalsMeanData(&g.herbivoresMeanSpeeds, len(g.herbivores), &g.herbivoresSpeeds)
		g.updateAnimalsMeanData(&g.herbivoresMeanBowelLengths, len(g.herbivores), &g.herbivoresBowelLengths)
		g.updateAnimalsMeanData(&g.herbivoresMeanFatLimits, len(g.herbivores), &g.herbivoresFatLimits)
		g.updateAnimalsMeanData(&g.herbivoresMeanLegsLengths, len(g.herbivores), &g.herbivoresLegsLengths)
		g.updateAnimalsMeanData(&g.carnivoresMeanSpeeds, len(g.carnivores), &g.carnivoresSpeeds)
		g.updateAnimalsMeanData(&g.carnivoresMeanBowelLengths, len(g.carnivores), &g.carnivoresBowelLengths)
		g.updateAnimalsMeanData(&g.carnivoresMeanFatLimits, len(g.carnivores), &g.carnivoresFatLimits)
		g.updateAnimalsMeanData(&g.carnivoresMeanLegsLengths, len(g.carnivores), &g.carnivoresLegsLengths)
	}

	return nil
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
	sc.drawCounters(screen, g)
	sc.drawSettings(screen, g)
	sc.drawButtons(screen)
	sc.drawHerbs(screen, g)
	sc.drawHerbivores(screen, g)
	sc.drawCarnivores(screen, g)

}
