package elements

import (
	"fmt"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/primitives"
	"github.com/lucasew/allegro_blasteroids_go/point"
	"github.com/lucasew/golog"
	"math"
	"sync"
)

var slog = golog.Default.NewLogger("spaceship")

const headingStep float32 = math.Pi / 18

type Spaceship struct {
	Health   int
	Speed    float32
	Position *point.HeadedPoint
	mu       sync.Mutex
}

func (a *Spaceship) Color() allegro.Color {
	return allegro.MapRGB(255, 255, 0)
}

func (a *Spaceship) Tick(tick float32, w, h int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Position.FixPosition(w, h)
}

func (a *Spaceship) MoveAhead() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Position.GoAhead(a.Speed)
}

func (a *Spaceship) MoveReverse() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Position.GoAhead(-a.Speed)
}

func (a *Spaceship) TurnLeft() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Position.Turn(-headingStep)
}

func (a *Spaceship) TurnRight() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Position.Turn(headingStep)
}

func (a *Spaceship) ToString() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	return fmt.Sprintf("Spaceship(%.2f px/s + %d %s)", a.Speed, a.Health, a.Position.ToString())
}

func (a *Spaceship) GetPosition() point.Point {
	a.mu.Lock()
	defer a.mu.Unlock()
	return point.Point{
		X: a.Position.X,
		Y: a.Position.Y,
	}
}

func (a *Spaceship) Draw() {
	slog.Info("draw")
	a.mu.Lock()
	defer a.mu.Unlock()
	var t allegro.Transform
	t.Identity()
	t.Rotate(a.Position.Heading.Heading)
	t.Translate(a.Position.X, a.Position.Y)
	allegro.UseTransform(&t)
	primitives.DrawLine(primitives.Point{X: -8, Y: 9}, primitives.Point{X: 0, Y: -11}, a.Color(), 2)
	primitives.DrawLine(primitives.Point{X: 0, Y: -11}, primitives.Point{X: 8, Y: 9}, a.Color(), 2)
	primitives.DrawLine(primitives.Point{X: -6, Y: 4}, primitives.Point{X: -1, Y: 4}, a.Color(), 2)
	primitives.DrawLine(primitives.Point{X: 6, Y: 4}, primitives.Point{X: 1, Y: 4}, a.Color(), 2)
}

func (a *Spaceship) DangerRadius() float32 {
	return 10
}

func (a *Spaceship) IsDead() bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Health <= 0
}

func (a *Spaceship) Hurt(howmuch int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if howmuch > 2 {
		howmuch = 0 // Não levar dano das próprias bullets
	}
	slog.Warn("Nave levou %d de dano!", howmuch)
	a.Health -= howmuch
}

func (a *Spaceship) GetPower() int {
	return 1
}

func (a *Spaceship) GetLife() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Health
}

func (a *Spaceship) Die() {}
