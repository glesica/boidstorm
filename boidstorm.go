package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/rect"
	"github.com/glesica/boidstorm/geometry/vector"
	"github.com/glesica/boidstorm/mechanics/navigator"
	"github.com/glesica/boidstorm/mover"
	"github.com/glesica/boidstorm/swarm"
	"github.com/glesica/boidstorm/view"
)

func run() {
	f := view.NewIMDrawFrame(1024, 768)
	r := rect.New(0, 0, 1024, 768)
	h := vector.New(1024/2, 768/2)
	c := boid.NewConfig().SetAvoidance(2.0).SetWanderlust(0.01, h)
	s := swarm.Random(100, r, c)
	for f.Active() {
		s = s.Update(func(b boid.T) boid.T {
			neighbors := s.Near(b.Position(), 100)
			acceleration := navigator.Apply(b, neighbors)
			return b.Accelerate(acceleration.Clamp(0.1)).Move(mover.Next, 0.1)
		})
		s.Draw(f)
		f.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
