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
	Rock01_png []byte
	//go:embed Rock02.png
	Rock02_png []byte
	//go:embed Rock03.png
	Rock03_png []byte
	//go:embed Rock04.png
	Rock04_png []byte
	//go:embed Rock05.png
	Rock05_png []byte
	//go:embed Rock06.png
	Rock06_png []byte
)

func init() {
	sheet, err := sprites.DecodeSpecs(specJSON)
	if err != nil {
		log.Fatalf("Unable to unmarshall specs: %v", err)
	}

	Rock02.Initialize(rock02_png, sheet)
}
