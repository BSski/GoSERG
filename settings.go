package main

type settings struct {
	boardSize int

	tempo float64

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
