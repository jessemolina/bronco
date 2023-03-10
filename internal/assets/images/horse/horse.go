package horse

import (
	_ "embed"
	_ "image/png"
	"log"

	"github.com/jessemolina/bronco/pkg/sprites"
)

const specJSON = "internal/assets/images/horse/horse.json"

var (
	Idle     *sprites.Sprites = &sprites.Sprites{Name: "Idle.png"}
	IdleLong *sprites.Sprites = &sprites.Sprites{Name: "IdleLong.png"}
	Jump     *sprites.Sprites = &sprites.Sprites{Name: "Jump.png"}
	Walk     *sprites.Sprites = &sprites.Sprites{Name: "Walk.png"}

	//go:embed Idle.png
	idle_png []byte
	//go:embed IdleLong.png
	idleLong_png []byte
	//go:embed Jump.png
	jump_png []byte
	//go:embed Walk.png
	walk_png []byte
)

func init() {
	sheet, err := sprites.DecodeSpecs(specJSON)
	if err != nil {
		log.Fatalf("Unable to unmarshall specs: %v", err)
	}

	Idle.Initialize(idle_png, sheet)
	IdleLong.Initialize(idleLong_png, sheet)
	Jump.Initialize(jump_png, sheet)
	Walk.Initialize(walk_png, sheet)
}
