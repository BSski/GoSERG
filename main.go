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
	debugMode = 0

	// Herbivores.
	startingHerbivoresCnt    = 5
	startingHerbivoresEnergy = 50
	herbivoresMoveCost       = 2
	herbivoresMaxEnergy      = 150

	// Carnivores.
	startingCarnivoresCnt    = 5
	startingCarnivoresEnergy = 50
	carnivoresMoveCost       = 2
	carnivoresMaxEnergy      = 150

	// Food.
	startingFoodsCnt                = 10
	startingFoodEnergy              = 20
	startingAdditionalMeatCnt       = 0
	startingAdditionalRottenMeatCnt = 0
	startingAdditionalVegetablesCnt = 0

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

// po cos chciales dodac sety, ale nie pamietam po co
// https://gist.github.com/bgadrian/cb8b9344d9c66571ef331a14eb7a2e80
// chyba zeby zrobic operacje Has

// dodaj plamę krwi w miejscu zjedzenia  -- a moze ogolnie w miejscu smierci? czyli po prostu meat mialby taka grafike
// dodaj aureolę naokoło świeżo urodzonych entities

// Features:
// chodzenie
// zjadanie miesozerca -> roslinozerca
// teleportacja on the edges
// 3 typy food
//

// TODO:
// dostawanie energii, gdy sie zje
// atakowanie przez carnivores zescheduluj
// zmienianie sie w warzywo jak sie umrze
// sprawdzanie czy poziom energii jest mniejszy niz 0, jesli tak to ded
// koszt ruchu
// plama krwi po zjedzeniu kogos na pare rund, byc moze niezalezna od foodu
// nowe ziomki maja żółty border
// sugarcoat the entire thing
// efficient UI
// zielone food się rozrasta na boki losowo i wgl rośnie wraz z turą
// jeśli food ma więcej energii, niż limit energii zwierzęcia, to zjada ono tylko do syta

// porownaj sobie funkcje carnivore i herbi czy sa takie same,
// czy moze gdzies zapomnialem zmienic w jednym strukcie

// FIXME: gdy na jednym tile jest vegetable i meat, to mrugają, bo kolejnosc iterowania po mapie
// g.foods jest zmienna. dlatego raz jeden, raz drugi są rysowane jako ostatni, przez co mruga
// albo uzyj OrderedMap, albo rozdziel g.foods na 3 listy
