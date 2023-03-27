package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"strconv"
)

type game struct {
	herbivores    map[*herbivore]struct{}
	herbivoresPos map[float64]map[float64]map[*herbivore]struct{} // Those must be float64 to be compatible with vectors.

	carnivores    map[*carnivore]struct{}
	carnivoresPos map[float64]map[float64]map[*carnivore]struct{}

	foods map[*food]struct{}

	meats          map[*food]struct{}
	meatPos        map[float64]map[float64]map[*food]struct{}
	meatCntP       *int
	rottenMeats    map[*food]struct{}
	rottenMeatPos  map[float64]map[float64]map[*food]struct{}
	rottenMeatCntP *int
	vegetables     map[*food]struct{}
	vegetablesPos  map[float64]map[float64]map[*food]struct{}
	vegetableCntP  *int

	bloodSpots    map[*bloodSpot]struct{}
	bloodSpotsPos map[float64]map[float64]map[*bloodSpot]struct{}

	counter  int
	tilesPos []float64

	paused bool
}

func newGame() *game {
	g := &game{}
	return reset(g)
}

func (g *game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (g *game) Update() error {
	checkKeybinds(g)

	if g.counter%updateInterval == 0 {
		if g.paused {
			g.counter += 1
			return nil
		}
		ageBloodSpots(g)
		growVegetables(g)
		doHerbivoreActions(g)
		doCarnivoreActions(g)
		//doFoodActions(g)
		//printHerbivores(g)
		//printMeat(g)
	}
	g.counter += 1
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, boardStartX, boardStartY, boardWidthPx, boardWidthPx, color.Gray{Y: 200})

	// drawAnimal tiles
	y := boardBorderWidth
	for i := 0; i < boardWidthTiles; i++ {
		x := boardBorderWidth
		for j := 0; j < boardWidthTiles; j++ {
			ebitenutil.DrawRect(
				screen,
				float64(x+boardStartX),
				float64(y+boardStartY),
				tileSize,
				tileSize,
				color.Gray{Y: 120},
			)
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	for i := range g.bloodSpots {
		i.drawMe(screen)
	}
	for i := range g.vegetables {
		i.drawMe(screen)
	}
	for i := range g.meats {
		i.drawMe(screen)
	}
	for i := range g.rottenMeats {
		i.drawMe(screen)
	}
	for i := range g.carnivores {
		i.drawMe(screen)
	}
	for i := range g.herbivores {
		i.drawMe(screen)
	}

	// UI entities counters
	drawText(screen, "H: "+strconv.Itoa(len(g.herbivores)), 10, 25)
	drawText(screen, "C: "+strconv.Itoa(len(g.carnivores)), 10, 45)
	drawText(screen, "meat: "+strconv.Itoa(*g.meatCntP), 10, 65)
	drawText(screen, "rottenMeat: "+strconv.Itoa(*g.rottenMeatCntP), 10, 85)
	drawText(screen, "vegetable: "+strconv.Itoa(*g.vegetableCntP), 10, 105)
}
