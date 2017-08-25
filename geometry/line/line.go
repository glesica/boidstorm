package line

import "github.com/glesica/boidstorm/geometry/vector"

type T interface {
	End() vector.T
	//Mid() vector.T
	Start() vector.T
}

type t struct {
	end, start vector.T
}

func New(start, end vector.T) T {
	return &t{
		end:   end,
		start: start,
	}
}

func (l *t) End() vector.T {
	return l.end
}

func (l *t) Start() vector.T {
	return l.start
}
