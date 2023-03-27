package main

import "math/rand"

func reset(g *game) *game {
	g.meatCntP = new(int)
	g.rottenMeatCntP = new(int)
	g.vegetableCntP = new(int)

	initPos(g)

	g.herbivores = make(map[*herbivore]struct{})
	for i := 0; i < startingHerbivoresCnt; i++ {
		newHerbiP := &herbivore{}
		newHerbiP.init(g, "A herbivore", nil, [2]any{nil, nil})
	}
	g.carnivores = make(map[*carnivore]struct{})
	for i := 0; i < startingCarnivoresCnt; i++ {
		newCarniP := &carnivore{}
		newCarniP.init(g, "A carnivore", nil, [2]any{nil, nil})
	}
	g.meats = make(map[*food]struct{})
	g.rottenMeats = make(map[*food]struct{})
	g.vegetables = make(map[*food]struct{})
	g.foods = make(map[*food]struct{})
	g.bloodSpots = make(map[*bloodSpot]struct{})

	for i := 0; i < startingRandomFoodsCnt; i++ {
		newFoodP := &food{}

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
		newFoodP := &food{}
		newFoodP.init(
			g,
			"meat",
			nil,
			[2]any{nil, nil},
		)
	}
	for i := 0; i < startingRottenMeatCnt; i++ {
		newFoodP := &food{}
		newFoodP.init(
			g,
			"rottenMeat",
			nil,
			[2]any{nil, nil},
		)
	}
	for i := 0; i < startingVegetablesCnt; i++ {
		newFoodP := &food{}
		newFoodP.init(
			g,
			"vegetable",
			nil,
			[2]any{nil, nil},
		)
	}
	return g
}

func initPos(g *game) {
	g.herbivoresPos = make(map[float64]map[float64]map[*herbivore]struct{})
	g.carnivoresPos = make(map[float64]map[float64]map[*carnivore]struct{})
	g.meatPos = make(map[float64]map[float64]map[*food]struct{})
	g.rottenMeatPos = make(map[float64]map[float64]map[*food]struct{})
	g.vegetablesPos = make(map[float64]map[float64]map[*food]struct{})
	g.bloodSpotsPos = make(map[float64]map[float64]map[*bloodSpot]struct{})
	g.tilesPos = make([]float64, 0)

	// Herbivores.
	y := 0
	for i := 0; i < boardWidthTiles; i++ {
		x := 0
		for j := 0; j < boardWidthTiles; j++ {
			if g.herbivoresPos[float64(y)] == nil {
				g.herbivoresPos[float64(y)] = make(map[float64]map[*herbivore]struct{})
			}
			g.herbivoresPos[float64(y)][float64(x)] = make(map[*herbivore]struct{})
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
				g.carnivoresPos[float64(y)] = make(map[float64]map[*carnivore]struct{})
			}
			g.carnivoresPos[float64(y)][float64(x)] = make(map[*carnivore]struct{}, 0)
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
				g.meatPos[float64(y)] = make(map[float64]map[*food]struct{})
			}
			g.meatPos[float64(y)][float64(x)] = make(map[*food]struct{}, 0)
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
				g.rottenMeatPos[float64(y)] = make(map[float64]map[*food]struct{})
			}
			g.rottenMeatPos[float64(y)][float64(x)] = make(map[*food]struct{}, 0)
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
				g.vegetablesPos[float64(y)] = make(map[float64]map[*food]struct{})
			}
			g.vegetablesPos[float64(y)][float64(x)] = make(map[*food]struct{}, 0)
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
				g.bloodSpotsPos[float64(y)] = make(map[float64]map[*bloodSpot]struct{})
			}
			g.bloodSpotsPos[float64(y)][float64(x)] = make(map[*bloodSpot]struct{}, 0)
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
