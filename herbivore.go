package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gonum.org/v1/gonum/mat"
	"image/color"
	"math/rand"
)

type Herbivore struct {
	gameP *Game

	name   string
	energy int
	pos    *mat.VecDense
	color  color.NRGBA
}

func (h *Herbivore) init(g *Game, name string) {
	h.gameP = g

	hColor := color.NRGBA{
		R: 30,
		G: 235,
		B: 30,
		A: 210,
	}
	h.name = name
	h.color = hColor
	h.pos = mat.NewVecDense(
		2,
		[]float64{
			float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
			float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
		},
	)
	h.energy = rand.Intn(startingHerbivoresEnergy)

}

func (h *Herbivore) move() {
	game := *h.gameP
	animalsPos := game.herbivoresPos

	x, y := h.pos.AtVec(0), h.pos.AtVec(1)
	delete(animalsPos[y][x], h)
	direction := mat.NewVecDense(
		2,
		[]float64{
			float64((rand.Intn(3) - 1) * (tileSize + boardTilesGapWidth)),
			float64((rand.Intn(3) - 1) * (tileSize + boardTilesGapWidth)),
		},
	)

	h.pos.AddVec(h.pos, direction)
	h.pos = h.teleportAtBoundary(h.pos)

	x, y = h.pos.AtVec(0), h.pos.AtVec(1)
	animalsPos[y][x][h] = struct{}{}
}

func (h *Herbivore) teleportAtBoundary(pos *mat.VecDense) *mat.VecDense {
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

func (h *Herbivore) drawMe(screen *ebiten.Image) {
	var x float64
	switch {
	case h.energy > 30:
		x = 1.0
	case h.energy > 20:
		x = 0.8
	default:
		x = 0.6
	}
	size := tileSize / 2 * x
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+h.pos.AtVec(0),
		boardStartY+boardBorderWidth+tileMiddlePx+h.pos.AtVec(1),
		size,
		h.color,
	)
}

func (h *Herbivore) getEaten() {
	x, y := h.pos.AtVec(0), h.pos.AtVec(1)
	game := *h.gameP

	// at my coordinates spawn a herb
	//spawnAHerb(x, y, h.energy)
	delete(game.herbivoresPos[y][x], h)
	delete(game.herbivores, h)
}
