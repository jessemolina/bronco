package images

import (
	_ "embed"
)

var (
	//go:embed horse/Walk.png
	HorseWalk_png []byte

	//go:embed horse/Jump.png
	HorseJump_png []byte
)
