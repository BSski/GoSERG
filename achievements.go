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

//// ACHIEVEMENTY:
//- Looks stable... yet: symulacja miala carnivores i herbivores powyzej 5 przez 2 miesiace~? have above 5 animals of each type for the first 3 months
//  - ten zrob na samym koncu, i wgl pierwsze achievementy niech dotyczą rzeczy, ktore nie wymagaja przyspieszenia
//- They can drown?! - herbivore or carnivore drowned
//- Long ride: max out all settings
//- Small values: min out all settings

var achievements = map[string]*achievement{
	"allAchievements": {
		false,
		"All achievements",
		achievementAllAchievements,
		"Complete all achievements!",
	},
	"allDead": {
		false,
		"The board is empty!",
		achievementAllDead,
		"Get no animals on the board.",
	},
	"allOver300": {
		false,
		"Getting a bit crowdy",
		achievementAllOver300,
		"Get over 300 animals of each\ntype.",
	},
	"massStarvation": {
		false,
		"Mass starvation",
		achievementMassStarvation,
		"All carnivores have speed 0.",
	},
	"brokenChart": {
		false,
		"Hey! You broke the chart!",
		achievementBrokenChart,
		"Get over 900 carnivores\nor herbivores.",
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
}

var achievementNames = []string{
	"allAchievements",
	"allDead",
	"allOver300",
	"massStarvation",
	"brokenChart",
	"placeholder2",
	"placeholder3",
	"placeholder4",
	"placeholder5",
	"placeholder6",
	"placeholder7",
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

func achievementAllOver300(g *game) {
	if len(g.herbivores) >= 300 && len(g.carnivores) >= 300 {
		g.a["allOver300"].completed = true
	}
}

func achievementBrokenChart(g *game) {
	if len(g.herbivores) > 900 || len(g.carnivores) > 900 {
		g.a["brokenChart"].completed = true
	}
}

func achievementMassStarvation(g *game) {
	if len(g.carnivores) == 0 {
		return
	}

	var sumC int
	for i := 0; i < len(g.d.carnivoresSpeeds); i++ {
		sumC += g.d.carnivoresSpeeds[i]
	}

	if int(sumC/len(g.d.carnivoresSpeeds)) == 0 {
		g.a["massStarvation"].completed = true
	}
}

// TODO: check if it works.
func achievementAllAchievements(g *game) {
	for i := range g.a {
		if i == "allAchievements" {
			continue
		}
		if !g.a[i].completed {
			return
		}
	}
	g.a["allAchievements"].completed = true
}
