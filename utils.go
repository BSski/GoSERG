package main

import "fmt"

func printHerbivores(g *Game) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	for _, z1 := range g.tilesPos {
		fmt.Printf("\nY %v:\n X:", z1)
		for _, z2 := range g.tilesPos {
			fmt.Printf("%v: %v ", z2, g.herbivoresPos[z1][z2])
		}
	}
}
