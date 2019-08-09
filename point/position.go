package point

import (
    "github.com/dradtke/go-allegro/allegro/primitives"
    "math"
    "fmt"
)
type Point struct {
    X, Y float32
}

func (p *Point) FixPosition(w, h int) {
    fw := float32(w)
    fh := float32(h)
    if (p.X > fw) {
        p.X = 0
    }
    if (p.X < 0) {
        p.X = fw
    }
    if (p.Y > fh) {
        p.Y = 0
    }
    if (p.Y < 0) {
        p.Y = fh
    }
}

func (p Point) ToString() string {
    return fmt.Sprintf("Point(%.2f, %.2f)", p.X, p.Y)
}

func (p Point) Multiply(by float32) Point {
    return Point{
        X: p.X * by,
        Y: p.Y * by,
    }
}

func (p *HeadedPoint) GoAhead(distance float32) {
    p.X += distance*float32(math.Sin(float64(p.Heading.Heading)))
    p.Y -= distance*float32(math.Cos(float64(p.Heading.Heading)))
}

func (p Point) ToPrimitivesPoint() primitives.Point {
    return primitives.Point{
        X: p.X,
        Y: p.Y,
    }
}

type Heading struct {
    Heading float32
}

func (h *Heading) Turn(rad float32) {
    h.Heading += rad
}


type HeadedPoint struct {
    Point
    Heading
}

func (p HeadedPoint) ToString() string {
    return fmt.Sprintf("HeadedPoint(%.2f, %.2f, %.2f rad)", p.X, p.Y, p.Heading.Heading)
}

func (p HeadedPoint) ToPoint() Point {
    return Point{
        X: p.X,
        Y: p.Y,
    }
}

func NewHeadedPoint(X, Y, heading float32) *HeadedPoint {
    return &HeadedPoint{
        Point{
            X: X,
            Y: Y,
        },
        Heading{
            Heading: heading,
        },
    }
}
