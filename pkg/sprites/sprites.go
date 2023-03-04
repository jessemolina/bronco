package sprites

import (
	"bytes"
	"encoding/json"
	"image"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
)

// ================================================================
// HELPER FUNCTIONS

func DecodeSpecs(file string) (*SpriteSheet, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	s := &SpriteSheet{}

	err = json.Unmarshal(data, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func DecodeImage(source []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(source))
	if err != nil {
		return nil, err
	}

	ebimg := ebiten.NewImageFromImage(img)

	return ebimg, nil
}

// ================================================================
// TYPES

type ImageSpecs struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Frames int    `json:"frames"`
}

type SpriteSheet struct {
	Images []ImageSpecs `json:"images"`
}

type Sprites struct {
	Name  string
	Image *ebiten.Image
	Specs *ImageSpecs
}

func (s *Sprites) Initialize(source []byte, sheet *SpriteSheet) error {
	for _, image := range sheet.Images {
		if s.Name == image.Name {
			img, err := DecodeImage(source)
			if err != nil {
				return err
			}
			s.Image = img
			s.Specs = &image
		}
	}

	return nil
}
