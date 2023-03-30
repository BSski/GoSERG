package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type game struct {
	cyclesPerSec int
	s            settings

	herbs      []herb
	herbivores []herbivore
	carnivores []carnivore

	herbivoresQuantities []int
	carnivoresQuantities []int

	herbivoresTotalQuantities []int
	carnivoresTotalQuantities []int

	herbivoresMeanSpeed        []float64
	herbivoresMeanBowelsLength []float64
	herbivoresMeanFatLimit     []float64
	herbivoresMeanLegsLength   []float64

	carnivoresMeanSpeed        []float64
	carnivoresMeanBowelsLength []float64
	carnivoresMeanFatLimit     []float64
	carnivoresMeanLegsLength   []float64

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

func newGame() *game {
	rightPanelOption1, _, err := ebitenutil.NewImageFromFile("sprites/right_panel_buttons1.png")
	if err != nil {
		log.Fatal(err)
	}
	rightPanelOption2, _, err := ebitenutil.NewImageFromFile("sprites/right_panel_buttons2.png")
	if err != nil {
		log.Fatal(err)
	}
	rightPanelOption3, _, err := ebitenutil.NewImageFromFile("sprites/right_panel_buttons3.png")
	if err != nil {
		log.Fatal(err)
	}

	g := &game{
		rightPanelSprites: [3]*ebiten.Image{rightPanelOption1, rightPanelOption2, rightPanelOption3},
		rightPanelOption:  2,
	}
	reset(g)
	return g
}

func reset(g *game) {
}
