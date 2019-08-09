package collision

import (
	"github.com/lucasew/allegro_blasteroids_go/interfaces"
)

func CheckCollision(a, b interfaces.Collisionable) int {
	distance := a.GetPosition().GetDistance(b.GetPosition())
	pts := 0
	if distance < a.DangerRadius()+b.DangerRadius() {
		pts = a.GetPower() + b.GetPower()
		a.Hurt(b.GetPower())
		b.Hurt(a.GetPower())
	}
	return pts
}
