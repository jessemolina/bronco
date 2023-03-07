package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type direction int

const (
	right direction = 1
	left  direction = -1
	down  direction = 1
	up    direction = -1

	output = `
frameWidth, frameHeight: %v x %v
frameX, frameY: %v x %v
targetX, targetY: %v x %v
scale: %v
`
)

func (d direction) invert() direction {
	return -d
}
type Object interface {
	Update(uint) error
	Draw(*ebiten.Image) error
}
