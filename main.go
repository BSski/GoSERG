package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	//defer profile.Start().Stop()
	if err := run(os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(stdout io.Writer) error {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(1061, 670)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// TODO: Check what will happen if you turn off vsync.
	_, sergLogo, err := ebitenutil.NewImageFromFile("sprites/serg.png")
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowIcon([]image.Image{sergLogo})
	ebiten.SetWindowTitle("GoSERG")

	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
	return nil
}

// TODO:
// make sure to mention on GoSERG website that everything was made by you, including charts

// dopisz do konca,
// postaw to na WASM na github pages
// opisz ladnie co to jest, co mozna zrobic, co sie klika, nie przesadź
// opisz ze sam zrobiles wszystko
// ze to oryginalnie projekt w pythonie i zrefactorowales go w Pythonie i przepisales na Go
