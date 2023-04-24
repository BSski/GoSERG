package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math/rand"
)

var herb0Spr *ebiten.Image
var herb1Spr *ebiten.Image
var herb2Spr *ebiten.Image

type herb struct {
	g      *game
	x      int
	y      int
	energy int
	age    int
	sprNr  int
	spr    *ebiten.Image
}

func (h *herb) init() {
	switch h.sprNr {
	case 0:
		h.spr = herb0Spr
	case 1:
		h.spr = herb1Spr
	case 2:
		h.spr = herb2Spr
	}
}

func init() {
	var err error
	herb0Reader := bytes.NewReader(spr.herb0Bytes)
	herb0Spr, _, err = ebitenutil.NewImageFromReader(herb0Reader)
	if err != nil {
		log.Fatal(err)
	}
	herb1Reader := bytes.NewReader(spr.herb1Bytes)
	herb1Spr, _, err = ebitenutil.NewImageFromReader(herb1Reader)
	if err != nil {
		log.Fatal(err)
	}
	herb2Reader := bytes.NewReader(spr.herb2Bytes)
	herb2Spr, _, err = ebitenutil.NewImageFromReader(herb2Reader)
	if err != nil {
		log.Fatal(err)
	}
}

func drawSingleHerb(screen *ebiten.Image, x, y float32, sprite *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x)+1, float64(y)+1)
	screen.DrawImage(sprite, options)
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
		sprNr:  rand.Intn(3),
	}
	h.init()
	g.herbs = append(g.herbs, &h)
	g.herbsPos[y][x] = append(g.herbsPos[y][x], &h)
	return true
}

func doHerbsActions(g *game) {
	if int(g.counterPrev) != int(g.counter) && int(g.counter)%speeds[g.s.herbsSpawnRate] == 0 {
		spawnHerbsOnRandomTiles(g, g.s.herbsPerSpawn)
	}

	if int(g.counterPrev) != int(g.counter) {
		for i := 0; i < len(g.herbs); i++ {
			g.herbs[i].age += 1
		}
	}
}
