package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
	"math/rand"
)

var herbiFull0Spr *ebiten.Image
var herbiFullNew0Spr *ebiten.Image
var herbiHungry0Spr *ebiten.Image
var herbiHungryNew0Spr *ebiten.Image

var herbiFull1Spr *ebiten.Image
var herbiFullNew1Spr *ebiten.Image
var herbiHungry1Spr *ebiten.Image
var herbiHungryNew1Spr *ebiten.Image

var herbiFull2Spr *ebiten.Image
var herbiFullNew2Spr *ebiten.Image
var herbiHungry2Spr *ebiten.Image
var herbiHungryNew2Spr *ebiten.Image

type herbivore struct {
	g           *game
	x           int
	y           int
	energy      int
	dna         [4]int
	speed       int
	bowelLength float64
	fatLimit    int
	legsLength  float64
	age         int

	sprNr        int
	sprFull      *ebiten.Image
	sprFullNew   *ebiten.Image
	sprHungry    *ebiten.Image
	sprHungryNew *ebiten.Image
}

func (h *herbivore) init() {
	h.speed = speeds[h.dna[0]]
	h.bowelLength = bowelLengths[h.dna[1]]
	h.fatLimit = fatLimits[h.dna[2]]
	h.legsLength = legsLengths[h.dna[3]]

	switch h.sprNr {
	case 0:
		h.sprFull = herbiFull0Spr
		h.sprFullNew = herbiFullNew0Spr
		h.sprHungry = herbiHungry0Spr
		h.sprHungryNew = herbiHungryNew0Spr
	case 1:
		h.sprFull = herbiFull1Spr
		h.sprFullNew = herbiFullNew1Spr
		h.sprHungry = herbiHungry1Spr
		h.sprHungryNew = herbiHungryNew1Spr
	case 2:
		h.sprFull = herbiFull2Spr
		h.sprFullNew = herbiFullNew2Spr
		h.sprHungry = herbiHungry2Spr
		h.sprHungryNew = herbiHungryNew2Spr
	}
}

func init() {
	var err error
	herbiFull0Reader := bytes.NewReader(spr.herbiFull0Bytes)
	herbiFull0Spr, _, err = ebitenutil.NewImageFromReader(herbiFull0Reader)
	if err != nil {
		log.Fatal(err)
	}
	herbiFullNew0Reader := bytes.NewReader(spr.herbiFullNew0Bytes)
	herbiFullNew0Spr, _, err = ebitenutil.NewImageFromReader(herbiFullNew0Reader)
	if err != nil {
		log.Fatal(err)
	}
	herbiHungry0Reader := bytes.NewReader(spr.herbiHungry0Bytes)
	herbiHungry0Spr, _, err = ebitenutil.NewImageFromReader(herbiHungry0Reader)
	if err != nil {
		log.Fatal(err)
	}
	herbiHungryNew0Reader := bytes.NewReader(spr.herbiHungryNew0Bytes)
	herbiHungryNew0Spr, _, err = ebitenutil.NewImageFromReader(herbiHungryNew0Reader)
	if err != nil {
		log.Fatal(err)
	}

	herbiFull1Reader := bytes.NewReader(spr.herbiFull1Bytes)
	herbiFull1Spr, _, err = ebitenutil.NewImageFromReader(herbiFull1Reader)
	if err != nil {
		log.Fatal(err)
	}
	herbiFullNew1Reader := bytes.NewReader(spr.herbiFullNew1Bytes)
	herbiFullNew1Spr, _, err = ebitenutil.NewImageFromReader(herbiFullNew1Reader)
	if err != nil {
		log.Fatal(err)
	}
	herbiHungry1Reader := bytes.NewReader(spr.herbiHungry1Bytes)
	herbiHungry1Spr, _, err = ebitenutil.NewImageFromReader(herbiHungry1Reader)
	if err != nil {
		log.Fatal(err)
	}
	herbiHungryNew1Reader := bytes.NewReader(spr.herbiHungryNew1Bytes)
	herbiHungryNew1Spr, _, err = ebitenutil.NewImageFromReader(herbiHungryNew1Reader)
	if err != nil {
		log.Fatal(err)
	}

	herbiFull2Reader := bytes.NewReader(spr.herbiFull2Bytes)
	herbiFull2Spr, _, err = ebitenutil.NewImageFromReader(herbiFull2Reader)
	if err != nil {
		log.Fatal(err)
	}
	herbiFullNew2Reader := bytes.NewReader(spr.herbiFullNew2Bytes)
	herbiFullNew2Spr, _, err = ebitenutil.NewImageFromReader(herbiFullNew2Reader)
	if err != nil {
		log.Fatal(err)
	}
	herbiHungry2Reader := bytes.NewReader(spr.herbiHungry2Bytes)
	herbiHungry2Spr, _, err = ebitenutil.NewImageFromReader(herbiHungry2Reader)
	if err != nil {
		log.Fatal(err)
	}
	herbiHungryNew2Reader := bytes.NewReader(spr.herbiHungryNew2Bytes)
	herbiHungryNew2Spr, _, err = ebitenutil.NewImageFromReader(herbiHungryNew2Reader)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *herbivore) draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(h.g.board[h.y][h.x][0]), float64(h.g.board[h.y][h.x][1]))
	if h.energy >= h.g.s.herbivoresBreedLevel {
		screen.DrawImage(h.sprFull, options)
	} else {
		screen.DrawImage(h.sprHungry, options)
	}
}

func (h *herbivore) starve() {
	h.die()
	createHerbOnField(h.g, h.x, h.y)
}

func (h *herbivore) die() {
	for i, v := range h.g.herbivores {
		if v == h {
			h.g.herbivores = append(h.g.herbivores[:i], h.g.herbivores[i+1:]...)
			break
		}
	}
	for i, v := range h.g.herbivoresPos[h.y][h.x] {
		if v == h {
			h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x][:i], h.g.herbivoresPos[h.y][h.x][i+1:]...)
			break
		}
	}
	h.g.d.herbivoresSpeeds[h.dna[0]] -= 1
	h.g.d.herbivoresBowelLengths[h.dna[1]] -= 1
	h.g.d.herbivoresFatLimits[h.dna[2]] -= 1
	h.g.d.herbivoresLegsLengths[h.dna[3]] -= 1
}

func (h *herbivore) action() {
	if h.energy >= h.g.s.herbivoresBreedLevel {
		h.breed()
	} else {
		h.eat()
		if h.energy > h.fatLimit {
			h.energy = h.fatLimit
		}
	}
}

func (h *herbivore) breed() {
	if len(h.g.herbivoresPos[h.y][h.x]) <= 1 {
		return
	}
	for _, v := range h.g.herbivoresPos[h.y][h.x] {
		if v == h {
			continue
		}
		if v.energy >= h.g.s.herbivoresBreedLevel {
			h.energy = h.energy / 2
			v.energy = v.energy / 2
			h.giveBirth(h.g, h.x, h.y, h.dna, v.dna, h.sprNr, v.sprNr)
			break
		}
	}
}

func (_ *herbivore) giveBirth(g *game, x, y int, dna1, dna2 [4]int, spr1, spr2 int) {
	newDna := [4]int{}
	for i := 0; i < len(newDna); i++ {
		if rand.Float64() >= float64(g.s.mutationChance)/100 {
			if rand.Intn(2) >= 1 {
				newDna[i] = dna1[i]
			} else {
				newDna[i] = dna2[i]
			}
		} else {
			newDna[i] = rand.Intn(8)
		}
	}

	var newSpr int
	if rand.Intn(2) == 1 {
		newSpr = spr1
	} else {
		newSpr = spr2
	}

	h := herbivore{
		g:      g,
		x:      x,
		y:      y,
		energy: g.s.herbivoresSpawnEnergy,
		dna:    newDna,
		sprNr:  newSpr,
	}
	h.init()
	g.herbivores = append(g.herbivores, &h)
	g.herbivoresPos[y][x] = append(g.herbivoresPos[y][x], &h)
	g.d.herbivoresSpeeds[h.dna[0]] += 1
	g.d.herbivoresBowelLengths[h.dna[1]] += 1
	g.d.herbivoresFatLimits[h.dna[2]] += 1
	g.d.herbivoresLegsLengths[h.dna[3]] += 1
}

func (h *herbivore) eat() {
	if len(h.g.herbsPos[h.y][h.x]) == 0 {
		return
	}
	f := h.g.herbsPos[h.y][h.x][0]
	h.energy += int(float64(f.energy) * h.bowelLength)
	f.die()
}

func spawnHerbivore(g *game, nr int) {
	for nr > 0 {
		y := rand.Intn(g.boardSize-2) + 2
		x := rand.Intn(g.boardSize-2) + 2
		if g.boardTilesType[y][x].tileType == 0 {
			continue
		}

		dnaRange := g.c.partialDnaRange
		h := herbivore{
			g:      g,
			x:      x,
			y:      y,
			energy: g.s.herbivoresSpawnEnergy,
			dna: [4]int{
				dnaRange[rand.Intn(len(dnaRange))],
				dnaRange[rand.Intn(len(dnaRange))],
				dnaRange[rand.Intn(len(dnaRange))],
				dnaRange[rand.Intn(len(dnaRange))],
			},
			sprNr: rand.Intn(3),
		}
		h.init()
		g.herbivores = append(g.herbivores, &h)
		g.herbivoresPos[y][x] = append(g.herbivoresPos[y][x], &h)
		g.d.herbivoresSpeeds[h.dna[0]] += 1
		g.d.herbivoresBowelLengths[h.dna[1]] += 1
		g.d.herbivoresFatLimits[h.dna[2]] += 1
		g.d.herbivoresLegsLengths[h.dna[3]] += 1
		nr -= 1
	}
}

func (h *herbivore) move() {
	if int(h.g.counterPrev) == int(h.g.counter) {
		return
	}
	if int(h.g.counter)%h.speed != 0 {
		return
	}

	// Remove the herbivore's current position.
	for i, v := range h.g.herbivoresPos[h.y][h.x] {
		if v == h {
			h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x][:i], h.g.herbivoresPos[h.y][h.x][i+1:]...)
			break
		}
	}

	h.subtractMoveCostFromEnergy()

	// Move away from the border.
	if h.x <= 1 || h.x >= h.g.boardSize || h.y <= 1 || h.y >= h.g.boardSize {
		h.moveAwayFromBorder()
		return
	}

	vectors := h.g.c.vonNeumannPerms[rand.Intn(24)]

	// Move away from close predators.
	isPredatorClose := len(h.g.carnivoresPos[h.y+1][h.x]) != 0 ||
		len(h.g.carnivoresPos[h.y-1][h.x]) != 0 ||
		len(h.g.carnivoresPos[h.y][h.x+1]) != 0 ||
		len(h.g.carnivoresPos[h.y][h.x-1]) != 0

	if isPredatorClose {
		h.runFromClosePredatorOrForbiddenTiles(vectors)
		return
	}

	// Move away from distant predators.
	xSum, ySum, xPresent, yPresent := h.scanForPredators()
	if xPresent > 0 || yPresent > 0 {
		forbiddenXSum, forbiddenYSum, forbiddenXPresent, forbiddenYPresent := h.scanForForbiddenTiles()
		xSum += forbiddenXSum
		ySum += forbiddenYSum
		xPresent += forbiddenXPresent
		yPresent += forbiddenYPresent
		h.runFromDistantPredator(xSum, ySum, xPresent, yPresent)
		return
	}

	// Move towards a mate.
	if h.energy >= h.g.s.herbivoresBreedLevel {
		for t := range vectors {
			if len(h.g.herbivoresPos[h.y+vectors[t][1]][h.x+vectors[t][0]]) > 0 {
				h.x += vectors[t][0]
				h.y += vectors[t][1]
				h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x], h)
				return
			}
		}
		xSum, ySum, xPresent, yPresent = h.scanDistantMates()
		if xPresent > 0 || yPresent > 0 {
			h.chaseDistantSubject(xSum, ySum, xPresent, yPresent)
			return
		}
	}

	// Move towards food.
	for t := range vectors {
		if len(h.g.herbsPos[h.y+vectors[t][1]][h.x+vectors[t][0]]) > 0 {
			h.x += vectors[t][0]
			h.y += vectors[t][1]
			h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x], h)
			return
		}
	}
	xSum, ySum, xPresent, yPresent = h.scanDistantFood()
	if xPresent > 0 || yPresent > 0 {
		h.chaseDistantSubject(xSum, ySum, xPresent, yPresent)
		return
	}

	h.makeRandomMove()
}

func (h *herbivore) subtractMoveCostFromEnergy() {
	moveCost := float64(h.g.s.herbivoresMoveCost)
	moveCost += float64(h.g.s.herbivoresMoveCost) * speedCosts[h.dna[0]]
	moveCost += float64(h.g.s.herbivoresMoveCost) * bowelLengthCosts[h.dna[1]]
	moveCost += float64(h.g.s.herbivoresMoveCost) * fatLimitCosts[h.dna[2]]
	moveCost += float64(h.g.s.herbivoresMoveCost) * legsLengthCosts[h.dna[3]]
	moveCost *= h.legsLength
	h.energy -= int(moveCost)
}

func (h *herbivore) moveAwayFromBorder() {
	if h.x <= 1 {
		h.x += 1
	} else if h.x >= h.g.boardSize {
		h.x -= 1
	} else if h.y <= 1 {
		h.y += 1
	} else {
		h.y -= 1
	}
	h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x], h)
	return
}

func (h *herbivore) runFromClosePredatorOrForbiddenTiles(vectors [4][2]int) {
	for _, v := range vectors {
		if h.g.boardTilesType[h.y+v[1]][h.x+v[0]].tileType == 0 {
			continue
		}
		if len(h.g.carnivoresPos[h.y+v[1]][h.x+v[0]]) != 0 {
			continue
		}
		h.x += v[0]
		h.y += v[1]
		break
	}
	h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x], h)
}

func (h *herbivore) runFromDistantPredator(xSum, ySum, xPresent, yPresent int) {
	if xPresent > 0 && yPresent > 0 {
		h.y, h.x = h.runFromXY(xSum, ySum)
	} else if xPresent > 0 {
		h.y, h.x = h.runFromX(xSum)
	} else {
		h.y, h.x = h.runFromY(ySum)
	}
	h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x], h)
}

func (h *herbivore) scanForPredators() (xSum, ySum, xPresent, yPresent int) {
	for _, i := range [][2]int{
		{-2, -2}, {-2, -1}, {-2, 0}, {-2, 1}, {-2, 2},
		{-1, -2}, {-1, -1}, {-1, 1}, {-1, 2}, {0, -2},
		{0, 2}, {1, -2}, {1, -1}, {1, 1}, {1, 2},
		{2, -2}, {2, -1}, {2, 0}, {2, 1}, {2, 2},
	} {
		if len(h.g.carnivoresPos[h.y+i[1]][h.x+i[0]]) == 0 {
			continue
		}
		if i[0] != 0 {
			if i[0] < 0 {
				xSum += -1 * len(h.g.carnivoresPos[h.y+i[1]][h.x+i[0]])
			} else {
				xSum += len(h.g.carnivoresPos[h.y+i[1]][h.x+i[0]])
			}
			xPresent = 1
		}
		if i[1] != 0 {
			if i[1] < 0 {
				ySum += -1 * len(h.g.carnivoresPos[h.y+i[1]][h.x+i[0]])
			} else {
				ySum += len(h.g.carnivoresPos[h.y+i[1]][h.x+i[0]])
			}
			yPresent = 1
		}
	}
	return xSum, ySum, xPresent, yPresent
}

func (h *herbivore) scanForForbiddenTiles() (xSum, ySum, xPresent, yPresent int) {
	tempXSum, tempYSum := 0.0, 0.0
	for _, i := range [][2]int{
		{-2, -2}, {-2, -1}, {-2, 0}, {-2, 1}, {-2, 2},
		{-1, -2}, {-1, -1}, {-1, 1}, {-1, 2}, {0, -2},
		{0, 2}, {1, -2}, {1, -1}, {1, 1}, {1, 2},
		{2, -2}, {2, -1}, {2, 0}, {2, 1}, {2, 2},
	} {
		if h.g.boardTilesType[h.y+i[1]][h.x+i[0]].tileType != 0 {
			continue
		}
		if i[0] != 0 {
			if i[0] < 0 {
				tempXSum += -0.3
			} else {
				tempXSum += 0.3
			}
			xPresent = 1
		}
		if i[1] != 0 {
			if i[1] < 0 {
				tempYSum += -0.3
			} else {
				tempYSum += 0.3
			}
			yPresent = 1
		}
	}
	xSum = int(tempXSum)
	ySum = int(tempYSum)
	return xSum, ySum, xPresent, yPresent
}

func (h *herbivore) runFromXY(xSum, ySum int) (int, int) {
	if math.Abs(float64(xSum)) == math.Abs(float64(ySum)) {
		if xSum == 0 && ySum == 0 {
			r := rand.Float64()
			if r >= 0.75 {
				return h.y, h.x + 1
			} else if 0.75 > r && r >= 0.5 {
				return h.y, h.x - 1
			} else if 0.50 > r && r >= 0.25 {
				return h.y + 1, h.x
			} else {
				return h.y - 1, h.x
			}
		}
		if xSum > ySum {
			if rand.Float64() >= 0.5 {
				return h.y, h.x - 1
			} else {
				return h.y + 1, h.x
			}
		}
		if ySum > xSum {
			if rand.Float64() >= 0.5 {
				return h.y, h.x + 1
			} else {
				return h.y - 1, h.x
			}
		}
		if ySum == xSum && ySum < 0 {
			if rand.Float64() >= 0.5 {
				return h.y + 1, h.x
			} else {
				return h.y, h.x + 1
			}
		}
		if ySum == xSum && ySum > 0 {
			if rand.Float64() >= 0.5 {
				return h.y - 1, h.x
			} else {
				return h.y, h.x - 1
			}
		}
	} else if math.Abs(float64(xSum)) > math.Abs(float64(ySum)) {
		// Might need more nuanced choice here, similar to the conditions above.
		return h.y, h.x + int(math.Abs(float64(xSum))/float64(xSum)*-1)
	} else if math.Abs(float64(ySum)) > math.Abs(float64(xSum)) {
		// Might need more nuanced choice here, similar to the conditions above.
		return h.y + int(math.Abs(float64(ySum))/float64(ySum)*-1), h.x
	}
	return h.y, h.x
}

func (h *herbivore) runFromX(xSum int) (int, int) {
	if xSum == 0 {
		if rand.Float64() >= 0.5 {
			return h.y + 1, h.x
		} else {
			return h.y - 1, h.x
		}
	} else {
		return h.y, h.x + int(math.Abs(float64(xSum))/float64(xSum)*-1)
	}
}

func (h *herbivore) runFromY(ySum int) (int, int) {
	if ySum == 0 {
		if rand.Float64() >= 0.5 {
			return h.y, h.x + 1
		} else {
			return h.y, h.x - 1
		}
	} else {
		return h.y + int(math.Abs(float64(ySum))/float64(ySum)*-1), h.x
	}
}

func (h *herbivore) scanDistantMates() (xSum, ySum, xPresent, yPresent int) {
	for _, i := range [][2]int{
		{-2, -2}, {-2, -1}, {-2, 0}, {-2, 1}, {-2, 2},
		{-1, -2}, {-1, -1}, {-1, 1}, {-1, 2}, {0, -2},
		{0, 2}, {1, -2}, {1, -1}, {1, 1}, {1, 2},
		{2, -2}, {2, -1}, {2, 0}, {2, 1}, {2, 2},
	} {
		if len(h.g.herbivoresPos[h.y+i[1]][h.x+i[0]]) == 0 {
			continue
		}
		if i[0] != 0 {
			if i[0] < 0 {
				xSum += -1 * len(h.g.herbivoresPos[h.y+i[1]][h.x+i[0]])
			} else {
				xSum += len(h.g.herbivoresPos[h.y+i[1]][h.x+i[0]])
			}
			xPresent = 1
		}
		if i[1] != 0 {
			if i[1] < 0 {
				ySum += -1 * len(h.g.herbivoresPos[h.y+i[1]][h.x+i[0]])
			} else {
				ySum += len(h.g.herbivoresPos[h.y+i[1]][h.x+i[0]])
			}
			yPresent = 1
		}
	}
	return xSum, ySum, xPresent, yPresent
}

func (h *herbivore) scanDistantFood() (xSum, ySum, xPresent, yPresent int) {
	for _, i := range [][2]int{
		{-2, -2}, {-2, -1}, {-2, 0}, {-2, 1}, {-2, 2},
		{-1, -2}, {-1, -1}, {-1, 1}, {-1, 2}, {0, -2},
		{0, 2}, {1, -2}, {1, -1}, {1, 1}, {1, 2},
		{2, -2}, {2, -1}, {2, 0}, {2, 1}, {2, 2},
	} {
		if len(h.g.herbsPos[h.y+i[1]][h.x+i[0]]) == 0 {
			continue
		}
		if i[0] != 0 {
			if i[0] < 0 {
				xSum += -1 * len(h.g.herbsPos[h.y+i[1]][h.x+i[0]])
			} else {
				xSum += len(h.g.herbsPos[h.y+i[1]][h.x+i[0]])
			}
			xPresent = 1
		}
		if i[1] != 0 {
			if i[1] < 0 {
				ySum += -1 * len(h.g.herbsPos[h.y+i[1]][h.x+i[0]])
			} else {
				ySum += len(h.g.herbsPos[h.y+i[1]][h.x+i[0]])
			}
			yPresent = 1
		}
	}
	return xSum, ySum, xPresent, yPresent
}

func (h *herbivore) chaseDistantSubject(xSum, ySum, xPresent, yPresent int) {
	if xPresent == 1 && yPresent == 1 {
		h.y, h.x = h.chaseXY(xSum, ySum)
	} else if xPresent == 1 {
		h.x = h.chaseX(xSum)
	} else {
		h.y = h.chaseY(ySum)
	}
	h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x], h)
}

func (h *herbivore) chaseXY(xSum, ySum int) (y, x int) {
	if math.Abs(float64(xSum)) == math.Abs(float64(ySum)) {
		if xSum == 0 && ySum == 0 {
			r := rand.Float64()
			if r >= 0.75 {
				return h.y, h.x + 1
			} else if 0.75 > r && r >= 0.5 {
				return h.y, h.x - 1
			} else if 0.5 > r && r >= 0.25 {
				return h.y + 1, h.x
			} else {
				return h.y - 1, h.x
			}
		}
		if xSum > ySum {
			if rand.Float64() >= 0.5 {
				return h.y, h.x + 1
			} else {
				return h.y - 1, h.x
			}
		}
		if ySum > xSum {
			if rand.Float64() >= 0.5 {
				return h.y, h.x - 1
			} else {
				return h.y + 1, h.x
			}
		}
		if ySum == xSum && ySum < 0 {
			if rand.Float64() >= 0.5 {
				return h.y - 1, h.x
			} else {
				return h.y, h.x - 1
			}
		}
		if ySum == xSum && ySum > 0 {
			if rand.Float64() >= 0.5 {
				return h.y + 1, h.x
			} else {
				return h.y, h.x + 1
			}
		}
	} else if math.Abs(float64(xSum)) > math.Abs(float64(ySum)) {
		return h.y, h.x + int(math.Abs(float64(xSum))/float64(xSum))
	} else if math.Abs(float64(xSum)) < math.Abs(float64(ySum)) {
		return h.y + int(math.Abs(float64(ySum))/float64(ySum)), h.x
	}
	return h.y, h.x
}

func (h *herbivore) chaseX(xSum int) int {
	if xSum == 0 {
		if rand.Float64() >= 0.5 {
			return h.x + 1
		} else {
			return h.x - 1
		}
	} else {
		return h.x + int(math.Abs(float64(xSum))/float64(xSum))
	}
}

func (h *herbivore) chaseY(ySum int) int {
	if ySum == 0 {
		if rand.Float64() >= 0.5 {
			return h.y + 1
		} else {
			return h.y - 1
		}
	} else {
		return h.y + int(math.Abs(float64(ySum))/float64(ySum))
	}
}

func (h *herbivore) makeRandomMove() {
	r := rand.Float64()
	originalX := h.x
	originalY := h.y

	if r >= 0.75 {
		h.x = h.x + 1
	} else if 0.75 > r && r >= 0.5 {
		h.x = h.x - 1
	} else if 0.5 > r && r >= 0.25 {
		h.y = h.y + 1
	} else {
		h.y = h.y - 1
	}

	if h.g.boardTilesType[h.y][h.x].tileType == 0 {
		h.x = originalX
		h.y = originalY
		h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x], h)
		return
	}
	h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x], h)
}

func doHerbivoreActions(g *game) {
	for i := 0; i < len(g.herbivores); i++ {
		if g.herbivores[i].energy <= 0 {
			g.herbivores[i].starve()
		}
	}
	for i := 0; i < len(g.herbivores); i++ {
		g.herbivores[i].action()
		g.herbivores[i].age += 1
	}
	for i := 0; i < len(g.herbivores); i++ {
		g.herbivores[i].move()
	}
}
