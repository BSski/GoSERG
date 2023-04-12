package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type game struct {
	s                    settings
	c                    consts
	d                    data
	grid                 [][][2]float32
	boardSize            int
	regularTilesQuantity int

	animation        []rune
	animationCounter int

	counter            float64
	counterPrev        float64
	totalCyclesCounter int

	reset bool
	pause bool

	plotHistoricQuantitiesCheckbox bool

	chosenCyclesPerSec int
	cyclesPerSec       int
	cyclesPerSecList   [29]int

	herbs      []*herb
	herbivores []*herbivore
	carnivores []*carnivore

	herbsPos      [][][]*herb
	herbivoresPos [][][]*herbivore
	carnivoresPos [][][]*carnivore

	rightPanelSprites [3]*ebiten.Image
	rightPanelOption  int
}

func (g *game) init() {
	g.cyclesPerSec = g.cyclesPerSecList[g.chosenCyclesPerSec]
	g.regularTilesQuantity = (g.boardSize - 2) * (g.boardSize - 2)
}

func newGame() *game {
	rightPanelOption0, _, err := ebitenutil.NewImageFromFile("sprites/right_panel_buttons0.png")
	if err != nil {
		log.Fatal(err)
	}
	rightPanelOption1, _, err := ebitenutil.NewImageFromFile("sprites/right_panel_buttons1.png")
	if err != nil {
		log.Fatal(err)
	}
	rightPanelOption2, _, err := ebitenutil.NewImageFromFile("sprites/right_panel_buttons2.png")
	if err != nil {
		log.Fatal(err)
	}

	g := &game{
		s:         s,
		c:         c,
		d:         d,
		grid:      generateGrid(),
		boardSize: 41,

		animation:        []rune("||||////----\\\\\\\\"),
		animationCounter: 0,

		counter:            0,
		counterPrev:        0,
		totalCyclesCounter: 0,

		reset: true,
		pause: false,

		plotHistoricQuantitiesCheckbox: true,

		chosenCyclesPerSec: 3,
		cyclesPerSecList: [29]int{
			30,
			60,
			90,
			120,
			150,
			180,
			240,
			300,
			360,
			450,
			600,
			720,
			900,
			1200,
			1800,
			2400,
			3000,
			3600,
			4200,
			4800,
			5400,
			6000,
			8000,
			10000,
			12000,
			15000,
			18000,
			21000,
			25000,
		},
		cyclesPerSec: 0,

		herbs:      []*herb{},
		herbivores: []*herbivore{},
		carnivores: []*carnivore{},

		herbsPos:      generateHerbsPositions(),
		herbivoresPos: generateHerbivoresPositions(),
		carnivoresPos: generateCarnivoresPositions(),

		rightPanelSprites: [3]*ebiten.Image{rightPanelOption0, rightPanelOption1, rightPanelOption2},
		rightPanelOption:  0,
	}
	g.init()
	return g
}

func (g *game) resetGame() {
	g.counter = 0
	g.counterPrev = 0
	g.totalCyclesCounter = 0

	g.herbsPos = generateHerbsPositions()
	g.herbivoresPos = generateHerbivoresPositions()
	g.carnivoresPos = generateCarnivoresPositions()

	g.herbs = []*herb{}
	g.herbivores = []*herbivore{}
	g.carnivores = []*carnivore{}

	d := data{}
	g.d = d

	spawnHerbs(g, g.s.herbsStartingNr)
	spawnHerbivore(g, g.s.herbivoresStartingNr)
	spawnCarnivore(g, g.s.carnivoresStartingNr)
}
