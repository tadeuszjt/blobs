package geom

import (
	"math"
	"math/rand"
)

type Vec2 struct {
	X, Y float64
}

func (a Vec2) Plus(b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

func (a *Vec2) PlusEquals(b Vec2) {
	a.X += b.X
	a.Y += b.Y
}

func (a Vec2) Minus(b Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}

func (v Vec2) Scaled(f float64) Vec2 {
	return Vec2{v.X * f, v.Y * f}
}

func Vec2Rand(r Rect) Vec2 {
	return Vec2{
		rand.Float64()*r.Width() + r.Min.X,
		rand.Float64()*r.Height() + r.Min.Y,
	}
}

func Vec2RandNormal() Vec2 {
	theta := rand.Float64() * 2 * math.Pi
	return Vec2{
		math.Sin(theta),
		math.Cos(theta),
	}
}

func (v Vec2) Len2() float64 {
	return v.X*v.X + v.Y*v.Y
}
