package main

import (
	"github.com/glesica/boidstorm/view"
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/boidstorm/swarm"
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/mover"
)

// TODO: This method should move to the game.T type
func run() {
	v := view.New(1024, 768)
	s := swarm.Random(100, 0, 1024, 0, 768)
	for !v.Window().Closed() {
		s.Update(func(b *boid.T) *boid.T {
			return b.Moved(mover.Next(b, 1.0))
		})
		v.Swarm(s)
		v.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
