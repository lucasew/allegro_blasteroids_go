package elements

import (
    "math/rand"
    "github.com/dradtke/go-allegro/allegro"
    "github.com/lucasew/allegro_blasteroids_go/point"
)

func NewRandomAsteroid(w, h int) *Asteroid {
    return &Asteroid{
        Scale: float32((20 + rand.Int31n(50)) / 10),
        RotationSpeed: headingStep * float32(rand.Int31n(5)),
        Health: int(rand.Int31n(200)),
        Speed: float32(rand.Intn(50)),
        Color: NewRandomColor(),
        Position: point.NewHeadedPoint(float32(rand.Intn(w)), float32(rand.Intn(h)), 0),
    }
}

func NewRandomColor() allegro.Color {
    return allegro.MapRGB(byte(rand.Intn(255)), byte(rand.Intn(255)), byte(rand.Intn(255)))
}

func (s Spaceship) NewBullet() *Bullet {
    return &Bullet{
        Power: rand.Intn(200),
        Speed: float32(rand.Intn(50)),
        Color: NewRandomColor(),
        Position: point.NewHeadedPoint(s.Position.X, s.Position.Y, s.Position.Heading.Heading),
    }
}
