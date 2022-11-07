package main

func debugNewGame() *Game {
	g := newGame()
	debugSetting(g)
	return g
}
func debugSetting(g *Game) {
	meatExampleP := &Food{}
	meatExampleP.init(
		g,
		"meat",
		nil,
		[2]any{144, 144},
	)

	rottenMeatExampleP := &Food{}
	rottenMeatExampleP.init(
		g,
		"rottenMeat",
		nil,
		[2]any{144, 144},
	)

	vegetableExampleP := &Food{}
	vegetableExampleP.init(
		g,
		"vegetable",
		nil,
		[2]any{144, 144},
	)

}
