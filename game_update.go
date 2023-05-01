package main

import "github.com/hajimehoshi/ebiten/v2"

func updateTimeCounters(g *game) {
	g.counterPrev = g.counter
	g.counter += g.tempo
	if int(g.counter) >= 120 {
		g.counter = 0
	}

	if int(g.counterPrev) != int(g.counter) {
		g.timeHour += 1
		if g.timeHour >= 120 {
			g.timeHour = 0
			g.timeDay += 1
		}
		if g.timeDay > 30 {
			g.timeDay = 1
			g.timeMonth += 1
		}
		if g.timeMonth > 12 {
			g.timeMonth = 1
			g.timeYear += 1
		}
	}
}

func updateGameTimeVars(g *game) {
	if buttons["slowMode"].state == 1 {
		return
	}

	if g.timeTravelCounter > 0 {
		g.timeTravelCounter -= 1
		return
	}

	g.tempo = 0.2 * float64(g.chosenGameSpeed)
	ebiten.SetTPS(g.cyclesPerSec)
}
