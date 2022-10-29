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
	startingHerbivoresNr     = 10
	startingCarnivoresNr     = 10
	startingFoodsNr          = 10
	startingHerbivoresEnergy = 50
	startingCarnivoresEnergy = 50
	startingFoodEnergy       = 20

	tileSize           = 14 // only evens, please
	boardStartX        = 200
	boardStartY        = 18
	boardBorderWidth   = 3
	boardTilesGapWidth = 2
	boardWidthTiles    = 10
	boardWidthPx       = 2*boardBorderWidth + boardWidthTiles*tileSize + (boardWidthTiles-1)*boardTilesGapWidth
	lastTilePx         = (boardWidthTiles-1)*tileSize + (boardWidthTiles-1)*boardTilesGapWidth
	screenWidth        = boardWidthPx + boardStartX + boardStartY
	screenHeight       = boardWidthPx + 2*boardStartY
	tileMiddlePx       = tileSize/2 + tileSize%2

	exitFail = 1 // exitFail is the exit code if the program fails.
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

// TOREMEMBER: everything should be fine if the first thing that an animal does in the Update() is move.
// to chyba bylo przez to, ze w pierwszym ruchu ziomki nie sa dodawane do carnivoresPos?

// przebieg rundy:
//   eat
//   move
//   atak

// po cos chciales dodac sety, ale nie pamietam po co
// https://gist.github.com/bgadrian/cb8b9344d9c66571ef331a14eb7a2e80
// chyba zeby zrobic operacje Has

// dodaj plamę krwi w miejscu zjedzenia  -- a moze ogolnie w miejscu smierci? czyli po prostu meat mialby taka grafike
// dodaj aureolę naokoło świeżo urodzonych entities
