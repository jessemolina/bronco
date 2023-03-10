package game

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	Banner
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

	g := &Game{mode: ModeTitle}

	g.objects = []objects.Object{
		objects.NewBackground(screenWidth, screenHeight),
		objects.NewBanner(screenWidth, screenHeight),
		objects.NewTiles(screenWidth, screenHeight, floorOffset),
		objects.NewHorse(screenWidth, screenHeight),
	}

	g.level = NewLevel(screenWidth, screenHeight)
	return g
}

// ================================================================
// CUSTOM TYPE

// Implements the Game interface.
// Methods required are Update, Draw, Layout.
type Game struct {
	tick    uint
	objects []objects.Object
	level   *level
	score   int
	mode    Mode
}

// Updates the Tick, the time unit for logical updates.
func (g *Game) Update() error {
	g.tick++

	switch g.mode {

	case ModeTitle:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.mode = ModeGameStart
			for _, o := range g.objects {
				o.Animation(int(g.mode))
			}
		}

		for _, o := range g.objects {
			o.Update(g.tick)
			/*
				if i == Horse {
					o.Animation("jump")
				} else {
					o.Animation("stop")
				}
			*/
		}

	case ModeGameStart:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.mode = ModeGameOver
			for _, o := range g.objects {
				o.Animation(int(g.mode))
			}

		}
		for _, o := range g.objects {
			o.Update(g.tick)

			/*
				if i == Horse {
					o.Animation("jump")
				} else {
					o.Animation("start")
				}
			*/
		}

		g.level.Update(g.tick)

	case ModeGameOver:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.mode = ModeTitle
			for _, o := range g.objects {
				o.Animation(int(g.mode))
			}

		}

		for _, o := range g.objects {
			o.Update(g.tick)

			/*
				if i == Horse {
					o.Animation("jump")
				} else {
					o.Animation("stop")
				}
			*/
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

	g.level.Draw(screen)

	message := fmt.Sprintf("Mode: %v", g.mode)
	ebitenutil.DebugPrint(screen, message)


}

// Returns the game's logical screen based on the given window
// width and height.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
