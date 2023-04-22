package main

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"io"
	"log"
	"os"
)

func main() {
	//defer profile.Start().Stop()
	if err := run(os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(stdout io.Writer) error {
	var imageBytes = []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 16, 0, 0, 0, 16, 8, 6, 0, 0, 0, 31, 243, 255, 97, 0, 0, 0, 4, 103, 65, 77, 65, 0, 0, 177, 143, 11, 252, 97, 5, 0, 0, 0, 9, 112, 72, 89, 115, 0, 0, 14, 196, 0, 0, 14, 196, 1, 149, 43, 14, 27, 0, 0, 0, 105, 73, 68, 65, 84, 56, 79, 99, 124, 253, 250, 245, 127, 6, 10, 0, 216, 128, 255, 26, 198, 80, 46, 4, 48, 222, 56, 203, 112, 252, 248, 113, 40, 15, 2, 44, 45, 45, 25, 28, 55, 27, 64, 121, 16, 176, 223, 247, 2, 3, 19, 148, 77, 54, 24, 120, 3, 168, 19, 136, 216, 2, 140, 88, 49, 188, 94, 168, 122, 147, 9, 198, 248, 0, 109, 3, 177, 77, 100, 58, 24, 227, 3, 212, 9, 68, 108, 41, 12, 91, 128, 97, 75, 177, 3, 159, 144, 40, 54, 128, 194, 64, 100, 96, 0, 0, 217, 1, 59, 90, 31, 64, 229, 54, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}
	imageReader := bytes.NewReader(imageBytes)
	_, sergLogoZ, err := ebitenutil.NewImageFromReader(imageReader)
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(1061, 670)
	ebiten.SetWindowIcon([]image.Image{sergLogoZ})
	ebiten.SetWindowTitle("GoSERG")

	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
	return nil
}
