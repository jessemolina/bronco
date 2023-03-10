package banner

import (
	_ "embed"
	_ "image/png"
	"log"

	"github.com/jessemolina/bronco/pkg/sprites"
)

const specJSON = "internal/assets/images/banner/banner.json"

var (
	Title *sprites.Sprites = &sprites.Sprites{Name: "Title.png"}

	//go:embed Title.png
	title_png []byte
)

func init() {
	sheet, err := sprites.DecodeSpecs(specJSON)
	if err != nil {
		log.Fatalf("Unable to unmarshall specs: %v", err)
	}

	Title.Initialize(title_png, sheet)
}
