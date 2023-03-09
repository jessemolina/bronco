package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/objects"
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGameStart
	ModeGameOver
)

const (
	// Objectes, from last to first in layer
	Background = iota
	Floor
	Obstacle
	Horse
)

const (
	screenWidth  = 960
	screenHeight = 540
	floorOffset  = 8
)

// ================================================================
// CONVENIENCE FUNCTIONS

func NewGame() *Game {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Go, Bronco, Go!")

	g := &Game{mode: ModeGameStart}

	g.objects = []objects.Object{
		objects.NewBackground(screenWidth, screenHeight),
		objects.NewTiles(screenWidth, screenHeight, floorOffset),
		objects.NewObstacle(screenWidth, screenHeight),
		objects.NewHorse(screenWidth, screenHeight),
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

	switch g.mode {
	case ModeTitle:
		for i, o := range g.objects {
			o.Update(g.tick)
			if i == Horse {
				o.Animation("jump")
			} else {
				o.Animation("stop")
			}
		}
		//log.Printf("Coordinates: %v", g.objects[Horse].Coordinates())
	case ModeGameStart:
		for i, o := range g.objects {
			o.Update(g.tick)
			if i == Horse {
				o.Animation("walk")
			} else {
				o.Animation("start")
			}
		}
	case ModeGameOver:
		for i, o := range g.objects {
			o.Update(g.tick)
			if i == Horse {
				o.Animation("walk")
			} else {
				o.Animation("stop")
			}
		}

	default:
		log.Printf("Mode Default: %v", "default")
	}

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
