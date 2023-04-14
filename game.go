package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func (g *game) Layout(_, _ int) (int, int) {
	return 1061, 670
}

func (g *game) Update() error {
	processEvents(g)

	if g.reset == true {
		g.reset = false
		g.resetGame()
	}

	g.cyclesPerSec = g.cyclesPerSecList[g.chosenCyclesPerSec]
	ebiten.SetTPS(g.cyclesPerSec)

	if g.pause {
		return nil
	}

	g.counterPrev = g.counter
	g.counter += g.s.tempo
	if int(g.counter) >= 120 {
		g.counter = 0
	}
	g.totalCyclesCounter += 1

	if int(g.counterPrev) != int(g.counter) && int(g.counter)%speeds[g.s.herbsSpawnRate] == 0 {
		spawnHerbs(g, g.s.herbsPerSpawn)
	}

	for i := 0; i < len(g.herbs); i++ {
		g.herbs[i].age += 1
	}

	for i := 0; i < len(g.carnivores); i++ {
		if g.carnivores[i].energy <= 0 {
			g.carnivores[i].starve()
		}
	}
	for i := 0; i < len(g.herbivores); i++ {
		if g.herbivores[i].energy <= 0 {
			g.herbivores[i].starve()
		}
	}

	for i := 0; i < len(g.carnivores); i++ {
		g.carnivores[i].action()
		g.carnivores[i].age += 1
	}
	for i := 0; i < len(g.herbivores); i++ {
		g.herbivores[i].action()
		g.herbivores[i].age += 1
	}

	for i := 0; i < len(g.carnivores); i++ {
		g.carnivores[i].move()
	}

	for i := 0; i < len(g.herbivores); i++ {
		g.herbivores[i].move()
	}

	if int(g.counterPrev) != int(g.counter) {
		//if len(g.d.herbivoresQuantities) >= 160 {
		//	g.d.herbivoresQuantities = (g.d.herbivoresQuantities)[1:]
		//}
		//g.d.herbivoresQuantities = append(g.d.herbivoresQuantities, len(g.herbivores))
		//g.d.herbivoresTotalQuantities = append(g.d.herbivoresTotalQuantities, len(g.herbivores))
		//
		//if len(g.d.carnivoresQuantities) >= 160 {
		//	g.d.carnivoresQuantities = (g.d.carnivoresQuantities)[1:]
		//}
		//g.d.carnivoresQuantities = append(g.d.carnivoresQuantities, len(g.carnivores))
		//g.d.carnivoresTotalQuantities = append(g.d.carnivoresTotalQuantities, len(g.carnivores))

		g.updateAnimalsMeanData(&g.d.herbivoresMeanSpeeds, len(g.herbivores), &g.d.herbivoresSpeeds)
		g.updateAnimalsMeanData(&g.d.herbivoresMeanBowelLengths, len(g.herbivores), &g.d.herbivoresBowelLengths)
		g.updateAnimalsMeanData(&g.d.herbivoresMeanFatLimits, len(g.herbivores), &g.d.herbivoresFatLimits)
		g.updateAnimalsMeanData(&g.d.herbivoresMeanLegsLengths, len(g.herbivores), &g.d.herbivoresLegsLengths)
		g.updateAnimalsMeanData(&g.d.carnivoresMeanSpeeds, len(g.carnivores), &g.d.carnivoresSpeeds)
		g.updateAnimalsMeanData(&g.d.carnivoresMeanBowelLengths, len(g.carnivores), &g.d.carnivoresBowelLengths)
		g.updateAnimalsMeanData(&g.d.carnivoresMeanFatLimits, len(g.carnivores), &g.d.carnivoresFatLimits)
		g.updateAnimalsMeanData(&g.d.carnivoresMeanLegsLengths, len(g.carnivores), &g.d.carnivoresLegsLengths)
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{Y: 239})

	sc.drawAnimation(screen, g)
	sc.drawLogo(screen)
	sc.drawGrid(screen)
	sc.drawMainUILines(screen)
	sc.drawChartsBg(screen)
	sc.drawQuantitiesChart(screen)
	sc.plotQuantities(screen, g)
	sc.drawHistoricQuantitiesChart(screen)
	sc.plotHistoricQuantities(screen, g)
	sc.drawRightPanel(screen, g)
	sc.plotRightPanel(screen, g)
	sc.drawCounters(screen, g)
	sc.drawSettings(screen, g)
	sc.drawButtons(screen)
	sc.drawHerbs(screen, g)
	sc.drawHerbivores(screen, g)
	sc.drawCarnivores(screen, g)

}
