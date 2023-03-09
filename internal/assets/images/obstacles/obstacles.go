package obstacles

import (
	_ "embed"
	_ "image/png"
	"log"

	"github.com/jessemolina/bronco/pkg/sprites"
)

const specJSON = "internal/assets/images/obstacles/obstacles.json"

var (
	Rock01 *sprites.Sprites = &sprites.Sprites{Name: "Rock01.png"}
	Rock02 *sprites.Sprites = &sprites.Sprites{Name: "Rock02.png"}
	Rock03 *sprites.Sprites = &sprites.Sprites{Name: "Rock03.png"}
	Rock04 *sprites.Sprites = &sprites.Sprites{Name: "Rock04.png"}
	Rock05 *sprites.Sprites = &sprites.Sprites{Name: "Rock05.png"}
	Rock06 *sprites.Sprites = &sprites.Sprites{Name: "Rock06.png"}

	//go:embed Rock01.png
	rock01_png []byte
	//go:embed Rock02.png
	rock02_png []byte
	//go:embed Rock03.png
	rock03_png []byte
	//go:embed Rock04.png
	rock04_png []byte
	//go:embed Rock05.png
	rock05_png []byte
	//go:embed Rock06.png
	rock06_png []byte
)

func init() {
	sheet, err := sprites.DecodeSpecs(specJSON)
	if err != nil {
		log.Fatalf("Unable to unmarshall specs: %v", err)
	}

	Rock01.Initialize(rock01_png, sheet)
	Rock02.Initialize(rock02_png, sheet)
	Rock03.Initialize(rock03_png, sheet)
	Rock04.Initialize(rock04_png, sheet)
	Rock05.Initialize(rock05_png, sheet)
	Rock06.Initialize(rock06_png, sheet)
}
