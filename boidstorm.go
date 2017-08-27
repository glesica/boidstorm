package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/boidstorm/boid"
	"github.com/glesica/boidstorm/geometry/rect"
	"github.com/glesica/boidstorm/mechanics/navigation"
	"github.com/glesica/boidstorm/mover"
	"github.com/glesica/boidstorm/swarm"
	"github.com/glesica/boidstorm/view"
	"github.com/glesica/boidstorm/geometry/vector"
)

func run() {
	f := view.NewIMDrawFrame(1024, 768)
	r := rect.New(0, 0, 1024, 768)
	c := boid.NewConfig().SetAvoidance(0.5).SetHome(vector.New(1024 / 2, 768 / 2))
	s := swarm.Random(100, r, c)
	for f.Active() {
		s = s.Update(func(b boid.T) boid.T {
			return navigation.Boid(s, b).Move(mover.Next, 1.0)
		})
		s.Draw(f)
		f.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
