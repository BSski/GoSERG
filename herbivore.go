package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	"math/rand"
)

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
}

func (h *herbivore) init() {
	h.speed = speeds[h.dna[0]]
	h.bowelLength = bowelLengths[h.dna[1]]
	h.fatLimit = fatLimits[h.dna[2]]
	h.legsLength = legsLengths[h.dna[3]]
}

func (h *herbivore) draw(screen *ebiten.Image) {
	var hColor color.RGBA
	if h.energy >= h.g.s.herbivoresBreedLevel {
		hColor = color.RGBA{R: 0, G: 123, B: 51, A: 255}
	} else {
		hColor = color.RGBA{R: 0, G: 255, B: 85, A: 255}
	}
	vector.DrawFilledRect(screen, h.g.grid[h.y][h.x][0]-2, h.g.grid[h.y][h.x][1]-2, 11, 11, color.Gray{Y: 45}, false)
	vector.DrawFilledRect(screen, h.g.grid[h.y][h.x][0]-1, h.g.grid[h.y][h.x][1]-1, 9, 9, hColor, false)
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
			h.giveBirth(
				h.g,
				h.x,
				h.y,
				h.dna,
				v.dna,
			)
			break
		}
	}
}

func (_ *herbivore) giveBirth(g *game, x, y int, dna1, dna2 [4]int) {
	newDna := [4]int{}
	for i := 0; i < len(newDna); i++ {
		if rand.Float64()/100 >= g.s.mutationChance {
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
	for i := 0; i < nr; i++ {
		y := rand.Intn(g.boardSize-2) + 2
		x := rand.Intn(g.boardSize-2) + 2
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
		g.d.herbivoresSpeeds[h.dna[0]] += 1
		g.d.herbivoresBowelLengths[h.dna[1]] += 1
		g.d.herbivoresFatLimits[h.dna[2]] += 1
		g.d.herbivoresLegsLengths[h.dna[3]] += 1
	}
}

func (h *herbivore) move() {
	if int(h.g.counterPrev) == int(h.g.counter) {
		return
	}
	if int(h.g.counter)%h.speed != 0 {
		return
	}

	for i, v := range h.g.herbivoresPos[h.y][h.x] {
		if v == h {
			h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x][:i], h.g.herbivoresPos[h.y][h.x][i+1:]...)
			break
		}
	}

	var moveCost float64
	moveCost += float64(h.g.s.herbivoresMoveCost)
	moveCost += float64(h.g.s.herbivoresMoveCost) * speedCosts[h.dna[0]]
	moveCost += float64(h.g.s.herbivoresMoveCost) * bowelLengthCosts[h.dna[1]]
	moveCost += float64(h.g.s.herbivoresMoveCost) * fatLimitCosts[h.dna[2]]
	moveCost += float64(h.g.s.herbivoresMoveCost) * legsLengthCosts[h.dna[3]]
	moveCost *= legsLengths[h.dna[3]]
	h.energy -= int(moveCost)

	// Move away from the border.
	if h.x <= 1 || h.x >= h.g.boardSize || h.y <= 1 || h.y >= h.g.boardSize {
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

	vectors := h.g.c.vonNeumannPerms[rand.Intn(24)]
	// Move away from close predators.
	isPredatorClose := (len(h.g.carnivoresPos[h.y+1][h.x]) != 0 || len(h.g.carnivoresPos[h.y-1][h.x]) != 0 || len(h.g.carnivoresPos[h.y][h.x+1]) != 0 || len(h.g.carnivoresPos[h.y][h.x-1]) != 0)
	if isPredatorClose {
		h.runFromClosePredator(vectors)
		return
	}

	// Move away from distant predators.
	xSum, ySum, xPresent, yPresent := h.scanForPredators()
	if xPresent > 0 || yPresent > 0 {
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
		xSum, ySum, xPresent, yPresent := h.scanDistantMates()
		if xPresent > 0 || yPresent > 0 {
			h.chaseDistantSubject(xSum, ySum, xPresent, yPresent)
			return
		}
		h.makeRandomMove()
		return
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

func (h *herbivore) runFromClosePredator(vectors [4][2]int) {
	for _, v := range vectors {
		if len(h.g.carnivoresPos[h.y+v[1]][h.x+v[0]]) == 0 {
			h.x += v[0]
			h.y += v[1]
			h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x], h)
			return
		}
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
	if r >= 0.75 {
		h.x = h.x + 1
	} else if 0.75 > r && r >= 0.5 {
		h.x = h.x - 1
	} else if 0.5 > r && r >= 0.25 {
		h.y = h.y + 1
	} else {
		h.y = h.y - 1
	}
	h.g.herbivoresPos[h.y][h.x] = append(h.g.herbivoresPos[h.y][h.x], h)
}
