package objects

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/assets/images/obstacles"
	"github.com/jessemolina/bronco/pkg/animate"
)

func NewObstacle(screenWidth int, screenHeight int) Object {

	frameW, frameH := obstacles.Rock05.Image.Size()

	diffY := frameH + 60
	//targetY := float64(screenHeight/2)
	targetY := float64(screenHeight - diffY)
	targetX := float64(screenWidth)
	//targetX := float64(screenWidth / 2)

	anm := &animate.Animation{
		Img:         obstacles.Rock05.Image,
		FrameWidth:  frameW,
		FrameHeight: frameH,
		TargetX:     targetX,
		TargetY:     targetY,
		Pace:        15,
		Scale:       2,
	}

	o := &Obstacle{
		anm,
		screenWidth,
		screenHeight,
	}

	return o
}

// Import images that are already decoded.
type Obstacle struct {
	animation    *animate.Animation
	screenWidth  int
	screenHeight int
}

func (o *Obstacle) Animation(set int) {
	switch set {
	case 0:
		o.animation.Pace = 0
	case 1:
		o.animation.Pace = 2
	case 2:
		o.animation.Pace = 0
	}

}

func (o *Obstacle) Coordinates() image.Rectangle {
	return o.animation.Rectangle()
}

func (o *Obstacle) Update(tick uint) error {
	//o.animation.UpdateScrollWidth(float64(o.screenWidth)+60, -1)
	direction := -1
	o.animation.TargetX += (o.animation.Pace * float64(direction))
	max := float64(o.screenWidth * direction)
	if o.animation.TargetX < max {
		o.animation.TargetX = float64(o.screenWidth)
	}

	return nil
}

func (o *Obstacle) Draw(target *ebiten.Image) error {
	// create as many tiles to match the screenWidth.
	// Multiply by 2 to enable scrolling.
	//targetW, _ := target.Size()
	//repeat := (targetW / o.animate.FrameWidth) * 2
	repeat := 1

	o.animation.DrawSequenceX(target, repeat)
	return nil
}
