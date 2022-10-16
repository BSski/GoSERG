package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gonum.org/v1/gonum/mat"
	"image/color"
	"math/rand"
)

type Food struct {
	gameP *Game

	energy   int
	pos      *mat.VecDense
	color    color.NRGBA
	foodType string
}

func (f *Food) init(g *Game, foodType string) {
	f.gameP = g

	f.color = color.NRGBA{
		R: 10,
		G: 140,
		B: 10,
		A: 230,
	}
	f.pos = mat.NewVecDense(
		2,
		[]float64{
			float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
			float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
		},
	)
	f.energy = rand.Intn(startingFoodEnergy)
	f.foodType = foodType
}

func (f *Food) drawMe(screen *ebiten.Image) {
	// zrob tu kolor zalezny od typu
	// moze zrob kompozycje i w Food zagniezdzaj meatFood, vegeFood, i tam przechowuj kolor?
	var x float64
	switch {
	case f.energy > 30:
		x = 1.0
	case f.energy > 20:
		x = 0.8
	default:
		x = 0.6
	}
	size := tileSize / 2 * x

	ebitenutil.DrawCircle(
		screen,
		boardStartX+boardBorderWidth+tileMiddlePx+f.pos.AtVec(0),
		boardStartY+boardBorderWidth+tileMiddlePx+f.pos.AtVec(1),
		size,
		f.color,
	)
}

func (f *Food) getEaten() {
	game := *f.gameP
	x, y := f.pos.AtVec(0), f.pos.AtVec(1)

	delete(game.foodsPos[y][x], f)
	delete(game.foods, f)
}

//func spawnAHerb(x, y float64, energy int) {
//	foods := *f.foodsP
//	foodsPos := *f.foodsPosP
//
//	foodsPos[y][x][]
//}
