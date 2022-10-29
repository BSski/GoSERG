package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func checkKeybinds(g *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.paused = checkPauseAction(g)
	}
}

func checkPauseAction(g *Game) bool {
	if g.paused == false {
		return true
	} else {
		return false
	}
}
