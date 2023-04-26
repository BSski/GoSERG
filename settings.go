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
	herbsSpawnRate:        2,
	herbsPerSpawn:         6,
	herbsEnergy:           180,
	herbsStartingNr:       180,
	herbivoresStartingNr:  60,
	carnivoresStartingNr:  15,
	herbivoresSpawnEnergy: 240,
	carnivoresSpawnEnergy: 220,
	herbivoresBreedLevel:  280,
	carnivoresBreedLevel:  240,
	herbivoresMoveCost:    8,
	carnivoresMoveCost:    2,
}
