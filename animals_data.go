package main

type data struct {
}

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
