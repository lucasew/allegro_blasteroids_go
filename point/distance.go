package point

import "math"

func (p Point) GetDistance(q Point) float32 {
	return float32(
		math.Sqrt(
			math.Pow(float64(p.X-q.X), 2) +
				math.Pow(float64(p.Y-q.Y), 2),
		),
	)
}

func (p HeadedPoint) GetDistance(q HeadedPoint) float32 {
	return p.ToPoint().GetDistance(q.ToPoint())
}
