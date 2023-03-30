package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type button struct {
	state      int
	sprites    [2]*ebiten.Image
	upSprite   string
	downSprite string
	x          float64
	y          float64
}

func getBtns() map[string]button {
	minusDownSpr, _, err := ebitenutil.NewImageFromFile("sprites/minus_down.png")
	if err != nil {
		log.Fatal(err)
	}
	minusUpSpr, _, err := ebitenutil.NewImageFromFile("sprites/minus_up.png")
	if err != nil {
		log.Fatal(err)
	}
	plusDownSpr, _, err := ebitenutil.NewImageFromFile("sprites/plus_down.png")
	if err != nil {
		log.Fatal(err)
	}
	plusUpSpr, _, err := ebitenutil.NewImageFromFile("sprites/plus_up.png")
	if err != nil {
		log.Fatal(err)
	}
	startDownSpr, _, err := ebitenutil.NewImageFromFile("sprites/start_down.png")
	if err != nil {
		log.Fatal(err)
	}
	startUpSpr, _, err := ebitenutil.NewImageFromFile("sprites/start_up.png")
	if err != nil {
		log.Fatal(err)
	}
	pauseDownSpr, _, err := ebitenutil.NewImageFromFile("sprites/pause_down.png")
	if err != nil {
		log.Fatal(err)
	}
	pauseUpSpr, _, err := ebitenutil.NewImageFromFile("sprites/pause_up.png")
	if err != nil {
		log.Fatal(err)
	}
	resetDownSpr, _, err := ebitenutil.NewImageFromFile("sprites/reset_down.png")
	if err != nil {
		log.Fatal(err)
	}
	resetUpSpr, _, err := ebitenutil.NewImageFromFile("sprites/reset_up.png")
	if err != nil {
		log.Fatal(err)
	}

	var buttons = map[string]button{
		"start": {
			state:   0,
			sprites: [2]*ebiten.Image{startUpSpr, startDownSpr},
			x:       36,
			y:       86,
		},
		"pause": {
			state:   0,
			sprites: [2]*ebiten.Image{pauseUpSpr, pauseDownSpr},
			x:       85,
			y:       86,
		},
		"reset": {
			state:   0,
			sprites: [2]*ebiten.Image{resetUpSpr, resetDownSpr},
			x:       136,
			y:       86,
		},
		"cps_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       85,
		},
		"cps_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       85,
		},
		"tempo_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       105,
		},
		"tempo_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       105,
		},
		"mutation_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       125,
		},
		"mutation_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       125,
		},
		"herbs_starting_nr_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       165,
		},
		"herbs_starting_nr_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       165,
		},
		"herbs_energy_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       185,
		},
		"herbs_energy_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       185,
		},
		"herbs_per_spawn_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       205,
		},
		"herbs_per_spawn_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       205,
		},
		"herbs_spawn_rate_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       225,
		},
		"herbs_spawn_rate_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       225,
		},
		"herbivores_starting_nr_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       265,
		},
		"herbivores_starting_nr_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       265,
		},
		"herbivores_spawn_energy_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       285,
		},
		"herbivores_spawn_energy_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       285,
		},
		"herbivores_breed_level_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       305,
		},
		"herbivores_breed_level_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       305,
		},
		"herbivores_move_cost_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       325,
		},
		"herbivores_move_cost_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       325,
		},
		"carnivores_starting_nr_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       365,
		},
		"carnivores_starting_nr_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       365,
		},
		"carnivores_spawn_energy_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       385,
		},
		"carnivores_spawn_energy_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       385,
		},
		"carnivores_breed_level_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       405,
		},
		"carnivores_breed_level_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       405,
		},
		"carnivores_move_cost_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       425,
		},
		"carnivores_move_cost_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       425,
		},
		"carnivores_kill_energy_plus": {
			state:   0,
			sprites: [2]*ebiten.Image{plusUpSpr, plusDownSpr},
			x:       826,
			y:       445,
		},
		"carnivores_kill_energy_minus": {
			state:   0,
			sprites: [2]*ebiten.Image{minusUpSpr, minusDownSpr},
			x:       811,
			y:       445,
		},
	}
	return buttons
}
