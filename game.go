package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func (g *game) Layout(_, _ int) (int, int) {
	return 1061, 670
}

func (g *game) Update() error {
	processEvents(g)

	updateGameTimeVars(g)

	if g.pause {
		return nil
	}

	updateTimeCounters(g)

	doHerbsActions(g)
	doCarnivoreActions(g)
	doHerbivoreActions(g)

	updateAnimalsData(g)

	chosenAchievementName := achievementNames[g.chosenAchievement]
	checkAchievement = achievements[chosenAchievementName].checkFunc

	if !achievements[chosenAchievementName].completed {
		checkAchievement(g)
	}

	g.cleanCurrentEvents()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{Y: 239})

	sc.drawAnimation(screen, g)
	sc.drawLogo(screen)
	sc.drawBoard(screen, g)
	sc.drawMainUILines(screen)
	sc.drawChartsBg(screen)
	sc.drawHistoricQuantitiesChart(screen)
	sc.plotHistoricQuantities(screen, g)
	sc.drawRightPanel(screen, g)
	sc.plotRightPanel(screen, g)
	sc.drawCounters(screen, g)
	sc.drawSettings(screen, g)
	sc.drawButtons(screen)
	sc.drawHerbs(screen, g)
	sc.drawHerbivores(screen, g)
	sc.drawCarnivores(screen, g)
	sc.drawAchievements(screen, g)
}
