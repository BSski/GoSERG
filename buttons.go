package main

import (
	"bytes"
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
	minusDownReader := bytes.NewReader(minusDownBytes)
	minusDownSpr, _, err := ebitenutil.NewImageFromReader(minusDownReader)
	if err != nil {
		log.Fatal(err)
	}

	minusUpReader := bytes.NewReader(minusUpBytes)
	minusUpSpr, _, err := ebitenutil.NewImageFromReader(minusUpReader)
	if err != nil {
		log.Fatal(err)
	}

	plusDownReader := bytes.NewReader(plusDownBytes)
	plusDownSpr, _, err := ebitenutil.NewImageFromReader(plusDownReader)
	if err != nil {
		log.Fatal(err)
	}

	plusUpReader := bytes.NewReader(plusUpBytes)
	plusUpSpr, _, err := ebitenutil.NewImageFromReader(plusUpReader)
	if err != nil {
		log.Fatal(err)
	}

	startDownReader := bytes.NewReader(startDownBytes)
	startDownSpr, _, err := ebitenutil.NewImageFromReader(startDownReader)
	if err != nil {
		log.Fatal(err)
	}

	startUpReader := bytes.NewReader(startUpBytes)
	startUpSpr, _, err := ebitenutil.NewImageFromReader(startUpReader)
	if err != nil {
		log.Fatal(err)
	}

	pauseDownReader := bytes.NewReader(pauseDownBytes)
	pauseDownSpr, _, err := ebitenutil.NewImageFromReader(pauseDownReader)
	if err != nil {
		log.Fatal(err)
	}

	pauseUpReader := bytes.NewReader(pauseUpBytes)
	pauseUpSpr, _, err := ebitenutil.NewImageFromReader(pauseUpReader)
	if err != nil {
		log.Fatal(err)
	}

	resetDownReader := bytes.NewReader(resetDownBytes)
	resetDownSpr, _, err := ebitenutil.NewImageFromReader(resetDownReader)
	if err != nil {
		log.Fatal(err)
	}

	resetUpReader := bytes.NewReader(resetUpBytes)
	resetUpSpr, _, err := ebitenutil.NewImageFromReader(resetUpReader)
	if err != nil {
		log.Fatal(err)
	}

	checkboxOffReader := bytes.NewReader(checkboxOffBytes)
	checkboxOff, _, err := ebitenutil.NewImageFromReader(checkboxOffReader)
	if err != nil {
		log.Fatal(err)
	}

	checkboxOnReader := bytes.NewReader(checkboxOnBytes)
	checkboxOn, _, err := ebitenutil.NewImageFromReader(checkboxOnReader)
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
		"totalHistoryOff": {
			state:   0,
			sprites: [2]*ebiten.Image{checkboxOn, checkboxOff},
			x:       351,
			y:       653,
		},
	}
	return buttons
}
