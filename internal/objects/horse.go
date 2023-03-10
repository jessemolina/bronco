package objects

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
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

	//targetY := float64(screenHeight/2)
	targetY := float64(185)
	targetX := float64(50)

	animate := &animate.Animation{
		Img:         horse.Walk.Image,
		FrameCount:  horse.Walk.Specs.Frames,
		FrameWidth:  frameW,
		FrameHeight: frameH,
		FrameX:      0,
		FrameY:      0,
		TargetX:     targetX,
		TargetY:     targetY,
		Pace:        4,
	}

	h := &Horse{
		animation:    animate,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}

	return h
}

// Import images that are already decoded.
type Horse struct {
	animation    *animate.Animation
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
		h.jump()
		opts.GeoM.Translate(h.animation.TargetX+30, h.animation.TargetY-30)
		//opts.GeoM.Translate(h.animation.TargetX, h.animation.TargetY)
	} else {
		h.walk()
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

	//message := fmt.Sprintf("jump: %v", h.isJump)
	//ebitenutil.DebugPrint(target, message)

	return nil
}

// Horse type with image and position.
func (h *Horse) Coordinates() image.Rectangle {
	return h.animation.Rectangle()
}

func (h *Horse) Animation(set int) {
	switch set {
	case 0:
		h.idle()
	case 1:
		h.walk()
	case 2:
		h.idle()
	}
}

func (h *Horse) walk() {

	//frameW, frameH := horse.Walk.Image.Size()
	//frameW = frameW / horse.Walk.Specs.Frames

	h.animation.Img = horse.Walk.Image
	//h.animation.FrameCount = horse.Walk.Specs.Frames
	//h.animation.FrameWidth = frameW
	//h.animation.FrameHeight = frameH
	log.Print("horse.Walk")
}

func (h *Horse) idle() {

	//frameW, frameH := horse.IdleLong.Image.Size()
	//frameW = frameW / horse.IdleLong.Specs.Frames

	h.animation.Img = horse.IdleLong.Image
	//h.animation.FrameCount = horse.IdleLong.Specs.Frames
	//h.animation.FrameWidth = frameW
	//h.animation.FrameHeight = frameH
	//h.animation.Pace = 10
	log.Print("horse.Idle")

}

func (h *Horse) jump() {

	//frameW, frameH := horse.Jump.Image.Size()
	//frameW = frameW / horse.Jump.Specs.Frames

	h.animation.Img = horse.Jump.Image
	//h.animation.FrameCount = horse.Jump.Specs.Frames
	//h.animation.FrameWidth = frameW
	//h.animation.FrameHeight = frameH
	log.Print("horse.Jump")

}
