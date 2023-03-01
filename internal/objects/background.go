package objects

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/resources/images"
	"github.com/jessemolina/bronco/pkg/tools"
)

func NewBackground(tx float64, ty float64) Object {

	// Decode into Ebiten Image
	img, err := tools.DecodeImage(images.BgPrairie_png)
	if err != nil {
		log.Fatalf("Unable to decode Horse: %v", err)
	}

	w, h := img.Size()
	frames := 1

	w = w / frames

	bg := &Background{
		img:         img,
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
	return nil
}

func (bg *Background) Draw(target *ebiten.Image) error {
	// Options for drawing image
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(bg.targetX, bg.targetY)

	target.DrawImage(bg.img, opts)

	return nil
}
