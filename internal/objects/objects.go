package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Object interface {
	Update() error
	Draw(*ebiten.Image) error
}
