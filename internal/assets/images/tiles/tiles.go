package tiles

import (
	_ "embed"
	_ "image/png"
	"log"

	"github.com/jessemolina/bronco/pkg/sprites"
)

const specJSON = "internal/assets/images/tiles/tiles.json"

var (
	Prairie02 *sprites.Sprites = &sprites.Sprites{Name: "Prairie02.png"}

	//go:embed Prairie02.png
	Prairie02_png []byte
)

func init() {
	sheet, err := sprites.DecodeSpecs(specJSON)
	if err != nil {
		log.Fatalf("Unable to unmarshall specs: %v", err)
	}

	Prairie02.Initialize(Prairie02_png, sheet)
}
