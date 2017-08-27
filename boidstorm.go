package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/rect"
	"github.com/glesica/boidstorm/mechanics/navigation"
	"github.com/glesica/boidstorm/mover"
	"github.com/glesica/boidstorm/swarm"
	"github.com/glesica/boidstorm/view"
)

// TODO: This method should move to the mechanics.T type
func run() {
	f := view.NewIMDrawFrame(1024, 768)
	r := rect.New(0, 0, 1024, 768)
	s := swarm.Random(100, r)
	for f.Active() {
		s = s.Update(func(b boid.T) boid.T {
			return navigation.UpdateBoid(s, b).Move(mover.Next(b, 1.0))
		})
		s.Draw(f)
		f.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
