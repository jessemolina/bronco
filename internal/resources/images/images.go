package images

import (
	_ "embed"
)

var (
	//go:embed horse/Walk.png
	HorseWalk_png []byte

	//go:embed horse/Jump.png
	HorseJump_png []byte

	//go:embed horse/Jump.png
	HorseIdle_png []byte

	//go:embed background/Prairie.png
	BgPrairie_png []byte
)
