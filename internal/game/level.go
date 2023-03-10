package game

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessemolina/bronco/internal/objects"
)

var Collision bool

const (
	first = 1
)

type level struct {
	Number int
	Obstacles []objects.Object
}

func NewLevel(screenWidth int, screenHeight int) *level {

	l := &level{
		Number: first,
	}

	for i:=0; i < first; i++ {
		o := objects.NewObstacle(screenWidth, screenHeight)
		l.Obstacles = append(l.Obstacles, o)
	}
	log.Printf("obstacles: %v", len(l.Obstacles))

	return l
}

func (l *level) Update(tick uint) {
	// draw X amount of rocks on the floor
	// this is controlled by the level number
	// spread them
	for _, o := range l.Obstacles {
		o.Update(tick)
	}
}

func (l *level) Draw(screen *ebiten.Image) {
	// draw X amount of rocks on the floor
	// this is controlled by the level number
	// spread them
	for _, o := range l.Obstacles {
		o.Draw(screen)
	}
}

func DetectCollision(a image.Rectangle, b image.Rectangle) {

}
/*
func (l *level) name(args) return type {

}
*/

func randomPoints(count int) []image.Point {
	return nil
}

// NewLevel() *level
// Initialize a set of obstacles and add them to the array

// level.Update
// for _, o := range l.Obstacles
// o.Update()
// Accept a image.Point and compare to its own
// var collission false or true

// helper function that checks collision between two boxes

// function that accepts a image.Point
// Each object needs to give a set of Coordinates; aka point.
