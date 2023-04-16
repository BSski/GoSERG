package main

import "github.com/hajimehoshi/ebiten/v2"

var d animalsData

type animalsData struct {
	herbivoresQuantities []int
	carnivoresQuantities []int

	herbivoresMeanSpeeds       []float64
	herbivoresMeanBowelLengths []float64
	herbivoresMeanFatLimits    []float64
	herbivoresMeanLegsLengths  []float64

	carnivoresMeanSpeeds       []float64
	carnivoresMeanBowelLengths []float64
	carnivoresMeanFatLimits    []float64
	carnivoresMeanLegsLengths  []float64

	herbivoresSpeeds       [8]int
	herbivoresBowelLengths [8]int
	herbivoresFatLimits    [8]int
	herbivoresLegsLengths  [8]int

	carnivoresSpeeds       [8]int
	carnivoresBowelLengths [8]int
	carnivoresFatLimits    [8]int
	carnivoresLegsLengths  [8]int
}

func updateTimeCounters(g *game) {
	g.counterPrev = g.counter
	g.counter += g.tempo
	if int(g.counter) >= 120 {
		g.counter = 0
	}

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

func updateTimeTravelStatus(g *game) {
	if g.timeTravelCounter > 0 {
		g.timeTravelCounter -= 1
	} else {
		ebiten.SetTPS(g.cyclesPerSec)
	}
}

func updateAnimalsData(g *game) {
	g.d.herbivoresQuantities = append(g.d.herbivoresQuantities, len(g.herbivores))
	g.d.carnivoresQuantities = append(g.d.carnivoresQuantities, len(g.carnivores))

	if int(g.timeMonth)%4 == 0 && len(g.d.herbivoresQuantities) >= 30000 {
		g.d.herbivoresQuantities = g.d.herbivoresQuantities[len(g.d.herbivoresQuantities)-30000:]
		g.d.carnivoresQuantities = g.d.carnivoresQuantities[len(g.d.carnivoresQuantities)-30000:]
	}

	updateAnimalsMeanData(&g.d.herbivoresMeanSpeeds, len(g.herbivores), &g.d.herbivoresSpeeds)
	updateAnimalsMeanData(&g.d.herbivoresMeanBowelLengths, len(g.herbivores), &g.d.herbivoresBowelLengths)
	updateAnimalsMeanData(&g.d.herbivoresMeanFatLimits, len(g.herbivores), &g.d.herbivoresFatLimits)
	updateAnimalsMeanData(&g.d.herbivoresMeanLegsLengths, len(g.herbivores), &g.d.herbivoresLegsLengths)
	updateAnimalsMeanData(&g.d.carnivoresMeanSpeeds, len(g.carnivores), &g.d.carnivoresSpeeds)
	updateAnimalsMeanData(&g.d.carnivoresMeanBowelLengths, len(g.carnivores), &g.d.carnivoresBowelLengths)
	updateAnimalsMeanData(&g.d.carnivoresMeanFatLimits, len(g.carnivores), &g.d.carnivoresFatLimits)
	updateAnimalsMeanData(&g.d.carnivoresMeanLegsLengths, len(g.carnivores), &g.d.carnivoresLegsLengths)
}

func updateAnimalsMeanData(
	meanValues *[]float64,
	animalsLen int,
	values *[8]int,
) {
	if len(*meanValues) >= 160 {
		*meanValues = (*meanValues)[1:]
	}
	var sum int
	for i := 0; i < 8; i++ {
		sum += (*values)[i] * i
	}
	if animalsLen > 0 {
		*meanValues = append(*meanValues, float64(sum)/float64(animalsLen))
	} else {
		if len(*meanValues) > 0 {
			*meanValues = (*meanValues)[1:]
		}
	}
}
