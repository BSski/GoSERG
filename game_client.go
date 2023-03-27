package main

type game struct {
}

func newGame() *game {
	g := &game{}
	reset(g)
	return g
}

func reset(g *game) {
}
