package main

func debugNewGame() *Game {
	g := newGame()
	debugSetting(g)
	return g
}

// FIXME: why do they appear at 0
func debugSetting(g *Game) {
	meatExampleP := &Food{}
	meatExampleP.init(
		g,
		"meat",
		[]any{144, 144},
		nil,
	)

	rottenMeatExampleP := &Food{}
	rottenMeatExampleP.init(
		g,
		"rottenMeat",
		[]any{144, 144},
		nil,
	)

	vegetableExampleP := &Food{}
	vegetableExampleP.init(
		g,
		"vegetable",
		[]any{144, 144},
		nil,
	)

}
