package objects

import (
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jessemolina/bronco/internal/resources/images"
	"github.com/jessemolina/bronco/pkg/tools"
)

var message = `g.count: %v
x0, y0: (%v,%v)
x1, y1: (%v,%v)
i: %v
size: %v, %v
`

func NewHorse(frames int, tx float64, ty float64) Object {

	// Decode into Ebiten Image
	img, err := tools.DecodeImage(images.HorseWalk_png)
	if err != nil {
		log.Fatalf("Unable to decode Horse: %v", err)
	}

	w, h := img.Size()

	w = w / frames

	horse := &Horse{
		img:         img,
		frameWidth:  w,
		frameHeight: h,
		frameX:      0,
		targetX:     tx,
		targetY:     ty,
	}

	return horse
}

// Import images that are already decoded.
type Horse struct {
	img         *ebiten.Image
	frameWidth  int     // frame width
	frameHeight int     // frame height
	frameX      float64 // starting point
	frameY      float64 // x + h
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
	log.Println("tick: ", tick, "frame: ", frame)

	// TODO update x0 and

	return nil
}

// Method for Drawing
func (h *Horse) Draw(target *ebiten.Image) error {

	// Options for drawing image
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(h.targetX, h.targetY)

	// coordinates
	// x0, y0 := int(h.x), int(h.x)+h.height
	// x1, y1 := x0 + h.width, y0 + h.height
	//
	x0, y0 := 0, 0
	x1, y1 := 72, 72

	// Crop spritesheet
	r := image.Rect(x0, y0, x1, y1)

	sub := h.img.SubImage(r).(*ebiten.Image)

	a, b := h.img.Size()
	stats := fmt.Sprintf(message, "g.count", x0, y0, x1, y1, "i", a, b)
	ebitenutil.DebugPrint(target, stats)

	target.DrawImage(sub, opts)

	return nil
}
