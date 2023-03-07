package objects

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/assets/images/tiles"
)

func NewTiles(screenWidth int, screenHeight int) Object {
	frameW, frameH := tiles.Prairie02.Image.Size()

	targetY := float64(screenHeight - frameH)
	targetX := float64(0)

	bg := &Tiles{
		img:         tiles.Prairie02.Image,
		frameWidth:  frameW,
		frameHeight: frameH,
		targetX:     targetX,
		targetY:     targetY,
	}

	return bg
}

// Import images that are already decoded.
type Tiles struct {
	img         *ebiten.Image
	frameWidth  int // frame width
	frameHeight int // frame height
	frameX      int // starting point
	frameY      int // x + h
	targetX     float64
	targetY     float64
	scale       float64
}

func (t *Tiles) Update(tick uint) error {
	pace, count := 5, 1
	frame := (int(tick) / pace) % count
	log.Println("tick: ", tick, "frame: ", frame, "scale:", t.scale,
		"h:", float64(t.frameHeight)*t.scale,
		"w:", float64(t.frameWidth)*t.scale,
	)

	t.frameX = frame*t.frameWidth + 10

	return nil
}

func (t *Tiles) Draw(target *ebiten.Image) error {
	// Options for drawing image
	//	opts.GeoM.Translate(t.targetX, t.targetY)

	targetW, _ := target.Size()
	repeat := targetW / t.frameWidth

	for j := 0; j < repeat; j++ {
		opts := &ebiten.DrawImageOptions{}
		tx := float64(t.frameWidth * j)
		opts.GeoM.Translate(tx, t.targetY)

		target.DrawImage(t.img, opts)

	}
	//opts.GeoM.Scale(2, 2)

	/*
		x0, y0 := bg.frameX, bg.frameY
		x1, y1 := x0 + bg.frameWidth - 100, y0 + bg.frameHeight

		// Crop spritesheet
		r := image.Rect(x0, y0, x1, y1)

		sub := bg.img.SubImage(r).(*ebiten.Image)


		target.DrawImage(sub, opts)

	*/
	return nil
}
