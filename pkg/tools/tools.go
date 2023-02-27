package tools

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func DecodeImage(b [] byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	ebimg := ebiten.NewImageFromImage(img)

	return ebimg, nil
}
