package objects

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/assets/images/tiles"
	"github.com/jessemolina/bronco/pkg/animate"
)

func NewTiles(screenWidth int, screenHeight int) Object {
	frameW, frameH := tiles.Prairie02.Image.Size()

	targetY := float64(screenHeight - frameH)
	targetX := float64(0)
	//targetX := float64(screenWidth / 2)

	anm := &animate.Animation{
		Img:         tiles.Prairie02.Image,
		FrameWidth:  frameW,
		FrameHeight: frameH,
		TargetX:     targetX,
		TargetY:     targetY,
		Pace: 2,
	}

	bg := &Tiles{
		anm,
		screenWidth,
		screenHeight,
	}

	return bg
}

// Import images that are already decoded.
type Tiles struct {
	animate *animate.Animation
	screenWidth int
	screenHeight int
}

func (t *Tiles) Update(tick uint) error {
	t.animate.UpdateScrollWidth(t.screenWidth, -1)

	return nil
}

func (t *Tiles) Draw(target *ebiten.Image) error {
	// create as many tiles to match the screenWidth.
	// Multiply by 2 to enable scrolling.
	targetW, _ := target.Size()
	repeat := (targetW / t.animate.FrameWidth) * 2

	t.animate.DrawSequenceX(target, repeat)
	return nil
}
