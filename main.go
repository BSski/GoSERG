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
	debugMode      = 0
	updateInterval = 15

	// Herbivores.
	startingHerbivoresCnt    = 0
	startingHerbivoresEnergy = 30 // Only evens, please.
	herbivoresMoveCost       = 0
	herbivoresMaxEnergy      = 150
	herbivoresBreedThreshold = 70

	// Carnivores.
	startingCarnivoresCnt    = 0
	startingCarnivoresEnergy = 30 // Only evens, please.
	carnivoresMoveCost       = 4
	carnivoresMaxEnergy      = 150
	carnivoresBreedThreshold = 80

	// Food.
	startingRandomFoodsCnt = 0
	startingFoodEnergy     = 50 // Can't be 0.
	startingMeatCnt        = 0
	startingRottenMeatCnt  = 0
	startingVegetablesCnt  = 50

	// Environment.
	foodPerInterval = 5

	// Board settings.
	tileSize           = 14 // only evens, please
	boardStartX        = 200
	boardStartY        = 18
	boardBorderWidth   = 3
	boardTilesGapWidth = 2
	boardWidthTiles    = 10
	boardWidthPx       = 2*boardBorderWidth + boardWidthTiles*tileSize + (boardWidthTiles-1)*boardTilesGapWidth
	lastTilePx         = (boardWidthTiles-1)*tileSize + (boardWidthTiles-1)*boardTilesGapWidth
	tileMiddlePx       = tileSize/2 + tileSize%2
	screenWidth        = boardWidthPx + boardStartX + boardStartY
	screenHeight       = boardWidthPx + 2*boardStartY

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

	if debugMode == 1 {
		if err := ebiten.RunGame(debugNewGame()); err != nil {
			log.Fatal(err)
		}
		return nil
	}
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

// Features:
// chodzenie
// zjadanie miesozerca -> roslinozerca
// teleportacja on the edges
// 3 typy food
// dostawanie energii z jedzenia
// jesli zje w tej turze, to sie w niej nie rusza
// jak umrze to zmienia sie w food (meat)
// jesli poziom energii < 0 => dead
// koszt ruchu
// jeśli food ma więcej energii, niż limit energii zwierzęcia, to zjada ono tylko do syta
// rozmnażanie
// randomowo syp vegetables per round

// TODO:
// add RESET button.
// dodaj plamę krwi w miejscu smierci,
// plama krwi po zjedzeniu kogos na pare rund, byc moze niezalezna od foodu
// nowe ziomki maja żółty border
// entities become old and die (track age)
// sugarcoat the entire thing
// efficient UI
// zielone food się rozrasta na boki losowo i wgl rośnie wraz z turą
// zrob wyswietlanie ze jak jest wiecej niz 1 food na tile, to wyswietlaja sie obok siebie zeby bylo widac
// architecture: moze zrob jeszcze jedna liste, na ktorej trzymalbys wszystkie food per tile, hm /\
// maybe instead of [2]any{x, y} I could just pass math.Vec?
// entities are browseable and one can see their children or even family tree
// counter mechanizm i chodzenie co ileś kroków; szybkość
// geny, cechy
// dodaj animacje umierania/zjadania (krew, czaszka) ktorej animacja trwa niezaleznie od ustawien szybkosci gry i fpsów UI
//   i powyzej jakiejs predkosci po prostu przestaje sie wyswietlac

// when introducing new pair/group behavior inside herbi or carbi
// group (not between groups): make sure to check if animal status is alive

// introduce your own Vector struct and add necessary methods, add, get x y
