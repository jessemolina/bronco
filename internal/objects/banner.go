package objects

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/assets/images/banner"
	"github.com/jessemolina/bronco/pkg/animate"
)

func NewBanner(screenWidth int, screenHeight int) Object {
	// set the frame to match the size of the bg image.
	frameW, frameH := banner.Title.Image.Size()

	// determine how much the bg image should scale to match
	// the screen's height.
	//scale := float64(screenHeight) / float64(frameH)
	//
	tx, ty := float64(screenWidth / 2) - (float64(frameW) / 2), float64(screenHeight / 2) - 128

	anm := &animate.Animation{
		Img:         banner.Title.Image,
		FrameWidth:  frameW,
		FrameHeight: frameH,
		FrameX:      0,
		FrameY:      0,
		TargetX:     tx,
		TargetY:     ty,
		Scale:       1,
		Pace:        2,
	}

	b := &Banner{
		anm,
		screenWidth,
		screenHeight,
	}

	return b
}

// Import images that are already decoded.
type Banner struct {
	animation    *animate.Animation
	screenWidth  int
	screenHeight int
}

func (b *Banner) Update(tick uint) error {
	return nil
}

func (b *Banner) Draw(target *ebiten.Image) error {
	opts := &ebiten.DrawImageOptions{}
	// draws the images next to each other
	opts.GeoM.Scale(b.animation.Scale, b.animation.Scale)
	opts.GeoM.Translate(b.animation.TargetX, b.animation.TargetY)

	//tx := float64(b.animation.FrameWidth) * b.animation.Scale

	//tx := float64(bg.frameWidth * i)
	//opts.GeoM.Translate(tx, 0)

	//opts.GeoM.Translate(offsetX, offsetY)
	target.DrawImage(b.animation.Img, opts)

	return nil
}

// Horse type with image and position.
func (b *Banner) Coordinates() image.Rectangle {
	return b.animation.Rectangle()
}

func (b *Banner) Animation(set int) {
	switch set {
	case 0:
		//b.animation.Img.Clear()
	case 1:
		b.animation.TargetY -= float64(b.screenHeight)
		b.animation.Pace = 2
	case 2:
		b.animation.TargetY += float64(b.screenHeight)
		b.animation.Pace += 1
	}
}
