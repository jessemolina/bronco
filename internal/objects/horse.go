package objects

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/assets/images/horse"
)

var message = `g.count: %v
x0, y0: (%v,%v)
x1, y1: (%v,%v)
i: %v
size: %v, %v
`

/*
// TODO Refactor NewHorse; set default frames per sprites via json file.
   NewHorse(name string, tx float64, ty float64)

   sprites := images.Horse[name]

   horse := &Horse{

   }
*/
func NewHorse(tx float64, ty float64) Object {

	w, h := horse.Walk.Image.Size()

	w = w / horse.Walk.Specs.Frames

	horse := &Horse{
		img:         horse.Walk.Image,
		frameWidth:  w,
		frameHeight: h,
		frameX:      0,
		frameY:      0,
		targetX:     tx,
		targetY:     ty,
	}

	return horse
}

// Import images that are already decoded.
type Horse struct {
	img         *ebiten.Image
	frameWidth  int // frame width
	frameHeight int // frame height
	frameX      int // starting point
	frameY      int // x + h
	targetX     float64
	targetY     float64
}

// Horse type with image and position.

// Method for Updating
func (h *Horse) Update(tick uint) error {
	// Pace determines how many ticks will occur before switching to next frame.
	// Count is the number of frames in the sprite sheet.
	pace, count := 5, 6
	frame := (int(tick) / pace) % count
	//log.Println("tick: ", tick, "frame: ", frame)

	h.frameX = frame * h.frameWidth

	return nil
}

// Method for Drawing
func (h *Horse) Draw(target *ebiten.Image) error {

	// Options for drawing image
	opts := &ebiten.DrawImageOptions{}
	// TODO dynamically update the targetX and Y for horse
	opts.GeoM.Translate(50, 182)
	opts.GeoM.Scale(2, 2)

	// coordinates
	// x0, y0 := int(h.x), int(h.x)+h.height
	// x1, y1 := x0 + h.width, y0 + h.height
	//
	x0, y0 := h.frameX, h.frameY
	x1, y1 := x0 + h.frameWidth, y0 + h.frameHeight

	// Crop spritesheet
	r := image.Rect(x0, y0, x1, y1)

	sub := h.img.SubImage(r).(*ebiten.Image)

	target.DrawImage(sub, opts)

	return nil
}
