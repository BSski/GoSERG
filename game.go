package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"strconv"
)

type Game struct {
	herbivores  map[*Herbivore]struct{}
	carnivores  map[*Carnivore]struct{}
	foods       map[*Food]struct{}
	meats       map[*Food]struct{}
	rottenMeats map[*Food]struct{}
	vegetables  map[*Food]struct{}

	herbivoresPos map[float64]map[float64]map[*Herbivore]struct{} // Those must be float64 to be compatible with vectors.
	carnivoresPos map[float64]map[float64]map[*Carnivore]struct{}

	meatPos        map[float64]map[float64]map[*Food]struct{}
	rottenMeatPos  map[float64]map[float64]map[*Food]struct{}
	vegetablesPos  map[float64]map[float64]map[*Food]struct{}
	meatCntP       *int
	rottenMeatCntP *int
	vegetableCntP  *int

	counter  int
	tilesPos []float64

	paused bool
}

func newGame() *Game {
	g := &Game{}
	reset(g)
	// TODO: Invoke a Reset function which empties everything and does everything from zero

	return g
}

func (g *Game) Update() error {
	checkKeybinds(g)

	if g.counter%updateInterval == 0 {
		if g.paused {
			g.counter += 1
			return nil
		}
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

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}
