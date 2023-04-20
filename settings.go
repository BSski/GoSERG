package main

type settings struct {
	mutationChance int

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
	mutationChance:        4,
	herbsSpawnRate:        3,
	herbsPerSpawn:         12,
	herbsEnergy:           200,
	herbsStartingNr:       500,
	herbivoresStartingNr:  200,
	carnivoresStartingNr:  30,
	herbivoresSpawnEnergy: 200,
	carnivoresSpawnEnergy: 220,
	herbivoresBreedLevel:  220,
	carnivoresBreedLevel:  240,
	herbivoresMoveCost:    7,
	carnivoresMoveCost:    3,
}
