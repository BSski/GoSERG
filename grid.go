package main

// Generating new grid positions.
func generateGrid() (grid [][][2]float32) {
	squareSize := 10
	y := -squareSize + 10
	for i := 0; i < 43; i++ {
		y += squareSize
		x := -squareSize + 213
		grid = append(grid, [][2]float32{})
		for j := 0; j < 43; j++ {
			x += squareSize
			grid[i] = append(grid[i], [2]float32{float32(x), float32(y)})
		}
	}
	return grid
}

// TODO: check if you must return pointers instead of values here
func generateHerbsPositions() (pos [][][]*herb) {
	for i := 0; i < 43; i++ {
		pos = append(pos, [][]*herb{})
		for j := 0; j < 43; j++ {
			pos[i] = append(pos[i], []*herb{})
		}
	}
	return pos
}

func generateHerbivoresPositions() (pos [][][]*herbivore) {
	for i := 0; i < 43; i++ {
		pos = append(pos, [][]*herbivore{})
		for j := 0; j < 43; j++ {
			pos[i] = append(pos[i], []*herbivore{})
		}
	}
	return pos
}

func generateCarnivoresPositions() (pos [][][]*carnivore) {
	for i := 0; i < 43; i++ {
		pos = append(pos, [][]*carnivore{})
		for j := 0; j < 43; j++ {
			pos[i] = append(pos[i], []*carnivore{})
		}
	}
	return pos
}
