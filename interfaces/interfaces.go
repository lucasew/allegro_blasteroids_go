package interfaces

import (
	"github.com/lucasew/allegro_blasteroids_go/point"
)

type Drawable interface {
	Draw()
}

type Stringable interface {
	ToString() string
}

type Positionable interface {
	GetPosition() point.Point
}

type Tickable interface {
	// w and h is the screen size
	Tick(tick float32, w int, h int)
}

type Collisionable interface {
	Positionable
	DangerRadius() float32
	IsDead() bool
	GetPower() int
	GetLife() int
	Hurt(amount int)
	Die()
}

type ListNodePayload interface {
	Drawable
	Stringable
	Collisionable
	Tickable
}
