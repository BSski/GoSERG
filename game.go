package main

import (
	"GoSERG/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

var (
	pressStart2P     font.Face
	mPlus1pRegular11 font.Face
	mPlus1pRegular18 font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	pressStart2P, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	tt2, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mPlus1pRegular11, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    11,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mPlus1pRegular18, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    18,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *game) Layout(_, _ int) (int, int) {
	return 1061, 670
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{Y: 239})
	text.Draw(screen, "SERG", pressStart2P, 40, 60, color.Gray{Y: 80})
	text.Draw(screen, "bsski 2023", mPlus1pRegular11, 82, 75, color.Gray{Y: 200})

	squareSize := 10
	y := 19
	x := 222

	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(42*squareSize-1), float64(42*squareSize-1), color.Gray{Y: 210})

	for i := 1; i < 42; i++ {
		ebitenutil.DrawLine(screen, float64(x), float64(y+i*squareSize), float64(x+42*squareSize-1), float64(y+i*squareSize), color.Gray{Y: 239})
	}
	for j := 1; j < 42; j++ {
		ebitenutil.DrawLine(screen, float64(x+j*squareSize), float64(y), float64(x+j*squareSize), float64(y+42*squareSize-1), color.Gray{Y: 239})
	}

	ebitenutil.DrawRect(screen, 29, 124, 162, 300, color.White)

	ebitenutil.DrawRect(screen, 860, 44, 176, 6, color.Gray{Y: 210})
	ebitenutil.DrawRect(screen, 850, 50, 196, 609, color.Gray{Y: 210})
	ebitenutil.DrawRect(screen, 37, 459, 801, 191, color.White)

	//ebitenutil.DrawLine(screen, float64(x+j*squareSize), float64(y), float64(x+j*squareSize), float64(y+42*squareSize-1), color.Gray{Y: 210})

}
