package main

import (
	"github.com/aquilax/go-perlin"
	"image/color"
	"math/rand"
)

func (g *game) clearGame() {
	d := animalsData{}
	g.d = d

	g.counter = 0
	g.counterPrev = 0
	g.timeHour = 0
	g.timeDay = 1
	g.timeMonth = 1
	g.timeYear = 0

	g.timeTravelCounter = 0

	g.herbsPos = generateHerbsPositions()
	g.herbivoresPos = generateHerbivoresPositions()
	g.carnivoresPos = generateCarnivoresPositions()

	g.herbs = []*herb{}
	g.herbivores = []*herbivore{}
	g.carnivores = []*carnivore{}

}

func (g *game) spawnStartingEntities() {
	spawnHerbsOnRandomTiles(g, g.s.herbsStartingNr)
	spawnHerbivore(g, g.s.herbivoresStartingNr)
	spawnCarnivore(g, g.s.carnivoresStartingNr)
}

func generateBoard() (board [][][2]float32) {
	squareSize := 10
	y := -squareSize + 10
	for i := 0; i < 43; i++ {
		y += squareSize
		x := -squareSize + 213
		board = append(board, [][2]float32{})
		for j := 0; j < 43; j++ {
			x += squareSize
			board[i] = append(board[i], [2]float32{float32(x), float32(y)})
		}
	}
	return board
}

func generateHerbsPositions() (pos [][][]*herb) {
	for i := 0; i < 43; i++ {
		pos = append(pos, [][]*herb{})
		for j := 0; j < 43; j++ {
			pos[i] = append(pos[i], []*herb{})
		}
	}
	return pos
}

func generateHerbivoresPositions() (pos [][][]*herbivore) {
	for i := 0; i < 43; i++ {
		pos = append(pos, [][]*herbivore{})
		for j := 0; j < 43; j++ {
			pos[i] = append(pos[i], []*herbivore{})
		}
	}
	return pos
}

func generateCarnivoresPositions() (pos [][][]*carnivore) {
	for i := 0; i < 43; i++ {
		pos = append(pos, [][]*carnivore{})
		for j := 0; j < 43; j++ {
			pos[i] = append(pos[i], []*carnivore{})
		}
	}
	return pos
}

func (g *game) generateNewTerrain() {
	noise := perlin.NewPerlin(2, 2, 3, int64(rand.Intn(100000)))
	a := float64(rand.Intn(18)) + 36
	off1 := float64(rand.Intn(2)) - 4
	off2 := float64(rand.Intn(2)) - 4
	var tilesType [][]tile
	for i := 0; i < 43; i++ {
		tilesType = append(tilesType, []tile{})
		for j := 0; j < 43; j++ {
			tilesType[i] = append(tilesType[i], generateTile(noise, i, j, a, off1, off2))
		}
	}

	g.boardTilesType = tilesType
}

func generateTile(noise *perlin.Perlin, x int, y int, a float64, o1 float64, o2 float64) (t tile) {
	height := noise.Noise2D(float64(x)/a+o1, float64(y)/a+o2)
	switch {
	case height > -0.20:
		// Grass.
		t.color = color.RGBA{
			R: 30,
			G: uint8(125 + int(height/0.11)*9),
			B: 55,
			A: 255,
		}
		t.tileType = 1
	case height <= -0.20 && height > -0.30:
		// Sand.
		t.color = color.RGBA{
			R: uint8(246 + (int(height/0.025)+15)*3 + rand.Intn(24) - 48),
			G: uint8(210 + (int(height/0.025)+15)*3 + rand.Intn(24) - 48),
			B: uint8(165 + rand.Intn(6) - 12),
			A: 255,
		}
		t.tileType = 1
	default:
		// Water.
		t.color = color.RGBA{
			R: 10,
			G: 55,
			B: uint8(240 + int(height/0.15)*25),
			A: 190,
		}
		t.tileType = 0
	}
	return
}
