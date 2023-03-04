package background

import (
	_ "embed"
	_ "image/png"
	"log"

	"github.com/jessemolina/bronco/pkg/sprites"
)

const specJSON = "internal/assets/images/background/background.json"

var (
	Prairie *sprites.Sprites = &sprites.Sprites{Name: "Prairie.png"}

	//go:embed Prairie.png
	Prairie_png []byte
)

func init() {
	sheet, err := sprites.DecodeSpecs(specJSON)
	if err != nil {
		log.Fatalf("Unable to unmarshall specs: %v", err)
	}

	Prairie.Initialize(Prairie_png, sheet)
}
