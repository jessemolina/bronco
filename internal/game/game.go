package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/objects"
)

type Mode int
type Object *objects.Object

const (
	//	screenWidth      = 640
	// screenHeight     = 480
	ModeTitle Mode = iota
	ModeGameStart
	ModeGameOver

	// Objectes, from last to first in layer
	Background = iota
	Floor
	Horse
	Obstacle

	screenWidth  = 960
	screenHeight = 540
)

// ================================================================
// CONVENIENCE FUNCTIONS

func NewGame() *Game {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Go Bronco!")
	g := &Game{}



	g.objects = []objects.Object{
		objects.NewBackground(screenWidth, screenHeight),
		objects.NewTiles(screenWidth, screenHeight),
		objects.NewHorse(screenWidth, screenHeight),
		objects.NewObstacle(screenWidth, screenHeight),
	}
	return g
}

// ================================================================
// CUSTOM TYPE

// Implements the Game interface.
// Methods required are Update, Draw, Layout.
type Game struct {
	tick    uint
	objects []objects.Object
	mode    Mode
}

// Updates the Tick, the time unit for logical updates.
func (g *Game) Update() error {
	g.tick++
	/*
	switch g.mode {
	case ModeTitle:
		g.objects[Horse].Coordinates()
	case ModeGameStart:
		// keep track of horse and obstable x and y's
		// TODO catch key press

		//key := ebiten.KeySpace
		//h.isJump = ebiten.IsKeyPressed(key)

	case ModeGameOver:
	}
	*/

	for _, o := range g.objects {
		o.Update(g.tick)
	}
	return nil
}

// Draws the image on screen every frame.
// Frame depends on the refresh rate of the monitor.
func (g *Game) Draw(screen *ebiten.Image) {
	for _, o := range g.objects {
		err := o.Draw(screen)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}

// Returns the game's logical screen based on the given window
// width and height.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
