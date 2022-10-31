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

func (f *Food) init(g *Game, foodType string, energy any, pos [2]any) {
	f.gameP = g
	game := *f.gameP

	if pos[0] == nil && pos[1] == nil {
		f.pos = mat.NewVecDense(
			2,
			[]float64{
				float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
				float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
			},
		)
	} else {
		x, _ := pos[0].(float64)
		y, _ := pos[1].(float64)
		f.pos = mat.NewVecDense(2, []float64{x, y})
	}

	if energy == nil {
		f.energy = startingFoodEnergy/2 + rand.Intn(startingFoodEnergy)
	} else {
		newEnergy, _ := energy.(int)
		f.energy = newEnergy
	}

	game.foods[f] = struct{}{}
	f.foodType = foodType
	switch foodType {
	case "meat":
		game.meats[f] = struct{}{}
		*game.meatCntP += 1
		f.currentTypePos = game.meatPos
		f.color = color.NRGBA{
			R: 95,
			G: 15,
			B: 15,
			A: 230,
		}
	case "rottenMeat":
		game.rottenMeats[f] = struct{}{}
		*game.rottenMeatCntP += 1
		f.currentTypePos = game.rottenMeatPos
		f.color = color.NRGBA{
			R: 70,
			G: 70,
			B: 30,
			A: 230,
		}
	case "vegetable":
		game.vegetables[f] = struct{}{}
		*game.vegetableCntP += 1
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

func (f *Food) getBitten(biteSize int) {
	f.energy -= biteSize
	if f.energy <= 0 {
		game := *f.gameP
		x, y := f.pos.AtVec(0), f.pos.AtVec(1)
		delete(game.foods, f)
		delete(f.currentTypePos[y][x], f)

		switch f.foodType {
		case "meat":
			delete(game.meats, f)
			*game.meatCntP -= 1
		case "rottenMeat":
			delete(game.rottenMeats, f)
			*game.rottenMeatCntP -= 1
		case "vegetable":
			delete(game.vegetables, f)
			*game.vegetableCntP -= 1
		}
	}
}

func spawnFood(g *Game, x, y float64, energy int, foodType string) {
	newFoodP := &Food{}
	newFoodP.init(g, foodType, energy, [2]any{x, y})
}

//func doFoodActions(g *Game) {
//	for i := range g.foods {
//		// becomeOlder()
//	}
//}
