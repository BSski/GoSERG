package main

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

var d animalsData

func (g *game) updateAnimalsMeanData(
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
