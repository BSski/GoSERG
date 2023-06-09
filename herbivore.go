package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
	"math/rand"
)

var herbiFullSpr *ebiten.Image
var herbiFullNewSpr *ebiten.Image
var herbiHungrySpr *ebiten.Image
var herbiHungryNewSpr *ebiten.Image

type herbivore struct {
	g            *game
	x            int
	y            int
	energy       int
	dna          [4]int
	speedDivider int
	bowelLength  float64
	fatLimit     int
	legsLength   float64
	age          int
}

func (h *herbivore) init() {
	h.speedDivider = speeds[h.dna[0]]
	h.bowelLength = bowelLengths[h.dna[1]]
	h.fatLimit = fatLimits[h.dna[2]]
	h.legsLength = legsLengths[h.dna[3]]
}

func init() {
	var err error
	herbiFullReader := bytes.NewReader(spr.herbiFullBytes)
	herbiFullSpr, _, err = ebitenutil.NewImageFromReader(herbiFullReader)
	if err != nil {
		log.Fatal(err)
	}
	herbiFullNewReader := bytes.NewReader(spr.herbiFullNewBytes)
	herbiFullNewSpr, _, err = ebitenutil.NewImageFromReader(herbiFullNewReader)
	if err != nil {
		log.Fatal(err)
	}
	herbiHungryReader := bytes.NewReader(spr.herbiHungryBytes)
	herbiHungrySpr, _, err = ebitenutil.NewImageFromReader(herbiHungryReader)
	if err != nil {
		log.Fatal(err)
	}
	herbiHungryNewReader := bytes.NewReader(spr.herbiHungryNewBytes)
	herbiHungryNewSpr, _, err = ebitenutil.NewImageFromReader(herbiHungryNewReader)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *herbivore) draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(h.g.board[h.y][h.x][0]), float64(h.g.board[h.y][h.x][1]))

	if h.age < 35 {
		if h.energy >= h.g.s.herbivoresBreedLevel {
			screen.DrawImage(herbiFullNewSpr, options)
		} else {
			screen.DrawImage(herbiHungryNewSpr, options)
		}
	} else {
		if h.energy >= h.g.s.herbivoresBreedLevel {
			screen.DrawImage(herbiFullSpr, options)
		} else {
			screen.DrawImage(herbiHungrySpr, options)
		}
	}
}

func (h *herbivore) starve() {
	h.die()
	createHerbOnField(h.g, h.x, h.y)
}

func (h *herbivore) die() {
	if h.g.boardTilesType[h.y][h.x].tileType == 0 {
		h.g.addEvent("herbivore drowned")
	}

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
	h.g.d.herbivoresSpeedsCounters[h.dna[0]] -= 1
	h.g.d.herbivoresBowelLengthsCounters[h.dna[1]] -= 1
	h.g.d.herbivoresFatLimitsCounters[h.dna[2]] -= 1
	h.g.d.herbivoresLegsLengthsCounters[h.dna[3]] -= 1
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
			h.giveBirth(h.g, h.x, h.y, h.dna, v.dna)
			break
		}
	}
}

func (_ *herbivore) giveBirth(g *game, x, y int, dna1, dna2 [4]int) {
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

	h := herbivore{
		g:      g,
		x:      x,
		y:      y,
		energy: g.s.herbivoresSpawnEnergy,
		dna:    newDna,
	}
	h.init()
	g.herbivores = append(g.herbivores, &h)
	g.herbivoresPos[y][x] = append(g.herbivoresPos[y][x], &h)
	g.d.herbivoresSpeedsCounters[h.dna[0]] += 1
	g.d.herbivoresBowelLengthsCounters[h.dna[1]] += 1
	g.d.herbivoresFatLimitsCounters[h.dna[2]] += 1
	g.d.herbivoresLegsLengthsCounters[h.dna[3]] += 1
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
		}
		h.init()
		g.herbivores = append(g.herbivores, &h)
		g.herbivoresPos[y][x] = append(g.herbivoresPos[y][x], &h)
		g.d.herbivoresSpeedsCounters[h.dna[0]] += 1
		g.d.herbivoresBowelLengthsCounters[h.dna[1]] += 1
		g.d.herbivoresFatLimitsCounters[h.dna[2]] += 1
		g.d.herbivoresLegsLengthsCounters[h.dna[3]] += 1
		nr -= 1
	}
}

func (h *herbivore) move() {
	if int(h.g.counterPrev) == int(h.g.counter) {
		return
	}
	if int(h.g.counter)%h.speedDivider != 0 {
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

	if h.g.boardTilesType[h.y][h.x].tileType == 0 {
		moveCost *= 2
	}

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
	}

	if int(g.counterPrev) != int(g.counter) {
		for i := 0; i < len(g.herbivores); i++ {
			g.herbivores[i].age += 1
		}
	}

	for i := 0; i < len(g.herbivores); i++ {
		g.herbivores[i].move()
	}
}
