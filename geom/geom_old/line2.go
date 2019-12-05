package geom

import (
	"math"
)

type Line2 struct {
	A, B Vec2
}

func (l Line2) Len2() float64 {
	dx := l.B.X - l.A.X
	dy := l.B.Y - l.A.Y
	return dx*dx + dy*dy
}

func (l Line2) Len() float64 {
	return math.Sqrt(l.Len2())
}

func (l Line2) Midpoint() Vec2 {
	return Vec2{
		(l.A.X + l.B.X) / 2,
		(l.A.Y + l.B.Y) / 2,
	}
}
