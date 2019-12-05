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

func (a Vec2) Dot(b Vec2) float64 {
	return a.X*b.X + a.Y*b.Y
}

func (a Vec2) Cross(b Vec2) float64 {
	return a.X*b.Y - a.Y*b.X
}

func (a Vec2) Len2() float64 {
	return a.X*a.X + a.Y*a.Y
}

func (a Vec2) Len() float64 {
	return math.Sqrt(a.Len2())
}

func (v Vec2) Normal() Vec2 {
	invLen := 1.0 / v.Len()
	return Vec2{v.X * invLen, v.Y * invLen}
}

func (v Vec2) Rotated(rad float64) Vec2 {
	c := math.Cos(rad)
	s := math.Sin(rad)
	return Vec2{c*v.X - s*v.Y, s*v.X + c*v.Y}
}

func Vec2Rand(r Rect) Vec2 {
	return Vec2{
		rand.Float64()*(r.Max.X-r.Min.X) + r.Min.X,
		rand.Float64()*(r.Max.Y-r.Min.Y) + r.Min.Y,
	}
}

func Vec2RandNorm() Vec2 {
	r := rand.Float64() * 2 * math.Pi
	return Vec2{math.Cos(r), math.Sin(r)}
}
