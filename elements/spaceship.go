package elements

import (
    "github.com/lucasew/allegro_blasteroids_go/point"
    "github.com/dradtke/go-allegro/allegro"
    "github.com/dradtke/go-allegro/allegro/primitives"
    "github.com/lucasew/golog"
    "fmt"
    "math"
)

var slog = golog.Default.NewLogger("spaceship")
const headingStep float32 = math.Pi/18

type Spaceship struct {
    Health int
    Speed float32
    Position *point.HeadedPoint
}

func (a Spaceship) Color() allegro.Color {
    return allegro.MapRGB(255, 255, 0)
}

func (a *Spaceship) Tick(tick float32, w, h int) {
    a.Position.FixPosition(w, h)
}

func (a *Spaceship) MoveAhead() {
    a.Position.GoAhead(a.Speed)
}

func (a *Spaceship) MoveReverse() {
    a.Position.GoAhead(-a.Speed)
}

func (a *Spaceship) TurnLeft() {
    a.Position.Turn(-headingStep)
}

func (a *Spaceship) TurnRight() {
    a.Position.Turn(headingStep)
}

func (a Spaceship) ToString() string {
    return fmt.Sprintf("Spaceship(%.2f px/s + %d %s)", a.Speed, a.Health, a.Position.ToString())
}

func (a Spaceship) GetPosition() point.Point {
    return point.Point{
        X: a.Position.X,
        Y: a.Position.Y,
    }
}

func (a Spaceship) Draw() {
    slog.Info("draw")
    var t allegro.Transform
    t.Identity()
    t.Rotate(a.Position.Heading.Heading)
    t.Translate(a.Position.X, a.Position.Y)
    allegro.UseTransform(&t)
    primitives.DrawLine(primitives.Point{X: -8, Y: 9}, primitives.Point{X: 0, Y: -11}, a.Color(), 2)
    primitives.DrawLine(primitives.Point{X: 0, Y: -11}, primitives.Point{X: 8, Y: 9}, a.Color(), 2)
    primitives.DrawLine(primitives.Point{X: -6, Y: 4}, primitives.Point{X: -1,Y: 4}, a.Color(), 2)
    primitives.DrawLine(primitives.Point{X: 6, Y: 4}, primitives.Point{X: 1, Y: 4}, a.Color(), 2)
}

func (a Spaceship) DangerRadius() float32 {
    return 10
}

func (a Spaceship) IsDead() bool {
    return a.Health <= 0
}

func (a *Spaceship) Hurt(howmuch int) {
    if (howmuch > 2) {
        howmuch = 0 // Não levar dano das próprias bullets
    }
    slog.Warn("Nave levou %d de dano!", howmuch)
    a.Health -= howmuch
}

func (a Spaceship) GetPower() int {
    return 1
}

func (a Spaceship) GetLife() int {
    return a.Health
}

func (a Spaceship) Die() {}
