package game

import (
	"bytes"
	"image"
	"log"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/resources/images"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

// ================================================================
// CONVENIENCE FUNCTIONS

func NewGame() *Game {
	ebiten.SetWindowSize(windowWidth, windowWidth)
	ebiten.SetWindowTitle("Go Bronco!")
	g := &Game{}
	return g
}

// ================================================================
// CUSTOM TYPE

// Implements the Game interface.
// Methods required are Update, Draw, Layout.
type Game struct {
	tick int
}

// Updates the Tick, the time unit for logical updates.
func (g *Game) Update() error {
	g.tick++
	return nil
}

// Draws the image on screen every frame.
// Frame depends on the refresh rate of the monitor.
func (g *Game) Draw(screen *ebiten.Image) {
	// TODO Refactor image decoding into internal/resources/images
	img, _, err := image.Decode(bytes.NewReader(images.HorseWalk_png))
	if err != nil {
		log.Fatal(err)
	}

	horse := ebiten.NewImageFromImage(img)

	screen.DrawImage(horse, nil)
}

// Returns the game's logical screen based on the given window
// width and height.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
