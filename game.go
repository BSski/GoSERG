package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"strconv"
)

type Game struct {
	herbivores    map[*Herbivore]struct{}
	herbivoresPos map[float64]map[float64]map[*Herbivore]struct{} // Those must be float64 to be compatible with vectors.

	carnivores    map[*Carnivore]struct{}
	carnivoresPos map[float64]map[float64]map[*Carnivore]struct{}

	foods map[*Food]struct{}

	meats          map[*Food]struct{}
	meatPos        map[float64]map[float64]map[*Food]struct{}
	meatCntP       *int
	rottenMeats    map[*Food]struct{}
	rottenMeatPos  map[float64]map[float64]map[*Food]struct{}
	rottenMeatCntP *int
	vegetables     map[*Food]struct{}
	vegetablesPos  map[float64]map[float64]map[*Food]struct{}
	vegetableCntP  *int

	bloodSpots    map[*BloodSpot]struct{}
	bloodSpotsPos map[float64]map[float64]map[*BloodSpot]struct{}

	counter  int
	tilesPos []float64

	paused bool
}

func newGame() *Game {
	g := &Game{}
	return reset(g)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
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

func (g *Game) Draw(screen *ebiten.Image) {
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
