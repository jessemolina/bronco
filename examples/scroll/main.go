package main

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jessemolina/bronco/examples/encode/images"
)

const (
	screenWidth  = 576
	screenHeight = 324

	output = `
screenWidth, screenHeight: %v x %v
imageWidth, imageHeight: %v x %v
x16, y16: %v x %v
offsetX: %v
offsetY: %v
`
)

var (
	bgImage *ebiten.Image
)

func init() {
	// original tile size is 256 x 256
	img, _, err := image.Decode(bytes.NewReader(images.Prairie_png))
	if err != nil {
		log.Fatal(err)
	}
	bgImage = ebiten.NewImageFromImage(img)
}

// viewPort
type viewport struct {
	x16 int
	y16 int
}

func (p *viewport) Move() {
	// 320 x 240
	w, h := bgImage.Size()
	maxX16 := w * 16
	maxY16 := h * 16

	p.x16 += w / 32
	p.y16 += h / 32
	p.x16 %= maxX16
	p.y16 %= maxY16
}

func (p *viewport) Position() (int, int) {
	return p.x16, p.y16
}

type Game struct {
	viewport viewport
}

func (g *Game) Update() error {
	g.viewport.Move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	x16, y16 := g.viewport.Position()

	// control the direction and the rate.

	// move image from left to right
	// move X by x16,
	// offsetX, offsetY := float64(x16/32), float64(y16)/float64(y16)

	// move image from right to left
	offsetX, offsetY := -float64(x16/32), float64(y16/16)/float64(y16/16)

	// move image from up to down
	// offsetX, offsetY := float64(x16)/float64(x16), float64(y16)

	// move image from down to up
	// offsetX, offsetY := float64(x16)/float64(x16), -float64(y16)

	// move image - test
	// increments / pace , y stays the same * 1
	// offsetX, offsetY := -float64(x16/32), float64(y16/16)/float64(y16/16)



	// Draw bgImage on the screen repeatedly.
	const repeat = 4
	screenWidth, screenHeight := screen.Size()
	imageWidth, imageHeight := bgImage.Size()

	// iterates over x and y axis - diagnal scroll
	for j := 0; j < repeat; j++ {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			// shift the image by image w x h
			 op.GeoM.Translate(float64(imageWidth*i), float64(imageHeight*j))
			// ofset sets the direction of the image
			op.GeoM.Translate(offsetX, offsetY)
			// draws the instance of the image
			screen.DrawImage(bgImage, op)
		}
	}

	message := fmt.Sprintf(output,
		screenWidth, screenHeight,
		imageWidth, imageHeight,
		g.viewport.x16, g.viewport.y16,
		offsetX, offsetY,
	)
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
	ebitenutil.DebugPrint(screen, message)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Infinite Scroll (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
