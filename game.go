package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math/rand"
	"strconv"
)

type Game struct {
	herbivores    map[*Herbivore]struct{}
	carnivores    map[*Carnivore]struct{}
	foods         map[*Food]struct{}
	herbivoresPos map[float64]map[float64]map[*Herbivore]struct{} // Those must be float64 to be compatible with vectors.
	carnivoresPos map[float64]map[float64]map[*Carnivore]struct{}

	meatPos       map[float64]map[float64]map[*Food]struct{}
	rottenMeatPos map[float64]map[float64]map[*Food]struct{}
	vegetablesPos map[float64]map[float64]map[*Food]struct{}

	counter  int
	tilesPos []float64

	paused bool
}

func newGame() *Game {
	g := &Game{}

	// TODO: Invoke a Reset function which empties everything and does everything from zero

	g.initPos()
	g.herbivores = make(map[*Herbivore]struct{})
	for i := 0; i < startingHerbivoresNr; i++ {
		newHerbiP := &Herbivore{}
		newHerbiP.init(g, "A herbivore")
	}
	g.carnivores = make(map[*Carnivore]struct{})
	for i := 0; i < startingCarnivoresNr; i++ {
		newCarniP := &Carnivore{}
		newCarniP.init(g, "A carnivore")
	}
	g.foods = make(map[*Food]struct{})
	for i := 0; i < startingFoodsNr; i++ {
		newFoodP := &Food{}

		foodTypes := map[int]string{
			0: "meat",
			1: "rottenMeat",
			2: "vegetable",
		}
		newFoodP.init(g, foodTypes[rand.Intn(len(foodTypes))])
	}
	return g
}

func (g *Game) Update() error {
	checkKeybinds(g)

	if g.counter%10 == 0 {
		if g.paused {
			g.counter += 1
			return nil
		}
		doHerbivoreActions(g)
		doCarnivoreActions(g)
		printHerbivores(g)
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

	for i := range g.foods {
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
	drawText(screen, "F: "+strconv.Itoa(len(g.foods)), 10, 65)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) initPos() {
	g.herbivoresPos = make(map[float64]map[float64]map[*Herbivore]struct{})
	g.carnivoresPos = make(map[float64]map[float64]map[*Carnivore]struct{})
	g.meatPos = make(map[float64]map[float64]map[*Food]struct{})
	g.rottenMeatPos = make(map[float64]map[float64]map[*Food]struct{})
	g.vegetablesPos = make(map[float64]map[float64]map[*Food]struct{})
	g.tilesPos = make([]float64, 0)

	// Herbivores.
	y := 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.herbivoresPos[float64(y)] == nil {
				g.herbivoresPos[float64(y)] = make(map[float64]map[*Herbivore]struct{})
			}
			g.herbivoresPos[float64(y)][float64(x)] = make(map[*Herbivore]struct{})
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	// Carnivores.
	y = 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.carnivoresPos[float64(y)] == nil {
				g.carnivoresPos[float64(y)] = make(map[float64]map[*Carnivore]struct{})
			}
			g.carnivoresPos[float64(y)][float64(x)] = make(map[*Carnivore]struct{}, 0)
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	// Meat.
	y = 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.meatPos[float64(y)] == nil {
				g.meatPos[float64(y)] = make(map[float64]map[*Food]struct{})
			}
			g.meatPos[float64(y)][float64(x)] = make(map[*Food]struct{}, 0)
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	// Rotten meat.
	y = 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.rottenMeatPos[float64(y)] == nil {
				g.rottenMeatPos[float64(y)] = make(map[float64]map[*Food]struct{})
			}
			g.rottenMeatPos[float64(y)][float64(x)] = make(map[*Food]struct{}, 0)
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	// Vegetables.
	y = 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.vegetablesPos[float64(y)] == nil {
				g.vegetablesPos[float64(y)] = make(map[float64]map[*Food]struct{})
			}
			g.vegetablesPos[float64(y)][float64(x)] = make(map[*Food]struct{}, 0)
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	// g.tilesPos is not really needed AFAIK, just used for printing positions. Maybe just create it if debugging?
	y = 0
	for i := 0; i < boardWidthTiles; i++ {
		g.tilesPos = append(g.tilesPos, float64(y))
		y += tileSize + boardTilesGapWidth
	}

}
