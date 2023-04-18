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
	spawnHerbs(g, g.s.herbsStartingNr)
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
	a := float64(rand.Intn(18)) + 18
	off1 := float64(rand.Intn(3)) - 6
	off2 := float64(rand.Intn(3)) - 6
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
	case height > -0.30:
		// Grass.
		t.color = color.RGBA{
			R: 20,
			G: uint8(135 + int(height/0.1)*10),
			B: 45,
			A: 255,
		}
		t.tileType = 1
	case height <= -0.30 && height > -0.40:
		// Sand.
		t.color = color.RGBA{
			R: uint8(240 + (int(height/0.025)+12)*8 + rand.Intn(12) - 24),
			G: uint8(200 + (int(height/0.025)+12)*8 + rand.Intn(12) - 24),
			B: uint8(160 + rand.Intn(8) - 16),
			A: 230,
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
