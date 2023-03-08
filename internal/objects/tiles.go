package objects

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/assets/images/tiles"
	"github.com/jessemolina/bronco/pkg/animate"
)

func NewTiles(screenWidth int, screenHeight int) Object {
	frameW, frameH := tiles.Prairie02.Image.Size()

	targetY := float64(screenHeight - frameH)
	targetX := float64(0)
	//targetX := float64(screenWidth / 2)

	anm := &animate.Animation{
		Img:         tiles.Prairie02.Image,
		FrameWidth:  frameW,
		FrameHeight: frameH,
		TargetX:     targetX,
		TargetY:     targetY,
		Pace: 2,
	}

	bg := &Tiles{
		anm,
		screenWidth,
		screenHeight,
	}

	return bg
}

// Import images that are already decoded.
type Tiles struct {
	animate *animate.Animation
	screenWidth int
	screenHeight int
}

func (t *Tiles) Update(tick uint) error {
	t.animate.UpdateScrollWidth(t.screenWidth, -1)
	/*
		pace, count := 5, 1
		frame := (int(tick) / pace) % count
		log.Println("tick: ", tick, "frame: ", frame, "scale:", t.scale,
			"h:", float64(t.frameHeight)*t.scale,
			"w:", float64(t.frameWidth)*t.scale,
		)

		t.frameX = frame*t.frameWidth + 10
	*/

	return nil
}

func (t *Tiles) Draw(target *ebiten.Image) error {
	// Options for drawing image
	/*
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(t.anm.TargetX, t.anm.TargetY)
	target.DrawImage(t.anm.Img, opts)
	*/
	targetW, _ := target.Size()
	repeat := (targetW / t.animate.FrameWidth) * 2

	t.animate.DrawSequenceX(target, repeat)

		//

	//opts.GeoM.Scale(2, 2)

	/*
		x0, y0 := bg.frameX, bg.frameY
		x1, y1 := x0 + bg.frameWidth - 100, y0 + bg.frameHeight

		// Crop spritesheet
		r := image.Rect(x0, y0, x1, y1)

		sub := bg.img.SubImage(r).(*ebiten.Image)


		target.DrawImage(sub, opts)

	*/
	return nil
}
