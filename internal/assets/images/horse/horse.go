package horse

import (
	_ "embed"
	_ "image/png"
	"log"

	"github.com/jessemolina/bronco/pkg/sprites"
)

const specJSON = "internal/assets/images/horse/horse.json"

var (
	Idle *sprites.Sprites = &sprites.Sprites{Name: "Idle.png"}
	Jump *sprites.Sprites = &sprites.Sprites{Name: "Jump.png"}
	Walk *sprites.Sprites = &sprites.Sprites{Name: "Walk.png"}

	//go:embed Idle.png
	Idle_png []byte
	//go:embed Jump.png
	Jump_png []byte
	//go:embed Walk.png
	Walk_png []byte
)

func init() {
	sheet, err := sprites.DecodeSpecs(specJSON)
	if err != nil {
		log.Fatalf("Unable to unmarshall specs: %v", err)
	}

	Idle.Initialize(Idle_png, sheet)
	Jump.Initialize(Jump_png, sheet)
	Walk.Initialize(Walk_png, sheet)
}
