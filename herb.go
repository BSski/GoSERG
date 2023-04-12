package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
)

type herb struct {
	g      *game
	x      int
	y      int
	energy int
	age    int
}

func (h *herb) draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		h.g.grid[h.y][h.x][0]+1,
		h.g.grid[h.y][h.x][1]+1,
		5,
		5,
		color.RGBA{R: 34, G: 139, B: 34, A: 255},
		false,
	)
}

func (h *herb) die() {
	for i, v := range h.g.herbs {
		if v == h {
			h.g.herbs = append(h.g.herbs[:i], h.g.herbs[i+1:]...)
			break
		}
	}
	for i, v := range h.g.herbsPos[h.y][h.x] {
		if v == h {
			h.g.herbsPos[h.y][h.x] = append(
				h.g.herbsPos[h.y][h.x][:i],
				h.g.herbsPos[h.y][h.x][i+1:]...,
			)
			break
		}
	}
}

func spawnHerbs(g *game, nr int) {
	for i := 0; i < nr; i++ {
		y := rand.Intn(g.boardSize-2) + 2
		x := rand.Intn(g.boardSize-2) + 2
		if len(g.herbsPos[y][x]) == 0 {
			h := herb{
				g:      g,
				x:      x,
				y:      y,
				energy: g.s.herbivoresSpawnEnergy,
			}
			g.herbs = append(g.herbs, &h)
			g.herbsPos[y][x] = append(g.herbsPos[y][x], &h)
		}
	}
}

func createHerbOnField(g *game, x, y int) {
	if len(g.herbsPos[y][x]) == 0 {
		h := herb{
			g:      g,
			x:      x,
			y:      y,
			energy: g.s.herbsEnergy,
		}
		g.herbs = append(g.herbs, &h)
		g.herbsPos[y][x] = append(g.herbsPos[y][x], &h)
	}
}
