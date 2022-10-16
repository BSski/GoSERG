package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gonum.org/v1/gonum/mat"
	"image/color"
	"math/rand"
)

type Carnivore struct {
	gameP *Game

	name   string
	energy int
	pos    *mat.VecDense
	color  color.NRGBA
}

func (c *Carnivore) init(g *Game, name string) {
	c.gameP = g

	cColor := color.NRGBA{
		R: 235,
		G: 30,
		B: 30,
		A: 200,
	}
	c.name = name
	c.color = cColor
	c.pos = mat.NewVecDense(
		2,
		[]float64{
			float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
			float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
		},
	)
	c.energy = rand.Intn(startingCarnivoresEnergy)

}

func (c *Carnivore) move() {
	game := *c.gameP
	animalsPos := game.carnivoresPos

	x, y := c.pos.AtVec(0), c.pos.AtVec(1)
	delete(animalsPos[y][x], c)
	direction := mat.NewVecDense(
		2,
		[]float64{
			float64((rand.Intn(3) - 1) * (tileSize + boardTilesGapWidth)),
			float64((rand.Intn(3) - 1) * (tileSize + boardTilesGapWidth)),
		},
	)

	c.pos.AddVec(c.pos, direction)
	c.pos = c.teleportAtBoundary(c.pos)

	x, y = c.pos.AtVec(0), c.pos.AtVec(1)
	animalsPos[y][x][c] = struct{}{}
}

func (c *Carnivore) teleportAtBoundary(pos *mat.VecDense) *mat.VecDense {
	// teleport at X boundaries
	if pos.AtVec(0) > lastTilePx {
		pos.SetVec(0, 0)
	} else if pos.AtVec(0) < 0 {
		pos.SetVec(0, lastTilePx)
	}

	// teleport at Y boundaries
	if pos.AtVec(1) > lastTilePx {
		pos.SetVec(1, 0)
	} else if pos.AtVec(1) < 0 {
		pos.SetVec(1, lastTilePx)
	}
	return pos
}

func (c *Carnivore) drawMe(screen *ebiten.Image) {
	var x float64
	switch {
	case c.energy > 30:
		x = 1.0
	case c.energy > 20:
		x = 0.8
	default:
		x = 0.6
	}
	size := tileSize / 2 * x
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+c.pos.AtVec(0),
		boardStartY+boardBorderWidth+tileMiddlePx+c.pos.AtVec(1),
		size,
		c.color,
	)
}

func (c *Carnivore) eat() {
	game := *c.gameP
	x, y := c.pos.AtVec(0), c.pos.AtVec(1)
	if len(game.herbivoresPos[y][x]) == 0 {
		return
	}

	animalPToEat := c.pickHerbivoreToEat(x, y)
	animalPToEat.getEaten()
}

func (c *Carnivore) pickHerbivoreToEat(x, y float64) *Herbivore {
	game := *c.gameP
	herbivoresAtThatSpot := game.herbivoresPos[y][x]

	k := rand.Intn(len(herbivoresAtThatSpot))
	i := 0
	for herbi := range game.herbivoresPos[y][x] {
		if i == k {
			return herbi
		}
		i++
	}
	return nil
}
