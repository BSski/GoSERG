package main

import (
	"GoSERG/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

var (
	pressStart2P   font.Face
	mPlus1pRegular font.Face
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
	mPlus1pRegular, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    36,
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
	//ebitenutil.DrawRect(screen, boardStartX, boardStartY, boardWidthPx, boardWidthPx, color.Gray{Y: 200})
	text.Draw(screen, "SERG", pressStart2P, 50, 80, color.Black)
	text.Draw(screen, "SERG", mPlus1pRegular, 50, 130, color.Black)
}
