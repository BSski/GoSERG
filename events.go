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
	} else if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.reset = true
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if 36 <= x && 76 >= x && 86 <= y && 106 >= y {
			// Start button
			g.pause = false
			buttons["start"].state = 1
		} else if 85 <= x && 124 >= x && 86 <= y && 106 >= y {
			// Pause button
			g.pause = true
			buttons["pause"].state = 1
		} else if 136 <= x && 175 >= x && 86 <= y && 106 >= y {
			// Reset button
			g.reset = true
			buttons["reset"].state = 1
		} else if 826 <= x && 839 >= x && 85 <= y && 98 >= y {
			// Cycles Per Second plus
			g.chosenCyclesPerSec += 1
			if g.chosenCyclesPerSec > 28 {
				g.chosenCyclesPerSec = 28
			}
			g.cyclesPerSec = g.cyclesPerSecList[g.chosenCyclesPerSec]
			buttons["cpsPlus"].state = 1
		} else if 811 <= x && 824 >= x && 85 <= y && 98 >= y {
			// Cycles Per Second minus
			g.chosenCyclesPerSec -= 1
			if g.chosenCyclesPerSec < 0 {
				g.chosenCyclesPerSec = 0
			}
			g.cyclesPerSec = g.cyclesPerSecList[g.chosenCyclesPerSec]
			buttons["cpsMinus"].state = 1
		} else if 826 <= x && 839 >= x && 105 <= y && 118 >= y {
			// g.s.tempo plus
			g.s.tempo += 0.09
			if g.s.tempo > 1.00 {
				g.s.tempo = 1.00
			}
			buttons["tempoPlus"].state = 1
		} else if 811 <= x && 824 >= x && 105 <= y && 118 >= y {
			// g.s.tempo minus
			g.s.tempo -= 0.09
			if g.s.tempo < 0.01 {
				g.s.tempo = 0.01
			}
			buttons["tempoMinus"].state = 1
		} else if 826 <= x && 839 >= x && 125 <= y && 138 >= y {
			// Mutation plus
			g.s.mutationChance += 0.01
			if g.s.mutationChance > 1.00 {
				g.s.mutationChance = 1.00
			}
			buttons["mutationPlus"].state = 1
		} else if 811 <= x && 824 >= x && 125 <= y && 138 >= y {
			// Mutation minus
			g.s.mutationChance -= 0.01
			if g.s.mutationChance < 0.00 {
				g.s.mutationChance = 0.00
			}
			buttons["mutationMinus"].state = 1
		} else if 826 <= x && 839 >= x && 165 <= y && 178 >= y {
			// Herbs starting nr plus
			g.s.herbsStartingNr += 50
			if g.s.herbsStartingNr > 1000 {
				g.s.herbsStartingNr = 1000
			}
			buttons["herbsStartingNrPlus"].state = 1
		} else if 811 <= x && 824 >= x && 165 <= y && 178 >= y {
			// Herbs starting nr minus
			g.s.herbsStartingNr -= 50
			if g.s.herbsStartingNr < 0 {
				g.s.herbsStartingNr = 0
			}
			buttons["herbsStartingNrMinus"].state = 1
		} else if 826 <= x && 839 >= x && 185 <= y && 198 >= y {
			// Herbs energy plus
			g.s.herbsEnergy += 50
			if g.s.herbsEnergy > 5500 {
				g.s.herbsEnergy = 5500
			}
			buttons["herbsEnergyPlus"].state = 1
		} else if 811 <= x && 824 >= x && 185 <= y && 198 >= y {
			// Herbs energy minus
			g.s.herbsEnergy -= 50
			if g.s.herbsEnergy < 0 {
				g.s.herbsEnergy = 0
			}
			buttons["herbsEnergyMinus"].state = 1
		} else if 826 <= x && 839 >= x && 205 <= y && 218 >= y {
			// Herbs per spawn plus
			g.s.herbsPerSpawn += 1
			if g.s.herbsPerSpawn > 100 {
				g.s.herbsPerSpawn = 100
			}
			buttons["herbsPerSpawnPlus"].state = 1
		} else if 811 <= x && 824 >= x && 205 <= y && 218 >= y {
			// Herbs per spawn minus
			g.s.herbsPerSpawn -= 1
			if g.s.herbsPerSpawn < 0 {
				g.s.herbsPerSpawn = 0
			}
			buttons["herbsPerSpawnMinus"].state = 1
		} else if 826 <= x && 839 >= x && 225 <= y && 238 >= y {
			// Herbs spawn rate plus
			g.s.herbsSpawnRate += 1
			if g.s.herbsSpawnRate > 7 {
				g.s.herbsSpawnRate = 7
			}
			buttons["herbsSpawnRatePlus"].state = 1
		} else if 811 <= x && 824 >= x && 225 <= y && 238 >= y {
			// Herbs spawn rate minus
			g.s.herbsSpawnRate -= 1
			if g.s.herbsSpawnRate < 0 {
				g.s.herbsSpawnRate = 0
			}
			buttons["herbsSpawnRateMinus"].state = 1
		} else if 826 <= x && 839 >= x && 265 <= y && 278 >= y {
			// Herbivores starting nr plus
			g.s.herbivoresStartingNr += 20
			if g.s.herbivoresStartingNr > 800 {
				g.s.herbivoresStartingNr = 800
			}
			buttons["herbivoresStartingNrPlus"].state = 1
		} else if 811 <= x && 824 >= x && 265 <= y && 278 >= y {
			// Herbivores starting nr minus
			g.s.herbivoresStartingNr -= 20
			if g.s.herbivoresStartingNr < 0 {
				g.s.herbivoresStartingNr = 0
			}
			buttons["herbivoresStartingNrMinus"].state = 1
		} else if 826 <= x && 839 >= x && 285 <= y && 298 >= y {
			// Herbivores spawn energy plus
			g.s.herbivoresSpawnEnergy += 50
			if g.s.herbivoresSpawnEnergy > 5500 {
				g.s.herbivoresSpawnEnergy = 5500
			}
			buttons["herbivoresSpawnEnergyPlus"].state = 1
		} else if 811 <= x && 824 >= x && 285 <= y && 298 >= y {
			// Herbivores spawn energy minus
			g.s.herbivoresSpawnEnergy -= 50
			if g.s.herbivoresSpawnEnergy < 0 {
				g.s.herbivoresSpawnEnergy = 0
			}
			buttons["herbivoresSpawnEnergyMinus"].state = 1
		} else if 826 <= x && 839 >= x && 305 <= y && 318 >= y {
			// Herbivores breeding level plus
			g.s.herbivoresBreedLevel += 50
			if g.s.herbivoresBreedLevel > 5500 {
				g.s.herbivoresBreedLevel = 5500
			}
			buttons["herbivoresBreedLevelPlus"].state = 1
		} else if 811 <= x && 824 >= x && 305 <= y && 318 >= y {
			// Herbivores breeding level minus
			g.s.herbivoresBreedLevel -= 50
			if g.s.herbivoresBreedLevel < 0 {
				g.s.herbivoresBreedLevel = 0
			}
			buttons["herbivoresBreedLevelMinus"].state = 1
		} else if 826 <= x && 839 >= x && 325 <= y && 338 >= y {
			// Herbivores move cost plus
			g.s.herbivoresMoveCost += 5
			if g.s.herbivoresMoveCost > 200 {
				g.s.herbivoresMoveCost = 200
			}
			buttons["herbivoresMoveCostPlus"].state = 1
		} else if 811 <= x && 824 >= x && 325 <= y && 338 >= y {
			// Herbivores move cost minus
			g.s.herbivoresMoveCost -= 5
			if g.s.herbivoresMoveCost < 0 {
				g.s.herbivoresMoveCost = 0
			}
			buttons["herbivoresMoveCostMinus"].state = 1
		} else if 826 <= x && 839 >= x && 365 <= y && 378 >= y {
			// Carnivores starting nr plus
			g.s.carnivoresStartingNr += 5
			if g.s.carnivoresStartingNr > 300 {
				g.s.carnivoresStartingNr = 300
			}
			buttons["carnivoresStartingNrPlus"].state = 1
		} else if 811 <= x && 824 >= x && 365 <= y && 378 >= y {
			// Carnivores starting nr minus
			g.s.carnivoresStartingNr -= 5
			if g.s.carnivoresStartingNr < 0 {
				g.s.carnivoresStartingNr = 0
			}
			buttons["carnivoresStartingNrMinus"].state = 1
		} else if 826 <= x && 839 >= x && 385 <= y && 398 >= y {
			// Carnivores spawn energy plus
			g.s.carnivoresSpawnEnergy += 50
			if g.s.carnivoresSpawnEnergy > 5500 {
				g.s.carnivoresSpawnEnergy = 5500
			}
			buttons["carnivoresSpawnEnergyPlus"].state = 1
		} else if 811 <= x && 824 >= x && 385 <= y && 398 >= y {
			// Carnivores spawn energy minus
			g.s.carnivoresSpawnEnergy -= 50
			if g.s.carnivoresSpawnEnergy < 0 {
				g.s.carnivoresSpawnEnergy = 0
			}
			buttons["carnivoresSpawnEnergyMinus"].state = 1
		} else if 826 <= x && 839 >= x && 405 <= y && 418 >= y {
			// Carnivores breeding level plus
			g.s.carnivoresBreedLevel += 50
			if g.s.carnivoresBreedLevel > 5500 {
				g.s.carnivoresBreedLevel = 5500
			}
			buttons["carnivoresBreedLevelPlus"].state = 1
		} else if 811 <= x && 824 >= x && 405 <= y && 418 >= y {
			// Carnivores breeding level minus
			g.s.carnivoresBreedLevel -= 50
			if g.s.carnivoresBreedLevel < 0 {
				g.s.carnivoresBreedLevel = 0
			}
			buttons["carnivoresBreedLevelMinus"].state = 1
		} else if 826 <= x && 839 >= x && 425 <= y && 438 >= y {
			// Carnivores move cost plus
			g.s.carnivoresMoveCost += 5
			if g.s.carnivoresMoveCost > 250 {
				g.s.carnivoresMoveCost = 250
			}
			buttons["carnivoresMoveCostPlus"].state = 1
		} else if 811 <= x && 824 >= x && 425 <= y && 438 >= y {
			// Carnivores move cost minus
			g.s.carnivoresMoveCost -= 5
			if g.s.carnivoresMoveCost < 0 {
				g.s.carnivoresMoveCost = 0
			}
			buttons["carnivoresMoveCostMinus"].state = 1
		} else if 860 <= x && 918 >= x && 11 <= y && 44 >= y {
			g.rightPanelOption = 0
		} else if 920 <= x && 978 >= x && 11 <= y && 44 >= y {
			g.rightPanelOption = 1
		} else if 980 <= x && 1036 >= x && 11 <= y && 44 >= y {
			g.rightPanelOption = 2
		}
	}
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for key := range buttons {
			buttons[key].state = 0
		}
	}
}
