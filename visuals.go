package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gonum.org/v1/gonum/mat"
	"image/color"
	"math/rand"
)

type BloodSpot struct {
	gameP *Game

	seed float64
	ttl  int
	pos  *mat.VecDense
}

func (b *BloodSpot) init(g *Game, pos [2]float64) {
	// insert a comment here and in other entities
	b.pos = mat.NewVecDense(2, []float64{pos[0], pos[1]})

	b.ttl = 16
	b.gameP = g
	b.seed = float64(rand.Intn(100)) / 100
	game := *b.gameP
	game.bloodSpots[b] = struct{}{}
	x, y := b.pos.AtVec(0), b.pos.AtVec(1)
	game.bloodSpotsPos[y][x][b] = struct{}{}
}

func (b *BloodSpot) drawMe(screen *ebiten.Image) {
	var x float64
	switch {
	case b.ttl > 12:
		x = 1.0
	case b.ttl > 8:
		x = 0.8
	case b.ttl > 4:
		x = 0.6
	default:
		x = 0.4
	}
	size := tileSize / 2 * x
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+b.pos.AtVec(0),
		boardStartY+boardBorderWidth+tileMiddlePx+b.pos.AtVec(1),
		size,
		color.NRGBA{
			R: 205,
			G: 10,
			B: 10,
			A: uint8(220 * x),
		},
	)
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+b.pos.AtVec(0)+tileSize/4,
		boardStartY+boardBorderWidth+tileMiddlePx+b.pos.AtVec(1)+tileSize/3,
		size*x*0.4,
		color.NRGBA{
			R: 255,
			G: 10,
			B: 10,
			A: uint8(200 * x * b.seed),
		},
	)
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+b.pos.AtVec(0)-tileSize/3,
		boardStartY+boardBorderWidth+tileMiddlePx+b.pos.AtVec(1)-tileSize/4,
		size*x*0.8,
		color.NRGBA{
			R: 255,
			G: 10,
			B: 10,
			A: uint8(180 * b.seed / 2),
		},
	)
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+b.pos.AtVec(0)+tileSize/4,
		boardStartY+boardBorderWidth+tileMiddlePx+b.pos.AtVec(1)-tileSize/3,
		size*x*0.8,
		color.NRGBA{
			R: 255,
			G: 10,
			B: 10,
			A: 100,
		},
	)
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+b.pos.AtVec(0)-tileSize/4,
		boardStartY+boardBorderWidth+tileMiddlePx+b.pos.AtVec(1)+tileSize/3,
		size*x*0.8,
		color.NRGBA{
			R: 255,
			G: 10,
			B: 10,
			A: 140,
		},
	)
}

func (b *BloodSpot) vanish() {
	game := *b.gameP
	x, y := b.pos.AtVec(0), b.pos.AtVec(1)
	delete(game.bloodSpotsPos[y][x], b)
	delete(game.bloodSpots, b)
}

func ageBloodSpots(g *Game) {
	for i := range g.bloodSpots {
		if i.ttl -= 1; i.ttl <= 0 {
			i.vanish()
		}
	}
}
