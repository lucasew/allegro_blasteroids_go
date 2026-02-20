package elements

import (
	"testing"
	"github.com/lucasew/allegro_blasteroids_go/point"
    "sync"
)

func TestSpaceshipRace(t *testing.T) {
	spaceship := &Spaceship{
		Health:   100,
		Speed:    20,
		Position: point.NewHeadedPoint(40, 40, 1),
	}

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        for i := 0; i < 1000; i++ {
            spaceship.MoveAhead()
        }
    }()

    go func() {
        defer wg.Done()
        for i := 0; i < 1000; i++ {
             _ = spaceship.GetPosition()
        }
    }()

    wg.Wait()
}
