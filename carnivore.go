package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gonum.org/v1/gonum/mat"
	"image/color"
	"math/rand"
)

type carnivore struct {
	gameP *game

	alive  bool
	name   string
	energy int
	pos    *mat.VecDense
	color  color.NRGBA
}

func (c *carnivore) init(g *game, name string, energy any, pos [2]any) {
	c.alive = true
	cColor := color.NRGBA{
		R: 80,
		G: 0,
		B: 0,
		A: 240,
	}
	c.name = name
	c.color = cColor

	// TODO: insert a comment here and in other entities
	if pos[0] == nil && pos[1] == nil {
		c.pos = mat.NewVecDense(
			2,
			[]float64{
				float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
				float64(rand.Intn(boardWidthTiles) * (tileSize + boardTilesGapWidth)),
			},
		)
	} else {
		x, _ := pos[0].(float64)
		y, _ := pos[1].(float64)
		c.pos = mat.NewVecDense(2, []float64{x, y})
	}

	// TODO: insert a comment here and in other entities
	if energy == nil {
		c.energy = startingCarnivoresEnergy/2 + rand.Intn(startingCarnivoresEnergy)
	} else {
		energyInt, _ := energy.(int)
		c.energy = energyInt
	}

	c.gameP = g
	game := *c.gameP
	game.carnivores[c] = struct{}{}
	x, y := c.pos.AtVec(0), c.pos.AtVec(1)
	game.carnivoresPos[y][x][c] = struct{}{}
}

func (c *carnivore) move() {
	game := *c.gameP

	x, y := c.pos.AtVec(0), c.pos.AtVec(1)
	delete(game.carnivoresPos[y][x], c)

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
	game.carnivoresPos[y][x][c] = struct{}{}
}

// If an animal crosses the board boundary, teleport it to the other side.
func (c *carnivore) teleportAtBoundary(pos *mat.VecDense) *mat.VecDense {
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

func (c *carnivore) drawMe(screen *ebiten.Image) {
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

func (c *carnivore) hunt() {
	game := *c.gameP

	x, y := c.pos.AtVec(0), c.pos.AtVec(1)
	if len(game.herbivoresPos[y][x]) == 0 {
		return
	}

	animalPToEat := c.pickHerbivoreToKill(x, y)
	animalPToEat.died(animalPToEat.energy)
	bloodSpot := bloodSpot{}
	bloodSpot.init(c.gameP, [2]float64{x, y})
}

func (c *carnivore) pickHerbivoreToKill(x, y float64) *herbivore {
	game := *c.gameP
	// Picking random one could be done just by picking first one, because
	// iterating over a map results in semi-random order. It's better to do it explicitly.
	k := rand.Intn(len(game.herbivoresPos[y][x]))
	i := 0
	for herbi := range game.herbivoresPos[y][x] {
		if i == k {
			return herbi
		}
		i++
	}
	return nil
}

func (c *carnivore) eat() bool {
	game := *c.gameP

	x, y := c.pos.AtVec(0), c.pos.AtVec(1)
	if len(game.meatPos[y][x])+len(game.rottenMeatPos[y][x]) == 0 {
		return false
	}

	foodP := c.pickFoodToEat(x, y)
	if c.energy+foodP.energy < carnivoresMaxEnergy {
		c.energy += foodP.energy
		foodP.getBitten(foodP.energy)
	} else {
		foodP.getBitten(carnivoresMaxEnergy - c.energy)
		c.energy = carnivoresMaxEnergy
	}

	if c.energy > carnivoresMaxEnergy {
		panic("c.energy > carnivoresMaxEnergy")
	}
	return true
}

func (c *carnivore) pickFoodToEat(x, y float64) *food {
	game := *c.gameP
	// Picking random one could be done just by picking first one, because
	// iterating over a map results in semi-random order. It's better to do it explicitly.
	if len(game.meatPos[y][x]) > 0 {
		k := rand.Intn(len(game.meatPos[y][x]))
		i := 0
		for food := range game.meatPos[y][x] {
			if i == k {
				return food
			}
			i++
		}
		return nil
	} else {
		k := rand.Intn(len(game.rottenMeatPos[y][x]))
		i := 0
		for food := range game.rottenMeatPos[y][x] {
			if i == k {
				return food
			}
			i++
		}

		getSickChance := rand.Intn(100)
		if getSickChance > 85 {
			c.energy /= 2
		}
		return nil
	}
}

func (c *carnivore) breed() bool {
	if c.energy < carnivoresBreedThreshold {
		return false
	}

	game := *c.gameP
	x, y := c.pos.AtVec(0), c.pos.AtVec(1)
	if len(game.carnivoresPos[y][x]) == 0 {
		return false
	}

	partnerP := c.pickPartnerToBreed(x, y)
	if partnerP == nil {
		return false
	}

	childEnergy := (c.energy + partnerP.energy) / 4
	childP := &carnivore{}
	childP.init(c.gameP, "A carnivore", childEnergy, [2]any{x, y})
	c.energy /= 2
	partnerP.energy /= 2
	return true
}

func (c *carnivore) pickPartnerToBreed(x, y float64) *carnivore {
	game := *c.gameP
	var potentialPartners []*carnivore
	for carni := range game.carnivoresPos[y][x] {
		if carni.energy < carnivoresBreedThreshold {
			continue
		}
		if !carni.alive {
			continue
		}
		potentialPartners = append(potentialPartners, carni)
	}
	if len(potentialPartners) == 0 {
		return nil
	}
	k := rand.Intn(len(potentialPartners))
	return potentialPartners[k]
}

func (c *carnivore) died(energy int) {
	game := *c.gameP
	x, y := c.pos.AtVec(0), c.pos.AtVec(1)
	spawnFood(c.gameP, x, y, energy, "meat")
	delete(game.carnivoresPos[y][x], c)
	delete(game.carnivores, c)
}

func doCarnivoreActions(g *game) {
	var toDelete []*carnivore
	for i := range g.carnivores {
		if i.energy -= carnivoresMoveCost; i.energy <= 0 {
			i.alive = false
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
		i.hunt()
	}
	for _, dead := range toDelete {
		dead.died(startingCarnivoresEnergy * 0.3)
	}
}
