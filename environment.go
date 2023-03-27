package main

func growVegetables(g *game) {
	for i := 0; i < foodPerInterval; i++ {
		newFoodP := &food{}
		newFoodP.init(
			g,
			"vegetable",
			nil,
			[2]any{nil, nil},
		)
	}
}
