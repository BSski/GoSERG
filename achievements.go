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

//var achievements = map[string]bool{
//	"allDead":      false,
//	"allOver200":   false,
//	"allOver300":   false,
//	"placeholder1": false,
//	"placeholder2": false,
//	"placeholder3": false,
//	"placeholder4": false,
//	"placeholder5": false,
//	"placeholder6": false,
//	"placeholder7": false,
//	"placeholder8": false,
//	"placeholder9": false,
//}

type achievement struct {
	state     bool
	checkFunc func(g *game)
}

var achievements = map[string]*achievement{
	"allDead":      {false, achievementAllDead},
	"allOver200":   {false, achievementAllOver200},
	"allOver300":   {false, achievementAllOver300},
	"placeholder1": {false, achievementAllDead},
	"placeholder2": {false, achievementAllDead},
	"placeholder3": {false, achievementAllDead},
	"placeholder4": {false, achievementAllDead},
	"placeholder5": {false, achievementAllDead},
	"placeholder6": {false, achievementAllDead},
	"placeholder7": {false, achievementAllDead},
	"placeholder8": {false, achievementAllDead},
}

var achievementsList = []string{
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
}

func achievementAllDead(g *game) {
	if len(g.herbivores) == 0 && len(g.carnivores) == 0 {
		g.a["allDead"].state = true

	}
}

func achievementAllOver200(g *game) {
	if len(g.herbivores) >= 200 && len(g.carnivores) >= 200 {
		g.a["allOver200"].state = true
	}
}

func achievementAllOver300(g *game) {
	if len(g.herbivores) >= 300 && len(g.carnivores) >= 300 {
		g.a["allOver300"].state = true
	}
}
