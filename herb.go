package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math/rand"
)

var herbSpr *ebiten.Image

type herb struct {
	g      *game
	x      int
	y      int
	energy int
	age    int
}

func init() {
	var err error
	herbReader := bytes.NewReader(herbBytes)
	herbSpr, _, err = ebitenutil.NewImageFromReader(herbReader)
	if err != nil {
		log.Fatal(err)
	}
}

func drawSingleHerb(screen *ebiten.Image, x, y float32) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(herbSpr, options)
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

func spawnHerbsOnRandomTiles(g *game, nr int) {
	maxTriesCounter := 200
	for nr > 0 && maxTriesCounter > 0 {
		y := rand.Intn(g.boardSize-2) + 2
		x := rand.Intn(g.boardSize-2) + 2

		created := createHerbOnField(g, x, y)

		if created {
			nr -= 1
		}
		maxTriesCounter -= 1
	}
}

func createHerbOnField(g *game, x, y int) (created bool) {
	if len(g.herbsPos[y][x]) != 0 {
		return false
	}
	if g.boardTilesType[y][x].tileType == 0 {
		return false
	}
	h := herb{
		g:      g,
		x:      x,
		y:      y,
		energy: g.s.herbsEnergy,
	}
	g.herbs = append(g.herbs, &h)
	g.herbsPos[y][x] = append(g.herbsPos[y][x], &h)
	return true
}

func doHerbsActions(g *game) {
	if int(g.counterPrev) != int(g.counter) && int(g.counter)%speeds[g.s.herbsSpawnRate] == 0 {
		spawnHerbsOnRandomTiles(g, g.s.herbsPerSpawn)
	}

	for i := 0; i < len(g.herbs); i++ {
		g.herbs[i].age += 1
	}
}
