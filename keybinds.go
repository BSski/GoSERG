package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func checkKeybinds(g *game) {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.paused = checkPauseAction(g)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		reset(g)
	}
}

func checkPauseAction(g *game) bool {
	if g.paused == false {
		return true
	} else {
		return false
	}
}
