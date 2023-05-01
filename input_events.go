package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func processEvents(g *game) {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft):
		// FIXME: disable this in release
		spawnHerbivore(g, 30)
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowRight):
		// FIXME: disable this in release
		spawnCarnivore(g, 10)
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowUp):
		g.chosenAchievement -= 1
		if g.chosenAchievement < 0 {
			g.chosenAchievement = 10
		}
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowDown):
		g.chosenAchievement += 1
		if g.chosenAchievement > 10 {
			g.chosenAchievement = 0
		}
	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		g.timeTravelCounter = 0
		if g.pause == false {
			g.pause = true
		} else {
			g.pause = false
		}
	case inpututil.IsKeyJustPressed(ebiten.Key1):
		g.rightPanelOption = 0
	case inpututil.IsKeyJustPressed(ebiten.Key2):
		g.rightPanelOption = 1
	case inpututil.IsKeyJustPressed(ebiten.Key3):
		g.rightPanelOption = 2
	case ebiten.IsKeyPressed(ebiten.KeyC):
		g.clearGame()
	case ebiten.IsKeyPressed(ebiten.KeyA):
		for i := range g.a {
			g.a[i].completed = false
		}

	case inpututil.IsKeyJustPressed(ebiten.KeyR):
		g.addEvent("simulation reset")
		g.clearGame()
		g.generateNewTerrain()
		g.spawnStartingEntities()
	case inpututil.IsKeyJustPressed(ebiten.KeyT):
		g.timeTravelCounter = 960
		g.tempo = 1
		ebiten.SetTPS(1000)

	case ebiten.IsKeyPressed(ebiten.KeyBackspace):
		// FIXME: why doesn't it work?
		buttons["slowMode"].state = 1
		g.timeTravelCounter = 0
		g.tempo = 0.1
		ebiten.SetTPS(30)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if 36 <= x && 74 >= x && 121 <= y && 159 >= y {
			buttons["slowMode"].state = 1
			g.timeTravelCounter = 0
			g.tempo = 0.1
			ebiten.SetTPS(30)
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		switch {
		case 36 <= x && 74 >= x && 71 <= y && 109 >= y:
			buttons["start"].state = 1
			g.pause = false
		case 85 <= x && 123 >= x && 71 <= y && 109 >= y:
			buttons["pause"].state = 1
			g.timeTravelCounter = 0
			g.pause = true
		case 136 <= x && 174 >= x && 71 <= y && 109 >= y:
			buttons["timeTravel"].state = 1
			g.timeTravelCounter = 960
			g.tempo = 1
			ebiten.SetTPS(1000)
		case 85 <= x && 123 >= x && 121 <= y && 159 >= y:
			buttons["reset"].state = 1
			g.addEvent("simulation reset")
			g.clearGame()
			g.generateNewTerrain()
			g.spawnStartingEntities()
		case 164 <= x && 177 >= x && 185 <= y && 198 >= y:
			buttons["chosenAchievementUp"].state = 1
			g.chosenAchievement -= 1
			if g.chosenAchievement < 0 {
				g.chosenAchievement = 0
			}
		case 179 <= x && 192 >= x && 185 <= y && 198 >= y:
			buttons["chosenAchievementDown"].state = 1
			g.chosenAchievement += 1
			if g.chosenAchievement > 10 {
				g.chosenAchievement = 10
			}
		case 826 <= x && 839 >= x && 94 <= y && 107 >= y:
			buttons["gameSpeedPlus"].state = 1
			g.chosenGameSpeed += 1
			if g.chosenGameSpeed > 5 {
				g.chosenGameSpeed = 5
			}
			g.cyclesPerSec = g.cyclesPerSecList[g.chosenGameSpeed-1]
			g.tempo = 0.2 * float64(g.chosenGameSpeed)
			ebiten.SetTPS(g.cyclesPerSec)
		case 811 <= x && 824 >= x && 94 <= y && 107 >= y:
			buttons["gameSpeedMinus"].state = 1
			g.chosenGameSpeed -= 1
			if g.chosenGameSpeed < 1 {
				g.chosenGameSpeed = 1
			}
			g.cyclesPerSec = g.cyclesPerSecList[g.chosenGameSpeed-1]
			g.tempo = 0.2 * float64(g.chosenGameSpeed)
			ebiten.SetTPS(g.cyclesPerSec)
		case 826 <= x && 839 >= x && 114 <= y && 127 >= y:
			buttons["mutationPlus"].state = 1
			g.s.mutationChance += 2
			if g.s.mutationChance > 20 {
				g.s.mutationChance = 20
			}
		case 811 <= x && 824 >= x && 114 <= y && 127 >= y:
			buttons["mutationMinus"].state = 1
			g.s.mutationChance -= 2
			if g.s.mutationChance < 0 {
				g.s.mutationChance = 0
			}
		case 826 <= x && 839 >= x && 154 <= y && 167 >= y:
			buttons["herbsStartingNrPlus"].state = 1
			g.s.herbsStartingNr += 30
			if g.s.herbsStartingNr > 300 {
				g.s.herbsStartingNr = 300
			}
		case 811 <= x && 824 >= x && 154 <= y && 167 >= y:
			buttons["herbsStartingNrMinus"].state = 1
			g.s.herbsStartingNr -= 30
			if g.s.herbsStartingNr < 30 {
				g.s.herbsStartingNr = 30
			}
		case 826 <= x && 839 >= x && 174 <= y && 187 >= y:
			buttons["herbsEnergyPlus"].state = 1
			g.s.herbsEnergy += 30
			if g.s.herbsEnergy > 450 {
				g.s.herbsEnergy = 450
			}
		case 811 <= x && 824 >= x && 174 <= y && 187 >= y:
			buttons["herbsEnergyMinus"].state = 1
			g.s.herbsEnergy -= 30
			if g.s.herbsEnergy < 30 {
				g.s.herbsEnergy = 30
			}
		case 826 <= x && 839 >= x && 194 <= y && 207 >= y:
			buttons["herbsPerSpawnPlus"].state = 1
			g.s.herbsPerSpawn += 3
			if g.s.herbsPerSpawn > 15 {
				g.s.herbsPerSpawn = 15
			}
		case 811 <= x && 824 >= x && 194 <= y && 207 >= y:
			buttons["herbsPerSpawnMinus"].state = 1
			g.s.herbsPerSpawn -= 3
			if g.s.herbsPerSpawn < 3 {
				g.s.herbsPerSpawn = 3
			}
		case 826 <= x && 839 >= x && 214 <= y && 227 >= y:
			// In the game it translates to max rate of 5
			buttons["herbsSpawnRatePlus"].state = 1
			g.s.herbsSpawnRate += 2
			if g.s.herbsSpawnRate > 8 {
				g.s.herbsSpawnRate = 8
			}
		case 811 <= x && 824 >= x && 214 <= y && 227 >= y:
			// In the game it translates to min rate of 1
			buttons["herbsSpawnRateMinus"].state = 1
			g.s.herbsSpawnRate -= 2
			if g.s.herbsSpawnRate < 0 {
				g.s.herbsSpawnRate = 0
			}
		case 826 <= x && 839 >= x && 254 <= y && 267 >= y:
			buttons["herbivoresStartingNrPlus"].state = 1
			g.s.herbivoresStartingNr += 20
			if g.s.herbivoresStartingNr > 300 {
				g.s.herbivoresStartingNr = 300
			}
		case 811 <= x && 824 >= x && 254 <= y && 267 >= y:
			buttons["herbivoresStartingNrMinus"].state = 1
			g.s.herbivoresStartingNr -= 20
			if g.s.herbivoresStartingNr < 20 {
				g.s.herbivoresStartingNr = 20
			}
		case 826 <= x && 839 >= x && 274 <= y && 287 >= y:
			buttons["herbivoresSpawnEnergyPlus"].state = 1
			g.s.herbivoresSpawnEnergy += 20
			if g.s.herbivoresSpawnEnergy > 300 {
				g.s.herbivoresSpawnEnergy = 300
			}
		case 811 <= x && 824 >= x && 274 <= y && 287 >= y:
			buttons["herbivoresSpawnEnergyMinus"].state = 1
			g.s.herbivoresSpawnEnergy -= 20
			if g.s.herbivoresSpawnEnergy < 60 {
				g.s.herbivoresSpawnEnergy = 60
			}
		case 826 <= x && 839 >= x && 294 <= y && 307 >= y:
			buttons["herbivoresBreedLevelPlus"].state = 1
			g.s.herbivoresBreedLevel += 20
			if g.s.herbivoresBreedLevel > 300 {
				g.s.herbivoresBreedLevel = 300
			}
		case 811 <= x && 824 >= x && 294 <= y && 307 >= y:
			buttons["herbivoresBreedLevelMinus"].state = 1
			g.s.herbivoresBreedLevel -= 20
			if g.s.herbivoresBreedLevel < 100 {
				g.s.herbivoresBreedLevel = 100
			}
		case 826 <= x && 839 >= x && 314 <= y && 327 >= y:
			buttons["herbivoresMoveCostPlus"].state = 1
			g.s.herbivoresMoveCost += 2
			if g.s.herbivoresMoveCost > 20 {
				g.s.herbivoresMoveCost = 20
			}
		case 811 <= x && 824 >= x && 314 <= y && 327 >= y:
			buttons["herbivoresMoveCostMinus"].state = 1
			g.s.herbivoresMoveCost -= 2
			if g.s.herbivoresMoveCost < 2 {
				g.s.herbivoresMoveCost = 2
			}
		case 826 <= x && 839 >= x && 354 <= y && 367 >= y:
			buttons["carnivoresStartingNrPlus"].state = 1
			g.s.carnivoresStartingNr += 5
			if g.s.carnivoresStartingNr > 75 {
				g.s.carnivoresStartingNr = 75
			}
		case 811 <= x && 824 >= x && 354 <= y && 367 >= y:
			buttons["carnivoresStartingNrMinus"].state = 1
			g.s.carnivoresStartingNr -= 5
			if g.s.carnivoresStartingNr < 5 {
				g.s.carnivoresStartingNr = 5
			}
		case 826 <= x && 839 >= x && 374 <= y && 387 >= y:
			buttons["carnivoresSpawnEnergyPlus"].state = 1
			g.s.carnivoresSpawnEnergy += 20
			if g.s.carnivoresSpawnEnergy > 300 {
				g.s.carnivoresSpawnEnergy = 300
			}
		case 811 <= x && 824 >= x && 374 <= y && 387 >= y:
			buttons["carnivoresSpawnEnergyMinus"].state = 1
			g.s.carnivoresSpawnEnergy -= 20
			if g.s.carnivoresSpawnEnergy < 80 {
				g.s.carnivoresSpawnEnergy = 80
			}
		case 826 <= x && 839 >= x && 394 <= y && 407 >= y:
			buttons["carnivoresBreedLevelPlus"].state = 1
			g.s.carnivoresBreedLevel += 20
			if g.s.carnivoresBreedLevel > 300 {
				g.s.carnivoresBreedLevel = 300
			}
		case 811 <= x && 824 >= x && 394 <= y && 407 >= y:
			buttons["carnivoresBreedLevelMinus"].state = 1
			g.s.carnivoresBreedLevel -= 20
			if g.s.carnivoresBreedLevel < 100 {
				g.s.carnivoresBreedLevel = 100
			}
		case 826 <= x && 839 >= x && 414 <= y && 427 >= y:
			buttons["carnivoresMoveCostPlus"].state = 1
			g.s.carnivoresMoveCost += 2
			if g.s.carnivoresMoveCost > 20 {
				g.s.carnivoresMoveCost = 20
			}
		case 811 <= x && 824 >= x && 414 <= y && 427 >= y:
			buttons["carnivoresMoveCostMinus"].state = 1
			g.s.carnivoresMoveCost -= 2
			if g.s.carnivoresMoveCost < 2 {
				g.s.carnivoresMoveCost = 2
			}
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
