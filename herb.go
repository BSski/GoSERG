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

func drawHerb(screen *ebiten.Image, x, y float32) {
	vector.DrawFilledCircle(
		screen,
		x+2,
		y+2,
		2,
		color.RGBA{R: 34, G: 139, B: 34, A: 255},
		false,
	)
	vector.DrawFilledRect(
		screen,
		x+6,
		y+1,
		1,
		1,
		color.RGBA{R: 34, G: 139, B: 34, A: 255},
		false,
	)
	vector.DrawFilledRect(
		screen,
		x+3,
		y+7,
		1,
		1,
		color.RGBA{R: 34, G: 139, B: 34, A: 255},
		false,
	)
	vector.DrawFilledCircle(
		screen,
		x+6,
		y+5,
		2,
		color.RGBA{R: 34, G: 139, B: 34, A: 255},
		false,
	)
	vector.DrawFilledCircle(
		screen,
		x+1,
		y+6,
		1,
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
		createHerbOnField(g, x, y)
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
