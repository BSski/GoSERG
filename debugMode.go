package main

func debugNewGame() *game {
	g := newGame()
	debugSetting(g)
	return g
}
func debugSetting(g *game) {
	meatExampleP := &food{}
	meatExampleP.init(
		g,
		"meat",
		nil,
		[2]any{144, 144},
	)

	rottenMeatExampleP := &food{}
	rottenMeatExampleP.init(
		g,
		"rottenMeat",
		nil,
		[2]any{144, 144},
	)

	vegetableExampleP := &food{}
	vegetableExampleP.init(
		g,
		"vegetable",
		nil,
		[2]any{144, 144},
	)

}
