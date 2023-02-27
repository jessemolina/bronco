package objects

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/resources/images"
	"github.com/jessemolina/bronco/pkg/tools"
)

func NewHorse(tx, ty float64) *Horse {

	return &Horse{
		tx: tx,
		ty: ty,
	}
}

// Import images that are already decoded.
type Horse struct {
	img *ebiten.Image
	w   int // frame width
	h   int // frame height
	x   float64
	y   float64
	tx  float64
	ty  float64
}

// Horse type with image and position.

// Method for Updating
func (h *Horse) Update() error {
	return nil
}

// Method for Drawing
func (h *Horse) Draw(target *ebiten.Image) error {

	// Decode into Ebiten Image
	img, err := tools.DecodeImage(images.HorseWalk_png)
	if err != nil {
		return err
	}

	// Options for drawing image
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(h.tx, h.ty)

	target.DrawImage(img, opts)

	return nil
}
