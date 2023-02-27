package main

import (
	"log"

	"github.com/jessemolina/bronco/internal/game"
)

func main() {
	g := game.NewGame()
	if err := g.Run(); err != nil {
		log.Fatalf("Error in Game: %v", err)
	}
}
