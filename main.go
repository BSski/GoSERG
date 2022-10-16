package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	startingHerbivoresNr     = 2
	startingCarnivoresNr     = 2
	startingFoodsNr          = 3
	startingHerbivoresEnergy = 50
	startingCarnivoresEnergy = 50
	startingFoodEnergy       = 20

	tileSize           = 16 // only evens, please
	boardStartX        = 200
	boardStartY        = 18
	boardBorderWidth   = 3
	boardTilesGapWidth = 2
	boardWidthTiles    = 5
	boardWidthPx       = 2*boardBorderWidth + boardWidthTiles*tileSize + (boardWidthTiles-1)*boardTilesGapWidth
	lastTilePx         = (boardWidthTiles-1)*tileSize + (boardWidthTiles-1)*boardTilesGapWidth
	screenWidth        = boardWidthPx + boardStartX + boardStartY
	screenHeight       = boardWidthPx + 2*boardStartY
	tileMiddlePx       = tileSize/2 + tileSize%2

	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	if err := run(os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(stdout io.Writer) error {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GoSERG")

	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
	return nil
}

// TODO: carnivores eating herbivores.
//   TODO: add printing for carnivores.
//   TODO: why printing shows wrong positions?
// TOREMEMBER: everything should be fine if the first thing that an animal does in the Update() is move.

// dzieki temu ze tam jest map, to usuwajac nie musze iterowac po wszystkich, tylko starczy dostep O(1) od delete
// a przy iterowaniu po keysach chyba kazdy bedzie printniety tylko raz

// przebieg rundy:
//   eat
//   move
//   atak

// wszedzie gdzie przekazujesz w argumentach jakies tam dziwne dlugie listy, po prostu bierz wartosc z gameState

// moze potrzebuje interfejsu Herbivore i Carnivore?
