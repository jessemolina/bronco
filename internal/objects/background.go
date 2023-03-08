package objects

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/assets/images/background"
)

func NewBackground(tx float64, ty float64) Object {

	w, h := background.Prairie.Image.Size()
	frames := 1

	w = w / frames

	bg := &Background{
		img:         background.Prairie.Image,
		frameWidth:  w,
		frameHeight: h,
		frameX:      0,
		frameY:      0,
		targetX:     tx,
		targetY:     ty,
	}

	return bg
}

// Import images that are already decoded.
type Background struct {
	img         *ebiten.Image
	frameWidth  int // frame width
	frameHeight int // frame height
	frameX      int // starting point
	frameY      int // x + h
	targetX     float64
	targetY     float64
}

func (bg *Background) Update(tick uint) error {
	pace, count := 5, 1
	frame := (int(tick) / pace) % count
	log.Println("tick: ", tick, "frame: ", frame)

	bg.frameX = frame * bg.frameWidth + 10


	return nil
}

func (bg *Background) Draw(target *ebiten.Image) error {
	// Options for drawing image
	opts := &ebiten.DrawImageOptions{}
	//opts.GeoM.Translate(bg.targetX, bg.targetY)
	opts.GeoM.Scale(2, 2)

	x0, y0 := bg.frameX, bg.frameY
	x1, y1 := x0 + bg.frameWidth - 100, y0 + bg.frameHeight

	// Crop spritesheet
	r := image.Rect(x0, y0, x1, y1)

	sub := bg.img.SubImage(r).(*ebiten.Image)


	target.DrawImage(sub, opts)

	return nil
}
