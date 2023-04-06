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

	g.counterForFps += 1
	if g.counterForFps >= 120 {
		g.counterForFps = 0
	}

	if int(g.counterPrev) != int(g.counter) && int(g.counter)%speed[g.s.herbsSpawnRate] == 0 {
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
	sc.drawCountersIcons(screen)
	sc.drawSettings(screen, g)
	sc.drawButtons(screen)
	sc.drawHerbs(screen, g)
	sc.drawHerbivores(screen, g)
	sc.drawCarnivores(screen, g)
}
