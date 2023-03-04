package tiles

import (
	_ "embed"
	_ "image/png"
	"log"

	"github.com/jessemolina/bronco/pkg/sprites"
)

const specJSON = "internal/assets/images/background/tiles.json"

var (
	PrairieTile *sprites.Sprites = &sprites.Sprites{Name: "PrairieTile.png"}

	//go:embed PrairieTile.png
	Prairie_png []byte
)

func init() {
	sheet, err := sprites.DecodeSpecs(specJSON)
	if err != nil {
		log.Fatalf("Unable to unmarshall specs: %v", err)
	}

	PrairieTile.Initialize(Prairie_png, sheet)
}
