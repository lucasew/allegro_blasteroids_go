package elements

import (
    "github.com/lucasew/allegro_blasteroids_go/point"
    "github.com/dradtke/go-allegro/allegro"
    "github.com/dradtke/go-allegro/allegro/primitives"
    "github.com/lucasew/golog"
    "fmt"
)

var blog = golog.Default.NewLogger("bullet")

type Bullet struct {
    Power int
    Speed float32
    Color allegro.Color
    Position *point.HeadedPoint
}

func (a Bullet) ToString() string {
    return fmt.Sprintf("Bullet(%.2f px/s pwr %d %s)", a.Speed, a.Power, a.Position.ToString())
}

func (a Bullet) GetPosition() point.Point {
    return a.Position.ToPoint()
}

func (a Bullet) Draw() {
    var t allegro.Transform
    t.Identity()
    t.Rotate(a.Position.Heading.Heading)
    t.Translate(a.Position.X, a.Position.Y)
    allegro.UseTransform(&t)
    primitives.DrawFilledRectangle(
        primitives.Point{X: -1, Y: -1},
        primitives.Point{X: 1, Y: 1},
        a.Color,
    )
}

func (a Bullet) DangerRadius() float32 {
    return 1
}

func (a Bullet) IsDead() bool {
    return a.Power <= 0
}

func (a Bullet) GetPower() int {
    return a.Power
}

func (a Bullet) GetLife() int {
    return a.Power
}

func (a *Bullet) Tick(tick float32, w, h int) {
    a.Position.GoAhead(tick*a.Speed)
    a.Position.FixPosition(w, h)
}

func (a Bullet) Hurt(amount int) {
    a.Power = 0
}

func (a Bullet) Die() {}
