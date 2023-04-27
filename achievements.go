package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var checkAchievement func(g *game)
var achievementBarSpr *ebiten.Image
var achievementBarChosenSpr *ebiten.Image
var starUnlockedSpr *ebiten.Image
var starLockedSpr *ebiten.Image

type achievement struct {
	completed   bool
	fullName    string
	checkFunc   func(g *game)
	description string
}

var achievements = map[string]*achievement{
	"allDead": {
		false,
		"The board is empty!",
		achievementAllDead,
		"Get no animals on the board",
	},
	"allOver200": {
		false,
		"Getting a bit crowdy",
		achievementAllOver200,
		"Get over 200 animals of each type",
	},
	"allOver300": {
		false,
		"Even more crowdy",
		achievementAllOver300,
		"Get over 300 animals of each type",
	},
	"placeholder1": {
		false,
		"Placeholder 1",
		achievementAllDead,
		"placeholder1desc",
	},
	"placeholder2": {
		false,
		"Placeholder 2",
		achievementAllDead,
		"placeholder2desc",
	},
	"placeholder3": {
		false,
		"Placeholder 3",
		achievementAllDead,
		"placeholder3desc",
	},
	"placeholder4": {
		false,
		"Placeholder 4",
		achievementAllDead,
		"placeholder4desc",
	},
	"placeholder5": {
		false,
		"Placeholder 5",
		achievementAllDead,
		"placeholder5desc",
	},
	"placeholder6": {
		false,
		"Placeholder 6",
		achievementAllDead,
		"placeholder6desc",
	},
	"placeholder7": {
		false,
		"Placeholder 7",
		achievementAllDead,
		"placeholder7desc",
	},
	"placeholder8": {
		false,
		"Complete all achievements",
		achievementAllDead,
		"Complete all achievements!",
	},
}

var achievementNames = []string{
	"allDead",
	"allOver200",
	"allOver300",
	"placeholder1",
	"placeholder2",
	"placeholder3",
	"placeholder4",
	"placeholder5",
	"placeholder6",
	"placeholder7",
	"placeholder8",
}

func init() {
	var err error
	achievementBarReader := bytes.NewReader(spr.achievementBarBytes)
	achievementBarSpr, _, err = ebitenutil.NewImageFromReader(achievementBarReader)
	if err != nil {
		log.Fatal(err)
	}
	achievementBarChosenReader := bytes.NewReader(spr.achievementBarChosenBytes)
	achievementBarChosenSpr, _, err = ebitenutil.NewImageFromReader(achievementBarChosenReader)
	if err != nil {
		log.Fatal(err)
	}
	starUnlockedReader := bytes.NewReader(spr.starUnlockedBytes)
	starUnlockedSpr, _, err = ebitenutil.NewImageFromReader(starUnlockedReader)
	if err != nil {
		log.Fatal(err)
	}
	starLockedReader := bytes.NewReader(spr.starLockedBytes)
	starLockedSpr, _, err = ebitenutil.NewImageFromReader(starLockedReader)
	if err != nil {
		log.Fatal(err)
	}
}

func achievementAllDead(g *game) {
	if len(g.herbivores) == 0 && len(g.carnivores) == 0 {
		g.a["allDead"].completed = true

	}
}

func achievementAllOver200(g *game) {
	if len(g.herbivores) >= 200 && len(g.carnivores) >= 200 {
		g.a["allOver200"].completed = true
	}
}

func achievementAllOver300(g *game) {
	if len(g.herbivores) >= 300 && len(g.carnivores) >= 300 {
		g.a["allOver300"].completed = true
	}
}
