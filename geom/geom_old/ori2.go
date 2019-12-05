package geom

import (
	"math"
	"math/rand"
)

type Ori2 struct {
	X, Y, Theta float64
}

func (o Ori2) Vec2() Vec2 {
	return Vec2{o.X, o.Y}
}

func (a Ori2) Plus(b Ori2) Ori2 {
	return Ori2{a.X + b.X, a.Y + b.Y, a.Theta + b.Theta}
}

func (a *Ori2) PlusEquals(b Ori2) {
	a.X += b.X
	a.Y += b.Y
	a.Theta += b.Theta
}

func (a *Ori2) MinusEquals(b Ori2) {
	a.X -= b.X
	a.Y -= b.Y
	a.Theta -= b.Theta
}

func (a Ori2) Minus(b Ori2) Ori2 {
	return Ori2{a.X - b.X, a.Y - b.Y, a.Theta - b.Theta}
}

func (a Ori2) Times(b Ori2) Ori2 {
	return Ori2{a.X * b.X, a.Y * b.Y, a.Theta * b.Theta}
}

func (a Ori2) Divide(b Ori2) Ori2 {
	return Ori2{a.X / b.X, a.Y / b.Y, a.Theta / b.Theta}
}

func (a Ori2) Dot(b Ori2) float64 {
	return a.X*b.X + a.Y*b.Y + a.Theta*b.Theta
}

func (o Ori2) Scaled(f float64) Ori2 {
	return Ori2{o.X * f, o.Y * f, o.Theta * f}
}

func (o Ori2) Mat33Transform() Mat33 {
	c := math.Cos(o.Theta)
	s := math.Sin(o.Theta)
	return Mat33{
		{c, -s, o.X},
		{s, c, o.Y},
		{0, 0, 1},
	}
}

func (o Ori2) Mat33InvTransform() Mat33 {
	c := math.Cos(o.Theta)
	s := math.Sin(o.Theta)
	return Mat33{
		{c, s, -(c*o.X + s*o.Y)},
		{-s, c, -(-s*o.X + c*o.Y)},
		{0, 0, 1},
	}
}

func (a Ori2) ToWorld(b Ori2) Ori2 {
	s := math.Sin(a.Theta)
	c := math.Cos(a.Theta)
	return Ori2{
		c*b.X - s*b.Y + a.X,
		s*b.X + c*b.Y + a.Y,
		a.Theta + b.Theta,
	}
}

// TODO confirm this
func (a Ori2) ToLocal(b Ori2) Ori2 {
	s := math.Sin(a.Theta)
	c := math.Cos(a.Theta)
	b.X -= a.X
	b.Y -= a.Y
	return Ori2{
		c*b.X + s*b.Y,
		-s*b.X + c*b.Y,
		b.Theta - a.Theta,
	}
}

func Ori2Rand(r Rect) Ori2 {
	v := Vec2Rand(r)
	return Ori2{
		v.X,
		v.Y,
		rand.Float64() * 2. * math.Pi,
	}
}
