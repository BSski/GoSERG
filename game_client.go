package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

var (
	pressStart2P      font.Face
	openSansRegular11 font.Face
	openSansRegular12 font.Face
	openSansRegular13 font.Face
	openSansRegular14 font.Face

	buttons map[string]*button
)

type tile struct {
	color    color.RGBA
	tileType uint8
}

type game struct {
	s                    settings
	c                    consts
	d                    animalsData
	board                [][][2]float32
	boardTilesType       [][]tile
	boardSize            int
	regularTilesQuantity int

	animation        []rune
	animationCounter int

	counter           float64
	counterPrev       float64
	timeHour          int
	timeYear          int
	timeMonth         int
	timeDay           int
	timeTravelCounter int

	pause bool

	tempo            float64
	chosenGameSpeed  int
	cyclesPerSecList [5]int
	cyclesPerSec     int

	herbs      []*herb
	herbivores []*herbivore
	carnivores []*carnivore

	herbsPos      [][][]*herb
	herbivoresPos [][][]*herbivore
	carnivoresPos [][][]*carnivore

	rightPanelSprites [3]*ebiten.Image
	rightPanelOption  int
}

func (g *game) init() {
	g.cyclesPerSec = g.cyclesPerSecList[g.chosenGameSpeed-1]
	g.tempo = 0.2 * float64(g.chosenGameSpeed)
	g.regularTilesQuantity = (g.boardSize - 2) * (g.boardSize - 2)

	tt, err := opentype.Parse(PressStart2P_ttf)
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

	tt2, err := opentype.Parse(OpenSansRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	openSansRegular11, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    11,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	openSansRegular12, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	openSansRegular13, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    13,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	openSansRegular14, err = opentype.NewFace(tt2, &opentype.FaceOptions{
		Size:    14,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	buttons = getBtns()
}

func newGame() *game {
	rightPanelOption0Bytes := []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 176, 0, 0, 0, 33, 8, 6, 0, 0, 0, 121, 141, 33, 184, 0, 0, 0, 1, 115, 82, 71, 66, 0, 174, 206, 28, 233, 0, 0, 0, 4, 103, 65, 77, 65, 0, 0, 177, 143, 11, 252, 97, 5, 0, 0, 0, 9, 112, 72, 89, 115, 0, 0, 14, 194, 0, 0, 14, 194, 1, 21, 40, 74, 128, 0, 0, 1, 20, 73, 68, 65, 84, 120, 94, 237, 220, 177, 13, 131, 48, 16, 70, 225, 75, 36, 106, 36, 38, 160, 101, 20, 10, 6, 97, 4, 22, 128, 129, 40, 104, 216, 129, 146, 130, 29, 144, 152, 32, 81, 136, 41, 210, 37, 81, 34, 244, 251, 222, 87, 36, 216, 157, 229, 167, 43, 125, 105, 219, 246, 102, 14, 212, 117, 109, 203, 178, 132, 85, 188, 198, 113, 116, 117, 86, 23, 1, 151, 101, 105, 121, 158, 239, 151, 154, 36, 73, 216, 141, 207, 48, 12, 251, 255, 17, 176, 135, 179, 94, 247, 95, 64, 212, 203, 4, 78, 211, 52, 124, 197, 99, 219, 54, 38, 112, 132, 152, 192, 136, 2, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 227, 35, 77, 211, 216, 60, 207, 97, 245, 244, 238, 222, 63, 16, 240, 143, 157, 121, 153, 30, 157, 18, 112, 223, 247, 182, 174, 107, 88, 1, 223, 99, 2, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 227, 35, 93, 215, 89, 81, 20, 97, 245, 244, 238, 222, 63, 16, 240, 143, 157, 121, 153, 30, 157, 18, 112, 85, 85, 150, 101, 89, 88, 1, 223, 99, 2, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 143, 251, 69, 196, 227, 211, 82, 151, 105, 154, 92, 60, 175, 122, 4, 252, 120, 146, 51, 118, 71, 192, 30, 206, 234, 230, 125, 96, 196, 200, 236, 14, 38, 193, 114, 163, 194, 160, 69, 172, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}
	rightPanelOption0Reader := bytes.NewReader(rightPanelOption0Bytes)
	rightPanelOption0, _, err := ebitenutil.NewImageFromReader(rightPanelOption0Reader)
	if err != nil {
		log.Fatal(err)
	}
	rightPanelOption1Bytes := []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 176, 0, 0, 0, 33, 8, 6, 0, 0, 0, 121, 141, 33, 184, 0, 0, 0, 1, 115, 82, 71, 66, 0, 174, 206, 28, 233, 0, 0, 0, 4, 103, 65, 77, 65, 0, 0, 177, 143, 11, 252, 97, 5, 0, 0, 0, 9, 112, 72, 89, 115, 0, 0, 14, 194, 0, 0, 14, 194, 1, 21, 40, 74, 128, 0, 0, 1, 38, 73, 68, 65, 84, 120, 94, 237, 220, 177, 173, 131, 48, 20, 70, 97, 231, 73, 212, 72, 76, 64, 203, 40, 20, 12, 194, 8, 44, 0, 35, 48, 8, 5, 13, 19, 176, 0, 5, 13, 19, 32, 49, 65, 34, 192, 175, 75, 145, 40, 2, 233, 191, 62, 95, 145, 112, 221, 89, 62, 114, 233, 199, 56, 142, 79, 103, 220, 48, 12, 174, 44, 75, 215, 182, 173, 95, 177, 109, 223, 235, 60, 207, 126, 178, 107, 63, 215, 35, 224, 40, 138, 252, 146, 61, 125, 223, 31, 255, 161, 4, 156, 231, 185, 75, 211, 244, 8, 56, 132, 115, 253, 59, 126, 1, 81, 193, 222, 192, 113, 28, 251, 47, 59, 182, 109, 227, 6, 6, 148, 16, 48, 164, 17, 48, 164, 17, 48, 164, 17, 48, 164, 17, 48, 190, 82, 85, 149, 155, 166, 201, 79, 167, 79, 215, 174, 112, 123, 192, 119, 109, 12, 97, 224, 6, 190, 64, 215, 117, 110, 93, 87, 63, 157, 222, 173, 225, 119, 4, 12, 105, 4, 12, 105, 4, 12, 105, 4, 12, 105, 4, 12, 105, 4, 12, 105, 4, 12, 105, 4, 12, 105, 4, 12, 105, 4, 140, 175, 52, 77, 227, 178, 44, 243, 211, 233, 211, 181, 43, 220, 30, 240, 93, 27, 67, 24, 184, 129, 47, 80, 20, 133, 75, 146, 196, 79, 167, 119, 107, 248, 29, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 1, 67, 26, 143, 251, 25, 19, 218, 211, 82, 71, 192, 251, 51, 149, 214, 237, 1, 47, 203, 226, 39, 219, 254, 3, 14, 225, 92, 31, 117, 93, 155, 127, 31, 24, 86, 57, 247, 2, 248, 156, 114, 163, 223, 1, 61, 222, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}
	rightPanelOption1Reader := bytes.NewReader(rightPanelOption1Bytes)
	rightPanelOption1, _, err := ebitenutil.NewImageFromReader(rightPanelOption1Reader)
	if err != nil {
		log.Fatal(err)
	}
	rightPanelOption2Bytes := []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 176, 0, 0, 0, 33, 8, 6, 0, 0, 0, 121, 141, 33, 184, 0, 0, 0, 1, 115, 82, 71, 66, 0, 174, 206, 28, 233, 0, 0, 0, 4, 103, 65, 77, 65, 0, 0, 177, 143, 11, 252, 97, 5, 0, 0, 0, 9, 112, 72, 89, 115, 0, 0, 14, 194, 0, 0, 14, 194, 1, 21, 40, 74, 128, 0, 0, 1, 1, 73, 68, 65, 84, 120, 94, 237, 220, 177, 9, 131, 64, 24, 134, 225, 51, 96, 237, 92, 142, 226, 8, 46, 224, 10, 14, 98, 97, 227, 4, 46, 96, 97, 227, 4, 130, 19, 36, 104, 126, 187, 8, 41, 66, 224, 251, 124, 159, 70, 239, 186, 31, 94, 174, 56, 208, 108, 28, 199, 103, 50, 55, 12, 67, 170, 170, 42, 205, 243, 28, 59, 190, 206, 89, 219, 182, 141, 29, 111, 71, 192, 121, 158, 199, 210, 79, 223, 247, 199, 243, 12, 248, 46, 179, 222, 33, 224, 178, 44, 211, 35, 222, 1, 73, 156, 192, 70, 174, 78, 224, 162, 40, 226, 205, 199, 182, 109, 156, 192, 208, 71, 192, 144, 70, 192, 144, 70, 192, 144, 70, 192, 144, 70, 192, 144, 246, 247, 128, 235, 186, 78, 211, 52, 197, 202, 207, 167, 249, 220, 102, 238, 186, 46, 173, 235, 26, 171, 183, 111, 247, 126, 141, 19, 24, 210, 8, 24, 210, 8, 24, 210, 8, 24, 210, 8, 24, 210, 8, 24, 210, 8, 24, 210, 8, 24, 210, 8, 24, 210, 8, 24, 210, 8, 24, 210, 248, 164, 200, 8, 159, 20, 1, 98, 8, 24, 210, 8, 24, 210, 8, 24, 210, 8, 24, 210, 8, 24, 210, 8, 24, 210, 184, 7, 54, 114, 117, 15, 236, 106, 191, 7, 62, 2, 222, 127, 201, 233, 238, 12, 248, 46, 179, 46, 203, 18, 43, 111, 89, 211, 52, 246, 255, 7, 134, 171, 148, 94, 137, 176, 126, 243, 138, 164, 163, 177, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}
	rightPanelOption2Reader := bytes.NewReader(rightPanelOption2Bytes)
	rightPanelOption2, _, err := ebitenutil.NewImageFromReader(rightPanelOption2Reader)
	if err != nil {
		log.Fatal(err)
	}

	g := &game{
		s:         s,
		c:         c,
		d:         d,
		board:     generateBoard(),
		boardSize: 41,

		animation:        []rune("||||////----\\\\\\\\"),
		animationCounter: 0,

		pause: false,

		chosenGameSpeed:  2,
		cyclesPerSecList: [5]int{30, 60, 90, 120, 150},

		rightPanelSprites: [3]*ebiten.Image{rightPanelOption0, rightPanelOption1, rightPanelOption2},
		rightPanelOption:  0,
	}

	g.init()
	g.clearGame()
	g.generateNewTerrain()
	g.spawnStartingEntities()
	return g
}
