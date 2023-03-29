package main

type settings struct {
	tempo float32

	mutationChance float32

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
