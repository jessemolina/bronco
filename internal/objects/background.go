package objects

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/assets/images/background"
	"github.com/jessemolina/bronco/pkg/animate"
)

func NewBackground(screenWidth int, screenHeight int) Object {
	// set the frame to match the size of the bg image.
	frameW, frameH := background.Prairie.Image.Size()

	// determine how much the bg image should scale to match
	// the screen's height.
	scale := float64(screenHeight) / float64(frameH)

	anm := &animate.Animation{
		Img:         background.Prairie.Image,
		FrameWidth:  frameW,
		FrameHeight: frameH,
		FrameX:      0,
		FrameY:      0,
		TargetX:     0,
		TargetY:     0,
		Scale:       scale,
		Pace:        2,
	}

	bg := &Background{
		anm,
		screenWidth,
		screenHeight,
	}

	return bg
}

// Import images that are already decoded.
type Background struct {
	animation    *animate.Animation
	screenWidth  int
	screenHeight int
}

func (bg *Background) Update(tick uint) error {
	// TODO Switch implementation for status of the horse.
	// It will update the ebiten sprite.
	// bg.move should implement logic for moving the object.
	// every sprite will have it's on pace and formula for
	// change the targetX and targetY and at what pace.
	max := float64(bg.animation.FrameWidth) * bg.animation.Scale

	bg.animation.UpdateScrollWidth(max, -1)

	//bg.move()
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

	offsetX, offsetY := bg.animation.TargetX, 1.0

	const repeat = 4
	for j := 0; j < repeat; j++ {
		for i := 0; i < repeat; i++ {
			opts := &ebiten.DrawImageOptions{}
			// draws the images next to each other
			opts.GeoM.Scale(bg.animation.Scale, bg.animation.Scale)
			tx := float64(bg.animation.FrameWidth*j) * bg.animation.Scale

			//tx := float64(bg.frameWidth * i)
			opts.GeoM.Translate(tx, 0)

			opts.GeoM.Translate(offsetX, offsetY)
			target.DrawImage(bg.animation.Img, opts)
			/*
				message := fmt.Sprintf(output,
					bg.frameWidth, bg.frameHeight,
					bg.frameX, bg.frameY,
					bg.targetX, bg.targetY,
					bg.scale,
				)

				ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
				ebitenutil.DebugPrint(target, message)
			*/
		}

	}

	return nil
}

// Horse type with image and position.
func (bg *Background) Coordinates() image.Rectangle {
	return bg.animation.Rectangle()
}

func (bg *Background) Animation(set int) {
	switch set {
	case 0:
		bg.animation.Pace = 0
	case 1:
		bg.animation.Pace = 2
	case 2:
		bg.animation.Pace = 0
	}
}
