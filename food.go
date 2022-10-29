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

	energy         int
	pos            *mat.VecDense
	color          color.NRGBA
	foodType       string
	currentTypePos map[float64]map[float64]map[*Food]struct{}
}

func (f *Food) init(g *Game, foodType string) {
	f.gameP = g
	game := *f.gameP
	game.foods[f] = struct{}{}

	f.pos = mat.NewVecDense(
		2,
		[]float64{
			float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
			float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
		},
	)
	f.energy = rand.Intn(startingFoodEnergy)
	f.foodType = foodType

	switch foodType {
	case "meat":
		f.currentTypePos = game.meatPos
		f.color = color.NRGBA{
			R: 95,
			G: 15,
			B: 15,
			A: 230,
		}
	case "rottenMeat":
		f.currentTypePos = game.rottenMeatPos
		f.color = color.NRGBA{
			R: 70,
			G: 70,
			B: 30,
			A: 230,
		}
	case "vegetable":
		f.currentTypePos = game.vegetablesPos
		f.color = color.NRGBA{
			R: 10,
			G: 140,
			B: 10,
			A: 230,
		}
	}
	x, y := f.pos.AtVec(0), f.pos.AtVec(1)
	f.currentTypePos[y][x][f] = struct{}{}
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
	delete(f.currentTypePos[y][x], f)
	delete(game.foods, f)
}

//func spawnFood(x, y float64, energy int) {
//	foods := *f.foodsP
//	foodsPos := *f.foodsPosP
//
//	foodsPos[y][x][]
//}

//func doFoodActions(g *Game) {
//	for i := range g.foods {
//		// becomeOlder()
//	}
//}
