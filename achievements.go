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
	"allAchievements": {
		false,
		"All achievements",
		achievementAllAchievements,
		"Complete all achievements!",
	},
	"resettedSimulation": {
		false,
		"Reset it",
		achievementResettedSimulation,
		"Reset the simulation.",
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
	"brokenChart": {
		false,
		"Hey! You broke the chart!",
		achievementBrokenChart,
		"Get over 900 carnivores\nor herbivores.",
	},
	"theyCanDrown": {
		false,
		"They can drown?!",
		achievementTheyCanDrown,
		"A herbivore or a carnivore\nhas drowned",
	},
	"massStarvation": {
		false,
		"Mass starvation",
		achievementMassStarvation,
		"All carnivores have speed\nequal to 0.",
	},
	"fromHeroToZeroHerbi": {
		false,
		"From hero to zero",
		achievementFromHeroToZeroHerbi,
		"Get a herbivore with all\ngenes equal to 0.",
	},
	"fromZeroToHeroCarni": {
		false,
		"From zero to hero",
		achievementFromZeroToHeroCarni,
		"Get a carnivore with all\ngenes equal to 7.",
	},
	"theFastestWillPrevail": {
		false,
		"The fastest will prevail",
		achievementTheFastestWillPrevail,
		"Herbivores have average\nspeed equal to 6.9 or more.",
	},
	"armsRace": {
		false,
		"Arms race",
		achievementArmsRace,
		"Carnivores have average\nspeed equal to 6.5 or more.",
	},
}

var achievementNames = []string{
	"allAchievements",
	"resettedSimulation",
	"allDead",
	"allOver300",
	"theyCanDrown",
	"brokenChart",
	"massStarvation",
	"fromHeroToZeroHerbi",
	"fromZeroToHeroCarni",
	"theFastestWillPrevail",
	"armsRace",
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

func achievementResettedSimulation(g *game) {
	for _, event := range g.currentEvents {
		if event == "simulation reset" {
			g.a["resettedSimulation"].completed = true
			return
		}
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

func achievementTheyCanDrown(g *game) {
	for _, event := range g.currentEvents {
		if event == "herbivore drowned" || event == "carnivore drowned" {
			g.a["theyCanDrown"].completed = true
			return
		}
	}
}

func achievementMassStarvation(g *game) {
	if g.d.carnivoresSpeedsCounters[0] == 0 {
		return
	}

	for i := 1; i < len(g.d.carnivoresSpeedsCounters); i++ {
		if g.d.carnivoresSpeedsCounters[i] != 0 {
			return
		}
	}

	g.a["massStarvation"].completed = true
}

func achievementFromHeroToZeroHerbi(g *game) {
	for _, h := range g.herbivores {
		if h.dna[0] == 0 && h.dna[1] == 0 && h.dna[2] == 0 && h.dna[3] == 0 {
			g.a["fromHeroToZeroHerbi"].completed = true
		}
	}
}

func achievementFromZeroToHeroCarni(g *game) {
	for _, c := range g.carnivores {
		if c.dna[0] == 7 && c.dna[1] == 7 && c.dna[2] == 7 && c.dna[3] == 7 {
			g.a["fromZeroToHeroCarni"].completed = true
		}
	}
}

func achievementTheFastestWillPrevail(g *game) {
	if g.d.herbivoresSpeedsCounters[7] == 0 {
		return
	}

	sum := 0
	for i := 0; i < len(g.d.herbivoresSpeedsCounters); i++ {
		sum += g.d.herbivoresSpeedsCounters[i] * i
	}

	if float64(sum)/float64(len(g.herbivores)) >= 6.9 {
		g.a["theFastestWillPrevail"].completed = true
	}
}

func achievementArmsRace(g *game) {
	if g.d.carnivoresSpeedsCounters[7] == 0 {
		return
	}

	sum := 0
	for i := 0; i < len(g.d.carnivoresSpeedsCounters); i++ {
		sum += g.d.carnivoresSpeedsCounters[i] * i
	}

	if float64(sum)/float64(len(g.carnivores)) >= 6.5 {
		g.a["armsRace"].completed = true
	}
}
