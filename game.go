package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *game) Layout(_, _ int) (int, int) {
	return 1061, 670
}

func (g *game) Update() error {
	//checkKeybinds(g)
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {

}
