package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	img         *ebiten.Image
	frameWidth  int // frame width
	frameHeight int // frame height
	frameX      int // x coordinate on frame
	frameY      int // y coordinate on frame
	targetX     float64 // x coordinate on target (i.e. screen)
	targetY     float64 // y coordinate on target (i.e. screen)
	scale       float64 // scale of image
	pace        float64
}
