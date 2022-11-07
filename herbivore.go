package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gonum.org/v1/gonum/mat"
	"image/color"
	"math/rand"
)

type Herbivore struct {
	gameP *Game

	toRemove bool
	name     string
	energy   int
	pos      *mat.VecDense
	color    color.NRGBA
}

func (h *Herbivore) init(g *Game, name string, energy any, pos [2]any) {
	h.toRemove = true
	hColor := color.NRGBA{
		R: 30,
		G: 235,
		B: 30,
		A: 210,
	}
	h.name = name
	h.color = hColor

	// insert a comment here and in other entities
	if pos[0] == nil && pos[1] == nil {
		h.pos = mat.NewVecDense(
			2,
			[]float64{
				float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
				float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
			},
		)
	} else {
		x, _ := pos[0].(float64)
		y, _ := pos[1].(float64)
		h.pos = mat.NewVecDense(2, []float64{x, y})
	}

	// insert a comment here and in other entities
	if energy == nil {
		h.energy = startingHerbivoresEnergy/2 + rand.Intn(startingHerbivoresEnergy)
	} else {
		energyInt, _ := energy.(int)
		h.energy = energyInt
	}

	h.gameP = g
	game := *h.gameP
	game.herbivores[h] = struct{}{}
	x, y := h.pos.AtVec(0), h.pos.AtVec(1)
	game.herbivoresPos[y][x][h] = struct{}{}
}

func (h *Herbivore) move() {
	game := *h.gameP

	x, y := h.pos.AtVec(0), h.pos.AtVec(1)
	delete(game.herbivoresPos[y][x], h)

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
	game.herbivoresPos[y][x][h] = struct{}{}
}

// If an animal crosses the board boundary, teleport it to the other side.
func (h *Herbivore) teleportAtBoundary(pos *mat.VecDense) *mat.VecDense {
	if pos.AtVec(0) > lastTilePx {
		pos.SetVec(0, 0)
	} else if pos.AtVec(0) < 0 {
		pos.SetVec(0, lastTilePx)
	}
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

func (h *Herbivore) eat() bool {
	game := *h.gameP
	fmt.Println(h.energy)
	x, y := h.pos.AtVec(0), h.pos.AtVec(1)
	if len(game.vegetablesPos[y][x]) == 0 {
		return false
	}

	foodP := h.pickFoodToEat(x, y)
	if h.energy+foodP.energy < herbivoresMaxEnergy {
		h.energy += foodP.energy
		foodP.getBitten(foodP.energy)
	} else {
		foodP.getBitten(herbivoresMaxEnergy - h.energy)
		h.energy = herbivoresMaxEnergy
	}

	if h.energy > herbivoresMaxEnergy {
		panic("h.energy > herbivoresMaxEnergy")
	}
	return true
}

func (h *Herbivore) pickFoodToEat(x, y float64) *Food {
	game := *h.gameP
	// Picking random one could be done just by picking first one, because
	// iterating over a map results in semi-random order. It's better to do it explicitly.
	k := rand.Intn(len(game.vegetablesPos[y][x]))
	i := 0
	for food := range game.vegetablesPos[y][x] {
		if i == k {
			return food
		}
		i++
	}
	return nil
}

func (h *Herbivore) breed() bool {
	if h.energy < herbivoresBreedThreshold {
		return false
	}

	game := *h.gameP
	x, y := h.pos.AtVec(0), h.pos.AtVec(1)
	if len(game.herbivoresPos[y][x]) == 0 {
		return false
	}

	partnerP := h.pickPartnerToBreed(x, y)
	if partnerP == nil {
		return false
	}

	childEnergy := (h.energy + partnerP.energy) / 4
	childP := &Herbivore{}
	childP.init(h.gameP, "A herbivore", childEnergy, [2]any{x, y})
	h.energy /= 2
	partnerP.energy /= 2
	return true
}

func (h *Herbivore) pickPartnerToBreed(x, y float64) *Herbivore {
	game := *h.gameP
	var potentialPartners []*Herbivore
	for herbi := range game.herbivoresPos[y][x] {
		if herbi.energy < herbivoresBreedThreshold {
			continue
		}
		if !herbi.toRemove {
			continue
		}
		potentialPartners = append(potentialPartners, herbi)
	}
	if len(potentialPartners) == 0 {
		return nil
	}
	k := rand.Intn(len(potentialPartners))
	return potentialPartners[k]
}

func (h *Herbivore) died(energy int) {
	game := *h.gameP
	x, y := h.pos.AtVec(0), h.pos.AtVec(1)
	spawnFood(h.gameP, x, y, energy, "meat")
	delete(game.herbivoresPos[y][x], h)
	delete(game.herbivores, h)
}

func doHerbivoreActions(g *Game) {
	var toDelete []*Herbivore
	for i := range g.herbivores {
		if i.energy -= herbivoresMoveCost; i.energy <= 0 {
			i.toRemove = false
			toDelete = append(toDelete, i)
			continue
		}
		if i.breed() {
			continue
		}
		if i.eat() {
			continue
		}
		i.move()
	}
	for _, dead := range toDelete {
		dead.died(startingHerbivoresEnergy * 0.3)
	}
}
