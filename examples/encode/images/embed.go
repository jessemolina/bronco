package images

import (
	_ "embed"
)


var (
	//go:embed walk.png
	Walk_png []byte

	//go:embed runner.png
	Runner_png []byte

	//go:embed Prairie.png
	Prairie_png []byte

	//go:embed Prairie02.png
	Floor_png []byte

	//go:embed tile.png
	Tile_png []byte
)
