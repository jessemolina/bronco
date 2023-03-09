package animate

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	Right direction = 1
	Left  direction = -1
	Down  direction = 1
	Up    direction = -1
)

type direction int

func (d direction) invert() direction {
	return -d
}

type Animation struct {
	Img         *ebiten.Image
	FrameCount  int
	FrameWidth  int     // frame width
	FrameHeight int     // frame height
	FrameX      int     // x coordinate on frame
	FrameY      int     // y coordinate on frame
	TargetX     float64 // x coordinate on target (i.e. screen)
	TargetY     float64 // y coordinate on target (i.e. screen)
	Scale       float64 // scale of image
	Pace        float64
}

func (a *Animation) Rectangle() image.Rectangle {
	x0, y0 := int(a.TargetX), int(a.TargetY)
	x1, y1 := x0+a.FrameWidth, y0+a.FrameHeight

	return image.Rect(x0, y0, x1, y1)
}

// Prints the animation struct on screen.
// TODO fix animate debug message
func (a *Animation) DebugMessage(target *ebiten.Image, targetX int, targetY int) {
	output := `
frameWidth, frameHeight:\t%v x %v
frameX, frameY:\t%v x %v
targetX, taregetY:\t%v x %v
scale:\t%v
pace:\t%v
`
	message := fmt.Sprintf(output,
		a.FrameWidth, a.FrameHeight,
		a.FrameX, a.FrameY,
		a.TargetX, a.TargetY,
		a.Scale,
		a.Pace,
	)
	log.Print(message)
	ebitenutil.DebugPrintAt(target, message, targetX, targetY)
}

// animate

// move

// multiply
func (a *Animation) DrawSequenceX(target *ebiten.Image, count int) {
	for i := 0; i < count; i++ {
		opts := &ebiten.DrawImageOptions{}
		tx := float64(a.FrameWidth*i) + a.TargetX
		opts.GeoM.Scale(a.Scale, a.Scale)
		opts.GeoM.Translate(tx, a.TargetY)
		target.DrawImage(a.Img, opts)
	}
}

// scroll (direction, pace)
// TODO fix scrolling direction for left and right
func (a *Animation) UpdateScrollWidth(maxWidth float64, direction int) {
	switch direction {
	case -1:
		a.TargetX += (a.Pace * float64(direction))
	case 1:
		a.TargetX += (a.Pace * float64(direction))
	default:
	}
	//	a.TargetX -= (a.Pace * float64(direction))
	i := float64(direction)
	max := maxWidth * i
	if a.TargetX < max {
		a.TargetX = float64(0)
	}
}

// jump
