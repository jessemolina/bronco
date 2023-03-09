package objects

import (
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jessemolina/bronco/internal/assets/images/horse"
	"github.com/jessemolina/bronco/pkg/animate"
)

var message = `g.count: %v
x0, y0: (%v,%v)
x1, y1: (%v,%v)
i: %v
size: %v, %v
`

// TODO replace tx tx parameters for new horse
func NewHorse(screenWidth int, screenHeight int) Object {

	frameW, frameH := horse.Walk.Image.Size()

	frameW = frameW / horse.Walk.Specs.Frames

	animate := &animate.Animation{
		Img:         horse.Walk.Image,
		FrameCount:  horse.Walk.Specs.Frames,
		FrameWidth:  frameW,
		FrameHeight: frameH,
		FrameX:      0,
		FrameY:      0,
		TargetX:     50,
		TargetY:     185,
		Pace:        4,
	}

	h := &Horse{
		animation:      animate,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}

	return h
}

// Import images that are already decoded.
type Horse struct {
	animation *animate.Animation
	screenWidth  int
	screenHeight int
	isJump       bool
}


// Method for Updating
func (h *Horse) Update(tick uint) error {
	// Pace determines how many ticks will occur before switching to next frame.
	// Count is the number of frames in the sprite sheet.
	pace, count := int(h.animation.Pace), h.animation.FrameCount
	frame := (int(tick) / pace) % count
	//log.Println("tick: ", tick, "frame: ", frame)

	h.animation.FrameX = frame * h.animation.FrameWidth

	// TODO catch key press
	key := ebiten.KeySpace
	h.isJump = ebiten.IsKeyPressed(key)

	return nil
}

// Method for Drawing
func (h *Horse) Draw(target *ebiten.Image) error {

	// Options for drawing image
	opts := &ebiten.DrawImageOptions{}
	// TODO dynamically update the targetX and Y for horse
	if h.isJump {
		opts.GeoM.Translate(h.animation.TargetX, h.animation.TargetY-30)
	} else {
		opts.GeoM.Translate(h.animation.TargetX, h.animation.TargetY)
	}

	opts.GeoM.Scale(2, 2)

	// coordinates
	// x0, y0 := int(h.x), int(h.x)+h.height
	// x1, y1 := x0 + h.width, y0 + h.height
	//
	x0, y0 := h.animation.FrameX, h.animation.FrameY
	x1, y1 := x0+h.animation.FrameWidth, y0+h.animation.FrameHeight

	// Crop spritesheet
	r := image.Rect(x0, y0, x1, y1)

	sub := h.animation.Img.SubImage(r).(*ebiten.Image)

	target.DrawImage(sub, opts)

	message := fmt.Sprintf("jump: %v", h.isJump)
	ebitenutil.DebugPrint(target, message)

	return nil
}

// Horse type with image and position.
func (h *Horse) Coordinates() image.Rectangle {
	return h.animation.Rectangle()
}
