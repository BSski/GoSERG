package main

type game struct {
	rightPanelButtonClicked int
	cyclesPerSec            int
	s                       settings

	herbs      []herb
	herbivores []herbivore
	carnivores []carnivore
}

func newGame() *game {
	g := &game{rightPanelButtonClicked: 2}
	reset(g)
	return g
}

func reset(g *game) {
}
