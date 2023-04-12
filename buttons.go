package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type button struct {
	state      int
	sprites    [2]*ebiten.Image
	upSprite   string
	downSprite string
	x          float64
	y          float64
}

func getBtns() map[string]*button {
	minusDownSpr, _, err := ebitenutil.NewImageFromFile("sprites/minus_down.png")
	if err != nil {
		log.Fatal(err)
	}
	minusUpSpr, _, err := ebitenutil.NewImageFromFile("sprites/minus_up.png")
	if err != nil {
		log.Fatal(err)
	}
	plusDownSpr, _, err := ebitenutil.NewImageFromFile("sprites/plus_down.png")
	if err != nil {
		log.Fatal(err)
	}
	plusUpSpr, _, err := ebitenutil.NewImageFromFile("sprites/plus_up.png")
	if err != nil {
		log.Fatal(err)
	}
	startDownSpr, _, err := ebitenutil.NewImageFromFile("sprites/start_down.png")
	if err != nil {
		log.Fatal(err)
	}
	startUpSpr, _, err := ebitenutil.NewImageFromFile("sprites/start_up.png")
	if err != nil {
		log.Fatal(err)
	}
	pauseDownSpr, _, err := ebitenutil.NewImageFromFile("sprites/pause_down.png")
	if err != nil {
		log.Fatal(err)
	}
	pauseUpSpr, _, err := ebitenutil.NewImageFromFile("sprites/pause_up.png")
	if err != nil {
		log.Fatal(err)
	}
	resetDownSpr, _, err := ebitenutil.NewImageFromFile("sprites/reset_down.png")
	if err != nil {
		log.Fatal(err)
	}
	resetUpSpr, _, err := ebitenutil.NewImageFromFile("sprites/reset_up.png")
	if err != nil {
		log.Fatal(err)
	}

	var buttons = map[string]*button{
		"start": {
			state:   0,
			sprites: [2]*ebiten.Image{startUpSpr, startDownSpr},
			x:       36,
			y:       96,
		},
		"pause": {
			state:   0,
			sprites: [2]*ebiten.Image{pauseUpSpr, pauseDownSpr},
			x:       85,
			y:       96,
		},
		"reset": {
			state:   0,
			sprites: [2]*ebiten.Image{resetUpSpr, resetDownSpr},
			x:       136,
			y:       96,
		},
		"cpsPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       74,
		},
		"cpsMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       74,
		},
		"tempoPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       94,
		},
		"tempoMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       94,
		},
		"mutationPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       114,
		},
		"mutationMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       114,
		},
		"herbsStartingNrPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       154,
		},
		"herbsStartingNrMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       154,
		},
		"herbsEnergyPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       174,
		},
		"herbsEnergyMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       174,
		},
		"herbsPerSpawnPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       194,
		},
		"herbsPerSpawnMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       194,
		},
		"herbsSpawnRatePlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       214,
		},
		"herbsSpawnRateMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       214,
		},
		"herbivoresStartingNrPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       254,
		},
		"herbivoresStartingNrMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       254,
		},
		"herbivoresSpawnEnergyPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       274,
		},
		"herbivoresSpawnEnergyMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       274,
		},
		"herbivoresBreedLevelPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       294,
		},
		"herbivoresBreedLevelMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       294,
		},
		"herbivoresMoveCostPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       314,
		},
		"herbivoresMoveCostMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       314,
		},
		"carnivoresStartingNrPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       354,
		},
		"carnivoresStartingNrMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       354,
		},
		"carnivoresSpawnEnergyPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       374,
		},
		"carnivoresSpawnEnergyMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       374,
		},
		"carnivoresBreedLevelPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       394,
		},
		"carnivoresBreedLevelMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       394,
		},
		"carnivoresMoveCostPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       414,
		},
		"carnivoresMoveCostMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       414,
		},
	}
	return buttons
}
