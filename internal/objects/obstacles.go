package objects

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/assets/images/obstacles"
	"github.com/jessemolina/bronco/pkg/animate"
)

func NewObstacle(screenWidth int, screenHeight int) Object {

	frameW, frameH := obstacles.Rock02.Image.Size()

	targetY := float64(screenHeight/2)
	targetX := float64(screenWidth/2)
	//targetX := float64(screenWidth / 2)

	anm := &animate.Animation{
		Img:         obstacles.Rock02.Image,
		FrameWidth:  frameW,
		FrameHeight: frameH,
		TargetX:     targetX,
		TargetY:     targetY,
		Pace: 2,
		Scale: 2,
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
	animate *animate.Animation
	screenWidth int
	screenHeight int

}

func (o *Obstacle) Update(tick uint) error {
	//o.animate.UpdateScrollWidth(float64(o.screenWidth), -1)

	return nil
}

func (o *Obstacle) Draw(target *ebiten.Image) error {
	// create as many tiles to match the screenWidth.
	// Multiply by 2 to enable scrolling.
	// targetW, _ := target.Size()
	// repeat := (targetW / o.animate.FrameWidth) * 2
	repeat := 1

	o.animate.DrawSequenceX(target, repeat)
	return nil
}
