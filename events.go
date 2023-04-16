package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func processEvents(g *game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		spawnHerbivore(g, 50)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		spawnCarnivore(g, 10)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		spawnHerbivore(g, 1)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		spawnCarnivore(g, 1)
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.pause == false {
			g.pause = true
		} else {
			g.pause = false
		}
	} else if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.rightPanelOption = 0
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.rightPanelOption = 1
	} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
		g.rightPanelOption = 2
	} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.clearGame()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.clearGame()
		g.spawnStartingEntities()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		g.timeTravelCounter = 10000
		ebiten.SetTPS(100000)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		switch {
		case 36 <= x && 76 >= x && 96 <= y && 116 >= y:
			g.pause = false
			buttons["start"].state = 1
		case 85 <= x && 124 >= x && 96 <= y && 116 >= y:
			g.pause = true
			buttons["pause"].state = 1
		case 136 <= x && 175 >= x && 96 <= y && 116 >= y:
			g.clearGame()
			g.spawnStartingEntities()
			buttons["reset"].state = 1
		case 826 <= x && 839 >= x && 94 <= y && 107 >= y:
			g.chosenGameSpeed += 1
			if g.chosenGameSpeed > 5 {
				g.chosenGameSpeed = 5
			}
			g.cyclesPerSec = g.cyclesPerSecList[g.chosenGameSpeed-1]
			g.tempo = 0.2 * float64(g.chosenGameSpeed)
			ebiten.SetTPS(g.cyclesPerSec)
			buttons["gameSpeedPlus"].state = 1
		case 811 <= x && 824 >= x && 94 <= y && 107 >= y:
			g.chosenGameSpeed -= 1
			if g.chosenGameSpeed < 1 {
				g.chosenGameSpeed = 1
			}
			g.cyclesPerSec = g.cyclesPerSecList[g.chosenGameSpeed-1]
			g.tempo = 0.2 * float64(g.chosenGameSpeed)
			ebiten.SetTPS(g.cyclesPerSec)
			buttons["gameSpeedMinus"].state = 1
		case 826 <= x && 839 >= x && 114 <= y && 127 >= y:
			g.s.mutationChance += 0.01
			if g.s.mutationChance > 1.00 {
				g.s.mutationChance = 1.00
			}
			buttons["mutationPlus"].state = 1
		case 811 <= x && 824 >= x && 114 <= y && 127 >= y:
			g.s.mutationChance -= 0.01
			if g.s.mutationChance < 0.00 {
				g.s.mutationChance = 0.00
			}
			buttons["mutationMinus"].state = 1
		case 826 <= x && 839 >= x && 154 <= y && 167 >= y:
			g.s.herbsStartingNr += 50
			if g.s.herbsStartingNr > g.regularTilesQuantity {
				g.s.herbsStartingNr = g.regularTilesQuantity
			}
			buttons["herbsStartingNrPlus"].state = 1
		case 811 <= x && 824 >= x && 154 <= y && 167 >= y:
			g.s.herbsStartingNr -= 50
			if g.s.herbsStartingNr < 0 {
				g.s.herbsStartingNr = 0
			}
			buttons["herbsStartingNrMinus"].state = 1
		case 826 <= x && 839 >= x && 174 <= y && 187 >= y:
			g.s.herbsEnergy += 20
			if g.s.herbsEnergy > 980 {
				g.s.herbsEnergy = 980
			}
			buttons["herbsEnergyPlus"].state = 1
		case 811 <= x && 824 >= x && 174 <= y && 187 >= y:
			g.s.herbsEnergy -= 20
			if g.s.herbsEnergy < 0 {
				g.s.herbsEnergy = 0
			}
			buttons["herbsEnergyMinus"].state = 1
		case 826 <= x && 839 >= x && 194 <= y && 207 >= y:
			g.s.herbsPerSpawn += 3
			if g.s.herbsPerSpawn > 300 {
				g.s.herbsPerSpawn = 300
			}
			buttons["herbsPerSpawnPlus"].state = 1
		case 811 <= x && 824 >= x && 194 <= y && 207 >= y:
			g.s.herbsPerSpawn -= 3
			if g.s.herbsPerSpawn < 0 {
				g.s.herbsPerSpawn = 0
			}
			buttons["herbsPerSpawnMinus"].state = 1
		case 826 <= x && 839 >= x && 214 <= y && 227 >= y:
			g.s.herbsSpawnRate += 2
			if g.s.herbsSpawnRate > 8 {
				g.s.herbsSpawnRate = 8
			}
			buttons["herbsSpawnRatePlus"].state = 1
		case 811 <= x && 824 >= x && 214 <= y && 227 >= y:
			g.s.herbsSpawnRate -= 2
			if g.s.herbsSpawnRate < 0 {
				g.s.herbsSpawnRate = 0
			}
			buttons["herbsSpawnRateMinus"].state = 1
		case 826 <= x && 839 >= x && 254 <= y && 267 >= y:
			g.s.herbivoresStartingNr += 20
			if g.s.herbivoresStartingNr > g.regularTilesQuantity {
				g.s.herbivoresStartingNr = g.regularTilesQuantity
			}
			buttons["herbivoresStartingNrPlus"].state = 1
		case 811 <= x && 824 >= x && 254 <= y && 267 >= y:
			g.s.herbivoresStartingNr -= 20
			if g.s.herbivoresStartingNr < 0 {
				g.s.herbivoresStartingNr = 0
			}
			buttons["herbivoresStartingNrMinus"].state = 1
		case 826 <= x && 839 >= x && 274 <= y && 287 >= y:
			g.s.herbivoresSpawnEnergy += 20
			if g.s.herbivoresSpawnEnergy > 980 {
				g.s.herbivoresSpawnEnergy = 980
			}
			buttons["herbivoresSpawnEnergyPlus"].state = 1
		case 811 <= x && 824 >= x && 274 <= y && 287 >= y:
			g.s.herbivoresSpawnEnergy -= 20
			if g.s.herbivoresSpawnEnergy < 0 {
				g.s.herbivoresSpawnEnergy = 0
			}
			buttons["herbivoresSpawnEnergyMinus"].state = 1
		case 826 <= x && 839 >= x && 294 <= y && 307 >= y:
			g.s.herbivoresBreedLevel += 20
			if g.s.herbivoresBreedLevel > 980 {
				g.s.herbivoresBreedLevel = 980
			}
			buttons["herbivoresBreedLevelPlus"].state = 1
		case 811 <= x && 824 >= x && 294 <= y && 307 >= y:
			g.s.herbivoresBreedLevel -= 20
			if g.s.herbivoresBreedLevel < 0 {
				g.s.herbivoresBreedLevel = 0
			}
			buttons["herbivoresBreedLevelMinus"].state = 1
		case 826 <= x && 839 >= x && 314 <= y && 327 >= y:
			g.s.herbivoresMoveCost += 1
			if g.s.herbivoresMoveCost > 50 {
				g.s.herbivoresMoveCost = 50
			}
			buttons["herbivoresMoveCostPlus"].state = 1
		case 811 <= x && 824 >= x && 314 <= y && 327 >= y:
			g.s.herbivoresMoveCost -= 1
			if g.s.herbivoresMoveCost < 0 {
				g.s.herbivoresMoveCost = 0
			}
			buttons["herbivoresMoveCostMinus"].state = 1
		case 826 <= x && 839 >= x && 354 <= y && 367 >= y:
			g.s.carnivoresStartingNr += 5
			if g.s.carnivoresStartingNr > g.regularTilesQuantity {
				g.s.carnivoresStartingNr = g.regularTilesQuantity
			}
			buttons["carnivoresStartingNrPlus"].state = 1
		case 811 <= x && 824 >= x && 354 <= y && 367 >= y:
			g.s.carnivoresStartingNr -= 5
			if g.s.carnivoresStartingNr < 0 {
				g.s.carnivoresStartingNr = 0
			}
			buttons["carnivoresStartingNrMinus"].state = 1
		case 826 <= x && 839 >= x && 374 <= y && 387 >= y:
			g.s.carnivoresSpawnEnergy += 20
			if g.s.carnivoresSpawnEnergy > 980 {
				g.s.carnivoresSpawnEnergy = 980
			}
			buttons["carnivoresSpawnEnergyPlus"].state = 1
		case 811 <= x && 824 >= x && 374 <= y && 387 >= y:
			g.s.carnivoresSpawnEnergy -= 20
			if g.s.carnivoresSpawnEnergy < 0 {
				g.s.carnivoresSpawnEnergy = 0
			}
			buttons["carnivoresSpawnEnergyMinus"].state = 1
		case 826 <= x && 839 >= x && 394 <= y && 407 >= y:
			g.s.carnivoresBreedLevel += 20
			if g.s.carnivoresBreedLevel > 980 {
				g.s.carnivoresBreedLevel = 980
			}
			buttons["carnivoresBreedLevelPlus"].state = 1
		case 811 <= x && 824 >= x && 394 <= y && 407 >= y:
			g.s.carnivoresBreedLevel -= 20
			if g.s.carnivoresBreedLevel < 0 {
				g.s.carnivoresBreedLevel = 0
			}
			buttons["carnivoresBreedLevelMinus"].state = 1
		case 826 <= x && 839 >= x && 414 <= y && 427 >= y:
			g.s.carnivoresMoveCost += 1
			if g.s.carnivoresMoveCost > 50 {
				g.s.carnivoresMoveCost = 50
			}
			buttons["carnivoresMoveCostPlus"].state = 1
		case 811 <= x && 824 >= x && 414 <= y && 427 >= y:
			g.s.carnivoresMoveCost -= 1
			if g.s.carnivoresMoveCost < 0 {
				g.s.carnivoresMoveCost = 0
			}
			buttons["carnivoresMoveCostMinus"].state = 1
		case 860 <= x && 918 >= x && 11 <= y && 44 >= y:
			g.rightPanelOption = 0
		case 920 <= x && 978 >= x && 11 <= y && 44 >= y:
			g.rightPanelOption = 1
		case 980 <= x && 1036 >= x && 11 <= y && 44 >= y:
			g.rightPanelOption = 2
		}
	}
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for key := range buttons {
			buttons[key].state = 0
		}
	}
}
