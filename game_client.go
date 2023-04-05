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
	bigCounter         float64
	bigCounterPrev     float64
	counterForFps      float64
	totalCyclesCounter int

	reset bool
	pause bool

	chosenCyclesPerSec   int
	cyclesPerSec         int
	cyclesPerSecList     [14]int
	cyclesPerSecDividers [14]int
	chartsDrawingSpeed   int

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

	herbivoresMeanSpeed        float64
	herbivoresMeanBowelsLength float64
	herbivoresMeanFatLimit     float64
	herbivoresMeanLegsLength   float64

	carnivoresMeanSpeed        float64
	carnivoresMeanBowelsLength float64
	carnivoresMeanFatLimit     float64
	carnivoresMeanLegsLength   float64

	herbivoresMeanSpeeds       []float64
	herbivoresMeanBowelLengths []float64
	herbivoresMeanFatLimits    []float64
	herbivoresMeanLegsLengths  []float64

	carnivoresMeanSpeeds       []float64
	carnivoresMeanBowelLengths []float64
	carnivoresMeanFatLimits    []float64
	carnivoresMeanLegsLengths  []float64

	herbivoresSpeeds       []int
	herbivoresBowelLengths []int
	herbivoresFatLimits    []int
	herbivoresLegsLengths  []int

	carnivoresSpeeds       []int
	carnivoresBowelLengths []int
	carnivoresFatLimits    []int
	carnivoresLegsLengths  []int

	rightPanelSprites [3]*ebiten.Image
	rightPanelOption  int
}

func (g *game) init() {
	g.cyclesPerSec = g.cyclesPerSecList[g.chosenCyclesPerSec]
}

func newGame() *game {
	// This shouldn't be here if you will use this func as reset func.
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

	s := settings{
		boardSize:             41,
		tempo:                 0.28,
		mutationChance:        0.04,
		herbsSpawnRate:        6,
		herbsPerSpawn:         7,
		herbsEnergy:           2000,
		herbsStartingNr:       500,
		herbivoresStartingNr:  200,
		carnivoresStartingNr:  30,
		herbivoresSpawnEnergy: 1900,
		carnivoresSpawnEnergy: 2400,
		herbivoresBreedLevel:  2000,
		carnivoresBreedLevel:  2500,
		herbivoresMoveCost:    70,
		carnivoresMoveCost:    30,
	}
	c := consts{
		partialDnaRange: [4]int{2, 3, 4, 5},
		vonNeumannPerms: [24][4][2]int{
			{{1, 0}, {-1, 0}, {0, 1}, {0, -1}},
			{{1, 0}, {-1, 0}, {0, -1}, {0, 1}},
			{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
			{{1, 0}, {0, 1}, {0, -1}, {-1, 0}},
			{{1, 0}, {0, -1}, {-1, 0}, {0, 1}},
			{{1, 0}, {0, -1}, {0, 1}, {-1, 0}},
			{{-1, 0}, {1, 0}, {0, 1}, {0, -1}},
			{{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
			{{-1, 0}, {0, 1}, {1, 0}, {0, -1}},
			{{-1, 0}, {0, 1}, {0, -1}, {1, 0}},
			{{-1, 0}, {0, -1}, {1, 0}, {0, 1}},
			{{-1, 0}, {0, -1}, {0, 1}, {1, 0}},
			{{0, 1}, {1, 0}, {-1, 0}, {0, -1}},
			{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
			{{0, 1}, {-1, 0}, {1, 0}, {0, -1}},
			{{0, 1}, {-1, 0}, {0, -1}, {1, 0}},
			{{0, 1}, {0, -1}, {1, 0}, {-1, 0}},
			{{0, 1}, {0, -1}, {-1, 0}, {1, 0}},
			{{0, -1}, {1, 0}, {-1, 0}, {0, 1}},
			{{0, -1}, {1, 0}, {0, 1}, {-1, 0}},
			{{0, -1}, {-1, 0}, {1, 0}, {0, 1}},
			{{0, -1}, {-1, 0}, {0, 1}, {1, 0}},
			{{0, -1}, {0, 1}, {1, 0}, {-1, 0}},
			{{0, -1}, {0, 1}, {-1, 0}, {1, 0}},
		},
	}

	g := &game{
		s:    s,
		c:    c,
		grid: generateGrid(),

		animation:        []rune("||||////----\\\\\\\\"),
		animationCounter: 0,

		counter:            0,
		counterPrev:        0,
		bigCounter:         0,
		bigCounterPrev:     0,
		counterForFps:      0,
		totalCyclesCounter: 0,

		reset: true,
		pause: false,

		chosenCyclesPerSec:   13,
		cyclesPerSecList:     [14]int{30, 60, 90, 120, 150, 180, 240, 300, 360, 450, 600, 720, 900, 1200},
		cyclesPerSecDividers: [14]int{1, 2, 3, 4, 5, 6, 8, 10, 10, 12, 20, 20, 20, 20},
		chartsDrawingSpeed:   0,
		cyclesPerSec:         0,

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

		herbivoresMeanSpeed:        0,
		herbivoresMeanBowelsLength: 0,

		herbivoresMeanFatLimit:   0,
		herbivoresMeanLegsLength: 0,

		carnivoresMeanSpeed:        0,
		carnivoresMeanBowelsLength: 0,
		carnivoresMeanFatLimit:     0,
		carnivoresMeanLegsLength:   0,

		herbivoresMeanSpeeds:       []float64{},
		herbivoresMeanBowelLengths: []float64{},
		herbivoresMeanFatLimits:    []float64{},
		herbivoresMeanLegsLengths:  []float64{},

		carnivoresMeanSpeeds:       []float64{},
		carnivoresMeanBowelLengths: []float64{},
		carnivoresMeanFatLimits:    []float64{},
		carnivoresMeanLegsLengths:  []float64{},

		herbivoresSpeeds:       []int{},
		herbivoresBowelLengths: []int{},
		herbivoresFatLimits:    []int{},
		herbivoresLegsLengths:  []int{},

		carnivoresSpeeds:       []int{},
		carnivoresBowelLengths: []int{},
		carnivoresFatLimits:    []int{},
		carnivoresLegsLengths:  []int{},

		rightPanelSprites: [3]*ebiten.Image{rightPanelOption0, rightPanelOption1, rightPanelOption2},
		rightPanelOption:  0,
	}
	g.init()
	reset(g)
	return g
}

func reset(g *game) {
}
