package animate

import (
	"fmt"
)

const (
		output = `
frameWidth, frameHeight:\t%v x %v
frameX, frameY:\t%v x %v
targetX, taregetY:\t%v x %v
scale:\t%v
pace:\t%v
`
)

// Prints the animation struct on screen.
func debugMessage(a *Animation) string {
	message := fmt.Sprintf(output,
		a.frameWidth, a.frameHeight,
		a.frameX, a.frameY,
		a.targetX, a.targetY,
		a.scale,
		a.pace,
	)

	return message
}

type Animation struct {
	frameWidth  int // frame width
	frameHeight int // frame height
	frameX      int // x coordinate on frame
	frameY      int // y coordinate on frame
	targetX     float64 // x coordinate on target (i.e. screen)
	targetY     float64 // y coordinate on target (i.e. screen)
	scale       float64 // scale of image
	pace        float64
}

// animate

// move

// multiply
func (a *Animation) multiply() {

}

// scroll (direction, pace)

// jump
