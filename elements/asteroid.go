package elements

import (
    "github.com/lucasew/allegro_blasteroids_go/point"
    "github.com/dradtke/go-allegro/allegro"
    "github.com/dradtke/go-allegro/allegro/primitives"
    "github.com/lucasew/golog"
    "fmt"
)

var alog = golog.Default.NewLogger("asteroid")

var asteroidPoints = []point.Point{
    {X: -20,Y: 20},
    {X: -25,Y: 5},
    {X: -25,Y: -10},
    {X: -5,Y: -10},
    {X: -10,Y: -20},
    {X: 5,Y: -20},
    {X: 20,Y: -10},
    {X: 20,Y: -5},
    {X: 0,Y: 0},
    {X: 20,Y: 10},
    {X: 10,Y: 20},
    {X: 0,Y: 15},
}

type Asteroid struct {
    Scale float32
    RotationSpeed float32
    Health int
    Speed float32
    Color allegro.Color
    Position *point.HeadedPoint
}

func (a Asteroid) ToString() string {
    return fmt.Sprintf("Asteroid(%.2fX  %.2f rad/s %.2f px/s + %d %s)", a.Scale, a.RotationSpeed, a.Speed, a.Health, a.Position.ToString())
}

func (a Asteroid) GetPosition() point.Point {
    return point.Point{
        X: a.Position.X,
        Y: a.Position.Y,
    }
}

func (a Asteroid) Draw() {
    var t allegro.Transform
    t.Identity()
    t.Rotate(a.Position.Heading.Heading)
    t.Translate(a.Position.X, a.Position.Y)
    allegro.UseTransform(&t)
    for i := 0; i < len(asteroidPoints); i++ {
        primitives.DrawLine(
            asteroidPoints[i].Multiply(a.Scale).ToPrimitivesPoint(),
            asteroidPoints[(i + 1)%len(asteroidPoints)].Multiply(a.Scale).ToPrimitivesPoint(),
            a.Color,
            2,
        )
    }
}

func (a Asteroid) DangerRadius() float32 {
    return 22 * a.Scale
}

func (a Asteroid) IsDead() bool {
    return a.Health <= 0
}

func (a *Asteroid) Hurt(howmuch int) {
    a.Health -= howmuch
}

func (a Asteroid) GetPower() int {
    return int(a.Speed) / 10
}

func (a Asteroid) GetLife() int {
    return a.Health
}

func (a *Asteroid) Tick(tick float32, w, h int) {
    a.Position.GoAhead(tick*a.Speed)
    a.Position.FixPosition(w, h)
    a.Position.Turn(tick*a.RotationSpeed)
}

func (a Asteroid) Die() {}
