package objects

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Object interface {
	Update(uint) error
	Draw(*ebiten.Image) error
	Coordinates() image.Rectangle
	Animation(int)
}
