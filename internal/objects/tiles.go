package objects

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/assets/images/tiles"
	"github.com/jessemolina/bronco/pkg/animate"
)

func NewTiles(screenWidth int, screenHeight int, floorOffset int) Object {
	frameW, frameH := tiles.Prairie02.Image.Size()

	targetY := float64(screenHeight - frameH - floorOffset)
	targetX := float64(0)
	//targetX := float64(screenWidth / 2)

	anm := &animate.Animation{
		Img:         tiles.Prairie02.Image,
		FrameWidth:  frameW,
		FrameHeight: frameH,
		TargetX:     targetX,
		TargetY:     targetY,
		Pace:        2,
		Scale:       2,
	}

	t := &Tiles{
		anm,
		screenWidth,
		screenHeight,
	}

	return t
}

// Import images that are already decoded.
type Tiles struct {
	animation    *animate.Animation
	screenWidth  int
	screenHeight int
}

func (t *Tiles) Update(tick uint) error {
	t.animation.UpdateScrollWidth(float64(t.screenWidth), -1)

	return nil
}

func (t *Tiles) Draw(target *ebiten.Image) error {
	// create as many tiles to match the screenWidth.
	// Multiply by 2 to enable scrolling.
	targetW, _ := target.Size()
	repeat := (targetW / t.animation.FrameWidth) * 2

	t.animation.DrawSequenceX(target, repeat)
	return nil
}

// Horse type with image and position.
func (t *Tiles) Coordinates() image.Rectangle {
	return t.animation.Rectangle()
}

func (t *Tiles) Animation(s string) {
	switch s {
	case "stop":
		t.animation.Pace = 0
	case "start":
		t.animation.Pace = 2
	case "faster":
		t.animation.Pace += 1
	case "slower":
		t.animation.Pace -= 1
	}
}
