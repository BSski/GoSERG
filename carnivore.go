package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	"math/rand"
)

type carnivore struct {
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

func (c *carnivore) init() {
	c.speed = speeds[c.dna[0]]
	c.bowelLength = bowelLengths[c.dna[1]]
	c.fatLimit = fatLimits[c.dna[2]]
	c.legsLength = legsLengths[c.dna[3]]
}

func (c *carnivore) draw(screen *ebiten.Image) {
	var cColor color.RGBA
	if c.energy >= c.g.s.carnivoresBreedLevel {
		cColor = color.RGBA{R: 190, G: 46, B: 0, A: 255}
	} else {
		cColor = color.RGBA{R: 255, G: 112, B: 77, A: 255}
	}

	vector.DrawFilledRect(screen, c.g.grid[c.y][c.x][0]-1, c.g.grid[c.y][c.x][1]-1, 11, 11, color.Gray{Y: 45}, false)
	vector.DrawFilledRect(screen, c.g.grid[c.y][c.x][0], c.g.grid[c.y][c.x][1], 9, 9, cColor, false)
}

func (c *carnivore) starve() {
	c.die()
	createHerbOnField(c.g, c.x, c.y)
}

func (c *carnivore) die() {
	for i, v := range c.g.carnivores {
		if v == c {
			c.g.carnivores = append(c.g.carnivores[:i], c.g.carnivores[i+1:]...)
			break
		}
	}
	for i, v := range c.g.carnivoresPos[c.y][c.x] {
		if v == c {
			c.g.carnivoresPos[c.y][c.x] = append(c.g.carnivoresPos[c.y][c.x][:i], c.g.carnivoresPos[c.y][c.x][i+1:]...)
			break
		}
	}
	c.g.d.carnivoresSpeeds[c.dna[0]] -= 1
	c.g.d.carnivoresBowelLengths[c.dna[1]] -= 1
	c.g.d.carnivoresFatLimits[c.dna[2]] -= 1
	c.g.d.carnivoresLegsLengths[c.dna[3]] -= 1
}

func (c *carnivore) action() {
	if c.energy >= c.g.s.carnivoresBreedLevel {
		c.breed()
	} else {
		c.eat()
		if c.energy > c.fatLimit {
			c.energy = c.fatLimit
		}
	}
}

func (c *carnivore) breed() {
	if len(c.g.carnivoresPos[c.y][c.x]) <= 1 {
		return
	}
	for _, v := range c.g.carnivoresPos[c.y][c.x] {
		if v == c {
			continue
		}
		if v.energy >= c.g.s.carnivoresBreedLevel {
			c.energy = c.energy / 2
			v.energy = v.energy / 2
			c.giveBirth(
				c.g,
				c.x,
				c.y,
				c.dna,
				v.dna,
			)
			break
		}
	}
}

func (_ *carnivore) giveBirth(g *game, x, y int, dna1, dna2 [4]int) {
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
	c := carnivore{
		g:      g,
		x:      x,
		y:      y,
		energy: g.s.carnivoresSpawnEnergy,
		dna:    newDna,
	}
	c.init()
	g.carnivores = append(g.carnivores, &c)
	g.carnivoresPos[y][x] = append(g.carnivoresPos[y][x], &c)
	g.d.carnivoresSpeeds[c.dna[0]] += 1
	g.d.carnivoresBowelLengths[c.dna[1]] += 1
	g.d.carnivoresFatLimits[c.dna[2]] += 1
	g.d.carnivoresLegsLengths[c.dna[3]] += 1
}

func (c *carnivore) eat() {
	if len(c.g.herbivoresPos[c.y][c.x]) == 0 {
		return
	}
	f := c.g.herbivoresPos[c.y][c.x][0]
	c.energy += int(float64(f.energy) * c.bowelLength)
	f.die()
}

func spawnCarnivore(g *game, nr int) {
	for i := 0; i < nr; i++ {
		y := rand.Intn(g.boardSize-2) + 2
		x := rand.Intn(g.boardSize-2) + 2
		dnaRange := g.c.partialDnaRange
		c := carnivore{
			g:      g,
			x:      x,
			y:      y,
			energy: g.s.carnivoresSpawnEnergy,
			dna: [4]int{
				dnaRange[rand.Intn(len(dnaRange))],
				dnaRange[rand.Intn(len(dnaRange))],
				dnaRange[rand.Intn(len(dnaRange))],
				dnaRange[rand.Intn(len(dnaRange))],
			},
		}
		c.init()
		g.carnivores = append(g.carnivores, &c)
		g.carnivoresPos[y][x] = append(g.carnivoresPos[y][x], &c)
		g.d.carnivoresSpeeds[c.dna[0]] += 1
		g.d.carnivoresBowelLengths[c.dna[1]] += 1
		g.d.carnivoresFatLimits[c.dna[2]] += 1
		g.d.carnivoresLegsLengths[c.dna[3]] += 1
	}
}

func (c *carnivore) move() {
	if int(c.g.counterPrev) == int(c.g.counter) {
		return
	}
	if int(c.g.counter)%c.speed != 0 {
		return
	}

	for i, v := range c.g.carnivoresPos[c.y][c.x] {
		if v == c {
			c.g.carnivoresPos[c.y][c.x] = append(c.g.carnivoresPos[c.y][c.x][:i], c.g.carnivoresPos[c.y][c.x][i+1:]...)
			break
		}
	}

	var moveCost float64
	moveCost += float64(c.g.s.carnivoresMoveCost)
	moveCost += float64(c.g.s.carnivoresMoveCost) * speedCosts[c.dna[0]]
	moveCost += float64(c.g.s.carnivoresMoveCost) * bowelLengthCosts[c.dna[1]]
	moveCost += float64(c.g.s.carnivoresMoveCost) * fatLimitCosts[c.dna[2]]
	moveCost += float64(c.g.s.carnivoresMoveCost) * legsLengthCosts[c.dna[3]]
	moveCost *= legsLengths[c.dna[3]]
	c.energy -= int(moveCost)

	// Move away from the border.
	if c.x <= 1 || c.x >= c.g.boardSize || c.y <= 1 || c.y >= c.g.boardSize {
		if c.x <= 1 {
			c.x += 1
		} else if c.x >= c.g.boardSize {
			c.x -= 1
		} else if c.y <= 1 {
			c.y += 1
		} else {
			c.y -= 1
		}
		c.g.carnivoresPos[c.y][c.x] = append(c.g.carnivoresPos[c.y][c.x], c)
		return
	}

	vectors := c.g.c.vonNeumannPerms[rand.Intn(24)]

	// Move towards a mate.
	if c.energy >= c.g.s.carnivoresBreedLevel {
		for t := range vectors {
			if len(c.g.carnivoresPos[c.y+vectors[t][1]][c.x+vectors[t][0]]) > 0 {
				c.x += vectors[t][0]
				c.y += vectors[t][1]
				c.g.carnivoresPos[c.y][c.x] = append(c.g.carnivoresPos[c.y][c.x], c)
				return
			}
		}
		xSum, ySum, xPresent, yPresent := c.scanDistantMates()
		if xPresent > 0 || yPresent > 0 {
			c.chaseDistantSubject(xSum, ySum, xPresent, yPresent)
			return
		}
		c.makeRandomMove()
		return
	}

	// Move towards prey.
	for t := range vectors {
		if len(c.g.herbivoresPos[c.y+vectors[t][1]][c.x+vectors[t][0]]) > 0 {
			c.x += vectors[t][0]
			c.y += vectors[t][1]
			c.g.carnivoresPos[c.y][c.x] = append(c.g.carnivoresPos[c.y][c.x], c)
			return
		}
	}
	xSum, ySum, xPresent, yPresent := c.scanDistantFood()
	if xPresent > 0 || yPresent > 0 {
		c.chaseDistantSubject(xSum, ySum, xPresent, yPresent)
		return
	}

	c.makeRandomMove()
}

func (c *carnivore) scanDistantMates() (xSum, ySum, xPresent, yPresent int) {
	for _, i := range [][2]int{
		{-2, -2}, {-2, -1}, {-2, 0}, {-2, 1}, {-2, 2},
		{-1, -2}, {-1, -1}, {-1, 1}, {-1, 2}, {0, -2},
		{0, 2}, {1, -2}, {1, -1}, {1, 1}, {1, 2},
		{2, -2}, {2, -1}, {2, 0}, {2, 1}, {2, 2},
	} {
		if len(c.g.carnivoresPos[c.y+i[1]][c.x+i[0]]) == 0 {
			continue
		}
		if i[0] != 0 {
			if i[0] < 0 {
				xSum += -1 * len(c.g.carnivoresPos[c.y+i[1]][c.x+i[0]])
			} else {
				xSum += len(c.g.carnivoresPos[c.y+i[1]][c.x+i[0]])
			}
			xPresent = 1
		}
		if i[1] != 0 {
			if i[1] < 0 {
				ySum += -1 * len(c.g.carnivoresPos[c.y+i[1]][c.x+i[0]])
			} else {
				ySum += len(c.g.carnivoresPos[c.y+i[1]][c.x+i[0]])
			}
			yPresent = 1
		}
	}
	return xSum, ySum, xPresent, yPresent
}

func (c *carnivore) scanDistantFood() (xSum, ySum, xPresent, yPresent int) {
	for _, i := range [][2]int{
		{-2, -2}, {-2, -1}, {-2, 0}, {-2, 1}, {-2, 2},
		{-1, -2}, {-1, -1}, {-1, 1}, {-1, 2}, {0, -2},
		{0, 2}, {1, -2}, {1, -1}, {1, 1}, {1, 2},
		{2, -2}, {2, -1}, {2, 0}, {2, 1}, {2, 2},
	} {
		if len(c.g.herbivoresPos[c.y+i[1]][c.x+i[0]]) == 0 {
			continue
		}
		if i[0] != 0 {
			if i[0] < 0 {
				xSum += -1 * len(c.g.herbivoresPos[c.y+i[1]][c.x+i[0]])
			} else {
				xSum += len(c.g.herbivoresPos[c.y+i[1]][c.x+i[0]])
			}
			xPresent = 1
		}
		if i[1] != 0 {
			if i[1] < 0 {
				ySum += -1 * len(c.g.herbivoresPos[c.y+i[1]][c.x+i[0]])
			} else {
				ySum += len(c.g.herbivoresPos[c.y+i[1]][c.x+i[0]])
			}
			yPresent = 1
		}
	}
	return xSum, ySum, xPresent, yPresent
}

func (c *carnivore) chaseDistantSubject(xSum, ySum, xPresent, yPresent int) {
	if xPresent > 0 && yPresent > 0 {
		c.y, c.x = c.chaseXY(xSum, ySum)
	} else if xPresent > 0 {
		c.x = c.chaseX(xSum)
	} else {
		c.y = c.chaseY(ySum)
	}
	c.g.carnivoresPos[c.y][c.x] = append(c.g.carnivoresPos[c.y][c.x], c)
}

func (c *carnivore) chaseXY(xSum, ySum int) (y, x int) {
	if math.Abs(float64(xSum)) == math.Abs(float64(ySum)) {
		if xSum == 0 && ySum == 0 {
			r := rand.Float64()
			if r >= 0.75 {
				return c.y, c.x + 1
			} else if 0.75 > r && r >= 0.5 {
				return c.y, c.x - 1
			} else if 0.5 > r && r >= 0.25 {
				return c.y + 1, c.x
			} else {
				return c.y - 1, c.x
			}
		}
		if xSum > ySum {
			if rand.Float64() >= 0.5 {
				return c.y, c.x + 1
			} else {
				return c.y - 1, c.x
			}
		}
		if ySum > xSum {
			if rand.Float64() >= 0.5 {
				return c.y, c.x - 1
			} else {
				return c.y + 1, c.x
			}
		}
		if ySum == xSum && ySum < 0 {
			if rand.Float64() >= 0.5 {
				return c.y - 1, c.x
			} else {
				return c.y, c.x - 1
			}
		}
		if ySum == xSum && ySum > 0 {
			if rand.Float64() >= 0.5 {
				return c.y + 1, c.x
			} else {
				return c.y, c.x + 1
			}
		}
	} else if math.Abs(float64(xSum)) > math.Abs(float64(ySum)) {
		return c.y, c.x + int(math.Abs(float64(xSum))/float64(xSum))
	} else if math.Abs(float64(xSum)) < math.Abs(float64(ySum)) {
		return c.y + int(math.Abs(float64(ySum))/float64(ySum)), c.x
	}
	return c.y, c.x
}

func (c *carnivore) chaseX(xSum int) int {
	if xSum == 0 {
		if rand.Float64() >= 0.5 {
			return c.x + 1
		} else {
			return c.x - 1
		}
	} else {
		return c.x + int(math.Abs(float64(xSum))/float64(xSum))
	}
}

func (c *carnivore) chaseY(ySum int) int {
	if ySum == 0 {
		if rand.Float64() >= 0.5 {
			return c.y + 1
		} else {
			return c.y - 1
		}
	} else {
		return c.y + int(math.Abs(float64(ySum))/float64(ySum))
	}
}

func (c *carnivore) makeRandomMove() {
	r := rand.Float64()
	if r >= 0.75 {
		c.x = c.x + 1
	} else if 0.75 > r && r >= 0.5 {
		c.x = c.x - 1
	} else if 0.5 > r && r >= 0.25 {
		c.y = c.y + 1
	} else {
		c.y = c.y - 1
	}
	c.g.carnivoresPos[c.y][c.x] = append(c.g.carnivoresPos[c.y][c.x], c)
}
