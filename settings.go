package main

type settings struct {
	mutationChance float64

	herbsSpawnRate int
	herbsPerSpawn  int
	herbsEnergy    int

	herbsStartingNr      int
	herbivoresStartingNr int
	carnivoresStartingNr int

	herbivoresSpawnEnergy int
	carnivoresSpawnEnergy int

	herbivoresBreedLevel int
	carnivoresBreedLevel int

	herbivoresMoveCost int
	carnivoresMoveCost int
}

var s = settings{
	mutationChance:        0.04,
	herbsSpawnRate:        3,
	herbsPerSpawn:         12,
	herbsEnergy:           2000,
	herbsStartingNr:       500,
	herbivoresStartingNr:  200,
	carnivoresStartingNr:  30,
	herbivoresSpawnEnergy: 1900,
	carnivoresSpawnEnergy: 2400,
	herbivoresBreedLevel:  2000,
	carnivoresBreedLevel:  2500,
	herbivoresMoveCost:    70,
	carnivoresMoveCost:    30,
}
