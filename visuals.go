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

	b.ttl = 14
	b.gameP = g
	b.seed = float64(60+rand.Intn(30)) / 100
	game := *b.gameP
	game.bloodSpots[b] = struct{}{}
	x, y := b.pos.AtVec(0), b.pos.AtVec(1)
	game.bloodSpotsPos[y][x][b] = struct{}{}
}

func (b *BloodSpot) drawMe(screen *ebiten.Image) {
	var x float64
	switch {
	case b.ttl > 9:
		x = 1.0
	case b.ttl > 6:
		x = 0.8
	case b.ttl > 3:
		x = 0.6
	default:
		x = 0.4
	}
	size := tileSize / 2 * x
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+b.pos.AtVec(0),
		boardStartY+boardBorderWidth+tileMiddlePx+b.pos.AtVec(1),
		size*0.77,
		color.NRGBA{
			R: uint8(190 + 40*x),
			G: 10,
			B: 10,
			A: uint8(218 * x * (b.seed + 0.1)),
		},
	)
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+b.pos.AtVec(0),
		boardStartY+boardBorderWidth+tileMiddlePx+b.pos.AtVec(1),
		size*0.5,
		color.NRGBA{
			R: uint8(150 + 100*x),
			G: 10,
			B: 10,
			A: uint8(205 * b.seed),
		},
	)
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+b.pos.AtVec(0)+tileSize/4-x,
		boardStartY+boardBorderWidth+tileMiddlePx+b.pos.AtVec(1)+tileSize/4,
		size*x*0.4,
		color.NRGBA{
			R: 255,
			G: 10,
			B: 10,
			A: uint8(210 * x * (b.seed / 2)),
		},
	)
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+b.pos.AtVec(0)-tileSize/4+1-x,
		boardStartY+boardBorderWidth+tileMiddlePx+b.pos.AtVec(1)-tileSize/4+x,
		size*x*0.77,
		color.NRGBA{
			R: 255,
			G: 10,
			B: 10,
			A: uint8(180 * (x - 0.2)),
		},
	)
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+b.pos.AtVec(0)+tileSize/4-1,
		boardStartY+boardBorderWidth+tileMiddlePx+b.pos.AtVec(1)-tileSize/4-1,
		size*x*0.77,
		color.NRGBA{
			R: 255,
			G: 10,
			B: 10,
			A: uint8(160 * x * b.seed),
		},
	)
	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+b.pos.AtVec(0)-tileSize/4,
		boardStartY+boardBorderWidth+tileMiddlePx+b.pos.AtVec(1)+tileSize/4,
		size*x*0.77,
		color.NRGBA{
			R: 255,
			G: 10,
			B: 10,
			A: uint8(180 * x * b.seed),
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
