package main

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

	herbivoresSpeedsCounters       [8]int
	herbivoresBowelLengthsCounters [8]int
	herbivoresFatLimitsCounters    [8]int
	herbivoresLegsLengthsCounters  [8]int

	carnivoresSpeedsCounters       [8]int
	carnivoresBowelLengthsCounters [8]int
	carnivoresFatLimitsCounters    [8]int
	carnivoresLegsLengthsCounters  [8]int
}

func updateAnimalsData(g *game) {
	g.d.herbivoresQuantities = append(g.d.herbivoresQuantities, len(g.herbivores))
	g.d.carnivoresQuantities = append(g.d.carnivoresQuantities, len(g.carnivores))

	if int(g.timeMonth)%4 == 0 && len(g.d.herbivoresQuantities) >= 30000 {
		g.d.herbivoresQuantities = g.d.herbivoresQuantities[len(g.d.herbivoresQuantities)-30000:]
		g.d.carnivoresQuantities = g.d.carnivoresQuantities[len(g.d.carnivoresQuantities)-30000:]
	}

	updateAnimalsMeanData(&g.d.herbivoresMeanSpeeds, len(g.herbivores), &g.d.herbivoresSpeedsCounters)
	updateAnimalsMeanData(&g.d.herbivoresMeanBowelLengths, len(g.herbivores), &g.d.herbivoresBowelLengthsCounters)
	updateAnimalsMeanData(&g.d.herbivoresMeanFatLimits, len(g.herbivores), &g.d.herbivoresFatLimitsCounters)
	updateAnimalsMeanData(&g.d.herbivoresMeanLegsLengths, len(g.herbivores), &g.d.herbivoresLegsLengthsCounters)
	updateAnimalsMeanData(&g.d.carnivoresMeanSpeeds, len(g.carnivores), &g.d.carnivoresSpeedsCounters)
	updateAnimalsMeanData(&g.d.carnivoresMeanBowelLengths, len(g.carnivores), &g.d.carnivoresBowelLengthsCounters)
	updateAnimalsMeanData(&g.d.carnivoresMeanFatLimits, len(g.carnivores), &g.d.carnivoresFatLimitsCounters)
	updateAnimalsMeanData(&g.d.carnivoresMeanLegsLengths, len(g.carnivores), &g.d.carnivoresLegsLengthsCounters)
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
