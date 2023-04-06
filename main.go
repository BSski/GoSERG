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
// - herbivores seem to get not enough energy from herbs
// maybe add breed refraction time, in action you would check if it is 0 and then breed

// write entire thing til the end and then check all gameClient vars and delete not used

// cos powoduje ze carnivores (moze herbi tez) znikaja z carnivoresPos (CHYBA NIE JEDNAK)

// maybe you can remove tempo at all since you can ;;; no, you can't do TPS to 1 because then it will be too slow for buttons to work
// the simulation is not dying as quickly as in python, something might be wrong, check code

// make sure to mention on GoSERG website that everything was made by you, including charts

// herbivores quantity somehow appears on top right panel chart, right? the green dots
// move speed chart 1-2 pixels up
