package main

import "math/rand"

func reset(g *Game) *Game {
	g.meatCntP = new(int)
	g.rottenMeatCntP = new(int)
	g.vegetableCntP = new(int)

	initPos(g)

	g.herbivores = make(map[*Herbivore]struct{})
	for i := 0; i < startingHerbivoresCnt; i++ {
		newHerbiP := &Herbivore{}
		newHerbiP.init(g, "A herbivore", nil, [2]any{nil, nil})
	}
	g.carnivores = make(map[*Carnivore]struct{})
	for i := 0; i < startingCarnivoresCnt; i++ {
		newCarniP := &Carnivore{}
		newCarniP.init(g, "A carnivore", nil, [2]any{nil, nil})
	}
	g.meats = make(map[*Food]struct{})
	g.rottenMeats = make(map[*Food]struct{})
	g.vegetables = make(map[*Food]struct{})
	g.foods = make(map[*Food]struct{})
	g.bloodSpots = make(map[*BloodSpot]struct{})

	for i := 0; i < startingRandomFoodsCnt; i++ {
		newFoodP := &Food{}

		foodTypes := map[int]string{
			0: "meat",
			1: "rottenMeat",
			2: "vegetable",
		}
		newFoodP.init(
			g,
			foodTypes[rand.Intn(len(foodTypes))],
			nil,
			[2]any{nil, nil},
		)
	}
	for i := 0; i < startingMeatCnt; i++ {
		newFoodP := &Food{}
		newFoodP.init(
			g,
			"meat",
			nil,
			[2]any{nil, nil},
		)
	}
	for i := 0; i < startingRottenMeatCnt; i++ {
		newFoodP := &Food{}
		newFoodP.init(
			g,
			"rottenMeat",
			nil,
			[2]any{nil, nil},
		)
	}
	for i := 0; i < startingVegetablesCnt; i++ {
		newFoodP := &Food{}
		newFoodP.init(
			g,
			"vegetable",
			nil,
			[2]any{nil, nil},
		)
	}
	return g
}

func initPos(g *Game) {
	g.herbivoresPos = make(map[float64]map[float64]map[*Herbivore]struct{})
	g.carnivoresPos = make(map[float64]map[float64]map[*Carnivore]struct{})
	g.meatPos = make(map[float64]map[float64]map[*Food]struct{})
	g.rottenMeatPos = make(map[float64]map[float64]map[*Food]struct{})
	g.vegetablesPos = make(map[float64]map[float64]map[*Food]struct{})
	g.bloodSpotsPos = make(map[float64]map[float64]map[*BloodSpot]struct{})
	g.tilesPos = make([]float64, 0)

	// Herbivores.
	y := 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.herbivoresPos[float64(y)] == nil {
				g.herbivoresPos[float64(y)] = make(map[float64]map[*Herbivore]struct{})
			}
			g.herbivoresPos[float64(y)][float64(x)] = make(map[*Herbivore]struct{})
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	// Carnivores.
	y = 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.carnivoresPos[float64(y)] == nil {
				g.carnivoresPos[float64(y)] = make(map[float64]map[*Carnivore]struct{})
			}
			g.carnivoresPos[float64(y)][float64(x)] = make(map[*Carnivore]struct{}, 0)
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	// Meat.
	y = 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.meatPos[float64(y)] == nil {
				g.meatPos[float64(y)] = make(map[float64]map[*Food]struct{})
			}
			g.meatPos[float64(y)][float64(x)] = make(map[*Food]struct{}, 0)
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	// Rotten meat.
	y = 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.rottenMeatPos[float64(y)] == nil {
				g.rottenMeatPos[float64(y)] = make(map[float64]map[*Food]struct{})
			}
			g.rottenMeatPos[float64(y)][float64(x)] = make(map[*Food]struct{}, 0)
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	// Vegetables.
	y = 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.vegetablesPos[float64(y)] == nil {
				g.vegetablesPos[float64(y)] = make(map[float64]map[*Food]struct{})
			}
			g.vegetablesPos[float64(y)][float64(x)] = make(map[*Food]struct{}, 0)
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	// BloodSpots.
	y = 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.bloodSpotsPos[float64(y)] == nil {
				g.bloodSpotsPos[float64(y)] = make(map[float64]map[*BloodSpot]struct{})
			}
			g.bloodSpotsPos[float64(y)][float64(x)] = make(map[*BloodSpot]struct{}, 0)
			x += tileSize + boardTilesGapWidth
		}
		y += tileSize + boardTilesGapWidth
	}

	// FIXME:
	// g.tilesPos is not really needed AFAIK, just used for printing positions. Maybe just create it only if debugging?
	// No need to use append also, since the length is known before. Just do tilesPos[i] = ... .
	y = 0
	for i := 0; i < boardWidthTiles; i++ {
		g.tilesPos = append(g.tilesPos, float64(y))
		y += tileSize + boardTilesGapWidth
	}

}
