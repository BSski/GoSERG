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
	arrowDownDownReader := bytes.NewReader(spr.arrowDownDownBytes)
	arrowDownDownSpr, _, err := ebitenutil.NewImageFromReader(arrowDownDownReader)
	if err != nil {
		log.Fatal(err)
	}

	arrowDownUpReader := bytes.NewReader(spr.arrowDownUpBytes)
	arrowDownUpSpr, _, err := ebitenutil.NewImageFromReader(arrowDownUpReader)
	if err != nil {
		log.Fatal(err)
	}

	arrowUpDownReader := bytes.NewReader(spr.arrowUpDownBytes)
	arrowUpDownSpr, _, err := ebitenutil.NewImageFromReader(arrowUpDownReader)
	if err != nil {
		log.Fatal(err)
	}

	arrowUpUpReader := bytes.NewReader(spr.arrowUpUpBytes)
	arrowUpUpSpr, _, err := ebitenutil.NewImageFromReader(arrowUpUpReader)
	if err != nil {
		log.Fatal(err)
	}

	minusDownReader := bytes.NewReader(spr.minusDownBytes)
	minusDownSpr, _, err := ebitenutil.NewImageFromReader(minusDownReader)
	if err != nil {
		log.Fatal(err)
	}

	minusUpReader := bytes.NewReader(spr.minusUpBytes)
	minusUpSpr, _, err := ebitenutil.NewImageFromReader(minusUpReader)
	if err != nil {
		log.Fatal(err)
	}

	plusDownReader := bytes.NewReader(spr.plusDownBytes)
	plusDownSpr, _, err := ebitenutil.NewImageFromReader(plusDownReader)
	if err != nil {
		log.Fatal(err)
	}

	plusUpReader := bytes.NewReader(spr.plusUpBytes)
	plusUpSpr, _, err := ebitenutil.NewImageFromReader(plusUpReader)
	if err != nil {
		log.Fatal(err)
	}

	startDownReader := bytes.NewReader(spr.startDownBytes)
	startDownSpr, _, err := ebitenutil.NewImageFromReader(startDownReader)
	if err != nil {
		log.Fatal(err)
	}

	startUpReader := bytes.NewReader(spr.startUpBytes)
	startUpSpr, _, err := ebitenutil.NewImageFromReader(startUpReader)
	if err != nil {
		log.Fatal(err)
	}

	pauseDownReader := bytes.NewReader(spr.pauseDownBytes)
	pauseDownSpr, _, err := ebitenutil.NewImageFromReader(pauseDownReader)
	if err != nil {
		log.Fatal(err)
	}

	pauseUpReader := bytes.NewReader(spr.pauseUpBytes)
	pauseUpSpr, _, err := ebitenutil.NewImageFromReader(pauseUpReader)
	if err != nil {
		log.Fatal(err)
	}

	resetDownReader := bytes.NewReader(spr.resetDownBytes)
	resetDownSpr, _, err := ebitenutil.NewImageFromReader(resetDownReader)
	if err != nil {
		log.Fatal(err)
	}

	resetUpReader := bytes.NewReader(spr.resetUpBytes)
	resetUpSpr, _, err := ebitenutil.NewImageFromReader(resetUpReader)
	if err != nil {
		log.Fatal(err)
	}

	slowModeUpReader := bytes.NewReader(spr.slowModeUpBytes)
	slowModeUpSpr, _, err := ebitenutil.NewImageFromReader(slowModeUpReader)
	if err != nil {
		log.Fatal(err)
	}

	slowModeDownReader := bytes.NewReader(spr.slowModeDownBytes)
	slowModeDownSpr, _, err := ebitenutil.NewImageFromReader(slowModeDownReader)
	if err != nil {
		log.Fatal(err)
	}

	timeTravelUpReader := bytes.NewReader(spr.timeTravelUpBytes)
	timeTravelUpSpr, _, err := ebitenutil.NewImageFromReader(timeTravelUpReader)
	if err != nil {
		log.Fatal(err)
	}
	timeTravelDownReader := bytes.NewReader(spr.timeTravelDownBytes)
	timeTravelDownSpr, _, err := ebitenutil.NewImageFromReader(timeTravelDownReader)
	if err != nil {
		log.Fatal(err)
	}

	cleanUpReader := bytes.NewReader(spr.cleanUpBytes)
	cleanUpSpr, _, err := ebitenutil.NewImageFromReader(cleanUpReader)
	if err != nil {
		log.Fatal(err)
	}
	cleanDownReader := bytes.NewReader(spr.cleanDownBytes)
	cleanDownSpr, _, err := ebitenutil.NewImageFromReader(cleanDownReader)
	if err != nil {
		log.Fatal(err)
	}

	var buttons = map[string]*button{
		"start": {
			state:   0,
			sprites: [2]*ebiten.Image{startUpSpr, startDownSpr},
			x:       36,
			y:       71,
		},
		"pause": {
			state:   0,
			sprites: [2]*ebiten.Image{pauseUpSpr, pauseDownSpr},
			x:       85,
			y:       71,
		},
		"reset": {
			state:   0,
			sprites: [2]*ebiten.Image{resetUpSpr, resetDownSpr},
			x:       85,
			y:       120,
		},
		"slowMode": {
			state:   0,
			sprites: [2]*ebiten.Image{slowModeUpSpr, slowModeDownSpr},
			x:       36,
			y:       120,
		},
		"timeTravel": {
			state:   0,
			sprites: [2]*ebiten.Image{timeTravelUpSpr, timeTravelDownSpr},
			x:       136,
			y:       71,
		},
		"clean": {
			state:   0,
			sprites: [2]*ebiten.Image{cleanUpSpr, cleanDownSpr},
			x:       136,
			y:       120,
		},
		"chosenAchievementUp": {
			state:   0,
			sprites: [2]*ebiten.Image{arrowUpUpSpr, arrowUpDownSpr},
			x:       164,
			y:       185,
		},
		"chosenAchievementDown": {
			state:   0,
			sprites: [2]*ebiten.Image{arrowDownUpSpr, arrowDownDownSpr},
			x:       179,
			y:       185,
		},
		"gameSpeedPlus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       94,
		},
		"gameSpeedMinus": {
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
