package objects

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jessemolina/bronco/internal/assets/images/background"
)

func NewBackground(screenWidth int, screenHeight int) Object {
	// set the frame to match the size of the bg image.
	frameW, frameH := background.Prairie.Image.Size()

	// determine how much the bg image should scale to match
	// the screen's height.
	scale := float64(screenHeight) / float64(frameH)

	bg := &Background{
		img:         background.Prairie.Image,
		frameWidth:  frameW,
		frameHeight: frameH,
		frameX:      0,
		frameY:      0,
		targetX:     0,
		targetY:     0,
		scale:       scale,
	}

	return bg
}

// Import images that are already decoded.
type Background struct {
	img         *ebiten.Image
	frameWidth  int // frame width
	frameHeight int // frame height
	frameX      int // starting point
	frameY      int // x + h
	targetX     int
	targetY     int
	scale       float64
}

func (bg *Background) Update(tick uint) error {
	// TODO Switch implementation for status of the horse.
	// It will update the ebiten sprite.
	// bg.move should implement logic for moving the object.
	// every sprite will have it's on pace and formula for
	// change the targetX and targetY and at what pace.

	bg.move()
	/*
		pace, count := 5, 1
		frame := (int(tick) / pace) % count
		log.Println("tick: ", tick, "frame: ", frame, "scale:", bg.scale,
			"h:", float64(bg.frameHeight)*bg.scale,
			"w:", float64(bg.frameWidth)*bg.scale,
		)

		bg.frameX = frame*bg.frameWidth + 10

	*/

	return nil
}

func (bg *Background) Draw(target *ebiten.Image) error {
	// Options for drawing image
	//opts := &ebiten.DrawImageOptions{}
	//opts.GeoM.Translate(bg.targetX, bg.targetY)
	//opts.GeoM.Scale(bg.scale, bg.scale)
	//
	offsetX, offsetY := float64(bg.targetX), 1.0

	const repeat = 4
	for j := 0; j < repeat; j++ {
		for i := 0; i < repeat; i++ {
			opts := &ebiten.DrawImageOptions{}
			// draws the images next to each other
			opts.GeoM.Scale(bg.scale, bg.scale)
			tx := float64(bg.frameWidth*j) * bg.scale
			//tx := float64(bg.frameWidth * i)
			opts.GeoM.Translate(tx, 0)

			opts.GeoM.Translate(offsetX, offsetY)
			target.DrawImage(bg.img, opts)
			message := fmt.Sprintf(output,
				bg.frameWidth, bg.frameHeight,
				bg.frameX, bg.frameY,
				bg.targetX, bg.targetY,
				bg.scale,
			)

			//ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
			ebitenutil.DebugPrint(target, message)

		}

	}

	return nil
}

func (bg *Background) move() {
	max := float64(bg.frameWidth) * bg.scale
	// TODO make pace dynamic
	bg.targetX -= 5
	if float64(bg.targetX) < -max {
		bg.targetX = 0
	}
}
