package main

func growVegetables(g *Game) {
	for i := 0; i < foodPerInterval; i++ {
		newFoodP := &Food{}
		newFoodP.init(
			g,
			"vegetable",
			nil,
			[2]any{nil, nil},
		)
	}
}
