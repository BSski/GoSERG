package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type game struct {
	s    settings
	c    consts
	grid [][][2]float32

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
	chartsDrawingSpeed int

	herbs      []*herb
	herbivores []*herbivore
	carnivores []*carnivore

	herbsPos      [][][]*herb
	herbivoresPos [][][]*herbivore
	carnivoresPos [][][]*carnivore

	herbivoresQuantities []int
	carnivoresQuantities []int

	herbivoresTotalQuantities []int
	carnivoresTotalQuantities []int

	herbivoresMeanSpeeds       []float64
	herbivoresMeanBowelLengths []float64
	herbivoresMeanFatLimits    []float64
	herbivoresMeanLegsLengths  []float64

	carnivoresMeanSpeeds       []float64
	carnivoresMeanBowelLengths []float64
	carnivoresMeanFatLimits    []float64
	carnivoresMeanLegsLengths  []float64

	herbivoresSpeeds       [8]int
	herbivoresBowelLengths [8]int
	herbivoresFatLimits    [8]int
	herbivoresLegsLengths  [8]int

	carnivoresSpeeds       [8]int
	carnivoresBowelLengths [8]int
	carnivoresFatLimits    [8]int
	carnivoresLegsLengths  [8]int

	rightPanelSprites [3]*ebiten.Image
	rightPanelOption  int
}

func (g *game) init() {
	g.cyclesPerSec = g.cyclesPerSecList[g.chosenCyclesPerSec]
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
		s:    s,
		c:    c,
		grid: generateGrid(),

		animation:        []rune("||||////----\\\\\\\\"),
		animationCounter: 0,

		counter:            0,
		counterPrev:        0,
		totalCyclesCounter: 0,

		reset: true,
		pause: false,

		chosenCyclesPerSec: 19,
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
		chartsDrawingSpeed: 0,
		cyclesPerSec:       0,

		herbs:      []*herb{},
		herbivores: []*herbivore{},
		carnivores: []*carnivore{},

		herbsPos:      generateHerbsPositions(),
		herbivoresPos: generateHerbivoresPositions(),
		carnivoresPos: generateCarnivoresPositions(),

		herbivoresQuantities: []int{},
		carnivoresQuantities: []int{},

		herbivoresTotalQuantities: []int{},
		carnivoresTotalQuantities: []int{},

		herbivoresMeanSpeeds:       []float64{},
		herbivoresMeanBowelLengths: []float64{},
		herbivoresMeanFatLimits:    []float64{},
		herbivoresMeanLegsLengths:  []float64{},

		carnivoresMeanSpeeds:       []float64{},
		carnivoresMeanBowelLengths: []float64{},
		carnivoresMeanFatLimits:    []float64{},
		carnivoresMeanLegsLengths:  []float64{},

		herbivoresSpeeds:       [8]int{0, 0, 0, 0, 0, 0, 0, 0},
		herbivoresBowelLengths: [8]int{0, 0, 0, 0, 0, 0, 0, 0},
		herbivoresFatLimits:    [8]int{0, 0, 0, 0, 0, 0, 0, 0},
		herbivoresLegsLengths:  [8]int{0, 0, 0, 0, 0, 0, 0, 0},

		carnivoresSpeeds:       [8]int{0, 0, 0, 0, 0, 0, 0, 0},
		carnivoresBowelLengths: [8]int{0, 0, 0, 0, 0, 0, 0, 0},
		carnivoresFatLimits:    [8]int{0, 0, 0, 0, 0, 0, 0, 0},
		carnivoresLegsLengths:  [8]int{0, 0, 0, 0, 0, 0, 0, 0},

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

	g.herbivoresQuantities = []int{}
	g.carnivoresQuantities = []int{}

	g.herbivoresTotalQuantities = []int{}
	g.carnivoresTotalQuantities = []int{}

	g.herbivoresMeanSpeeds = []float64{}
	g.herbivoresMeanBowelLengths = []float64{}
	g.herbivoresMeanFatLimits = []float64{}
	g.herbivoresMeanLegsLengths = []float64{}

	g.carnivoresMeanSpeeds = []float64{}
	g.carnivoresMeanBowelLengths = []float64{}
	g.carnivoresMeanFatLimits = []float64{}
	g.carnivoresMeanLegsLengths = []float64{}

	g.herbivoresSpeeds = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	g.herbivoresBowelLengths = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	g.herbivoresFatLimits = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	g.herbivoresLegsLengths = [8]int{0, 0, 0, 0, 0, 0, 0, 0}

	g.carnivoresSpeeds = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	g.carnivoresBowelLengths = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	g.carnivoresFatLimits = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	g.carnivoresLegsLengths = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
}
