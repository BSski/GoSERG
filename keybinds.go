package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func checkKeybinds(g *Game) {
	g.paused = pauseAction(g)
}

func pauseAction(g *Game) bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.paused == false {
			return true
		}
		return false
	}
	return g.paused
}
