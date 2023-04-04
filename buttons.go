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
			y:       86,
		},
		"pause": {
			state:   0,
			sprites: [2]*ebiten.Image{pauseUpSpr, pauseDownSpr},
			x:       85,
			y:       86,
		},
		"reset": {
			state:   0,
			sprites: [2]*ebiten.Image{resetUpSpr, resetDownSpr},
			x:       136,
			y:       86,
		},
		"cpsPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       85,
		},
		"cpsMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       85,
		},
		"tempoPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       105,
		},
		"tempoMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       105,
		},
		"mutationPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       125,
		},
		"mutationMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       125,
		},
		"herbsStartingNrPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       165,
		},
		"herbsStartingNrMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       165,
		},
		"herbsEnergyPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       185,
		},
		"herbsEnergyMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       185,
		},
		"herbsPerSpawnPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       205,
		},
		"herbsPerSpawnMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       205,
		},
		"herbsSpawnRatePlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       225,
		},
		"herbsSpawnRateMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       225,
		},
		"herbivoresStartingNrPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       265,
		},
		"herbivoresStartingNrMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       265,
		},
		"herbivoresSpawnEnergyPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       285,
		},
		"herbivoresSpawnEnergyMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       285,
		},
		"herbivoresBreedLevelPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       305,
		},
		"herbivoresBreedLevelMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       305,
		},
		"herbivoresMoveCostPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       325,
		},
		"herbivoresMoveCostMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       325,
		},
		"carnivoresStartingNrPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       365,
		},
		"carnivoresStartingNrMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       365,
		},
		"carnivoresSpawnEnergyPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       385,
		},
		"carnivoresSpawnEnergyMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       385,
		},
		"carnivoresBreedLevelPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       405,
		},
		"carnivoresBreedLevelMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       405,
		},
		"carnivoresMoveCostPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       425,
		},
		"carnivoresMoveCostMinus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       425,
		},
	}
	return buttons
}
