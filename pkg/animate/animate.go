package animate

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


const (

	right direction = 1
	left  direction = -1
	down  direction = 1
	up    direction = -1


	output = `
frameWidth, frameHeight:\t%v x %v
frameX, frameY:\t%v x %v
targetX, taregetY:\t%v x %v
scale:\t%v
pace:\t%v
`
)

type direction int

func (d direction) invert() direction {
	return -d
}

type Animation struct {
	Img         *ebiten.Image
	FrameWidth  int     // frame width
	FrameHeight int     // frame height
	FrameX      int     // x coordinate on frame
	FrameY      int     // y coordinate on frame
	TargetX     float64 // x coordinate on target (i.e. screen)
	TargetY     float64 // y coordinate on target (i.e. screen)
	Scale       float64 // scale of image
	Pace        float64
}

// Prints the animation struct on screen.
func (a *Animation) DebugMessage(target *ebiten.Image, targetX int, targetY int) {
	message := fmt.Sprintf(output,
		a.FrameWidth, a.FrameHeight,
		a.FrameX, a.FrameY,
		a.TargetX, a.TargetY,
		a.Scale,
		a.Pace,
	)
	ebitenutil.DebugPrintAt(target, message, targetX, targetY)
}


// animate

// move

// multiply
func (a *Animation) multiply(target *ebiten.Image, count int) {
	for i := 0; i < count; i++ {
		opts := &ebiten.DrawImageOptions{}
		tx := float64(a.FrameWidth * i)
		opts.GeoM.Translate(tx, a.TargetY)
		target.DrawImage(a.Img, opts)
	}
}

// scroll (direction, pace)

// jump
