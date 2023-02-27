package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jessemolina/bronco/examples/encode/images"
)

// ================================================================
// DEFAULTS

/*
const (
	// Size of the Window
	screenWidth  = 320
	screenHeight = 240

	// Frame Starting Postion
	frameOX = 0
	frameOY = 128

	// Single Frame Size & Count
	frameWidth  = 128
	frameHeight = 128
	frameCount  = 9
)
*/

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8
)

var (
	walkerImage *ebiten.Image
)

var message = `g.count: %v
x0, y0: (%v,%v)
x1, y1: (%v,%v)
i: %v
`

// ================================================================
// GAME - Custom Type

// Implements the Game interface.
// Methods required are Update, Draw, Layout.
type Game struct {
	// Starting position for bronco.
	count int
}

// Updates the Tick, the time unit for logical updates.
func (g *Game) Update() error {
	g.count++
	return nil
}

// Draws the image on screen every frame.
// Frame depends on the refresh rate of the monitor.
func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	// -16, -16
	//op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	// 160, 120
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (g.count / 50) % frameCount
	x0, y0:= frameOX+(i*frameWidth), frameOY
	x1, y1 := x0 + frameWidth, y0 + frameHeight
	r := image.Rect(x0, y0, x1, y1)

	sub := walkerImage.SubImage(r).(*ebiten.Image)

	stats := fmt.Sprintf(message,g.count,x0,y0,x1,y1,i)
	ebitenutil.DebugPrint(screen, stats)
	screen.DrawImage(sub, op)
}

// Returns the game's logical screen based on the given window
// width and height.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// ================================================================
// MAIN

func main() {

	// Fill the screen Red
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}

	walkerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Bronco Jump!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
