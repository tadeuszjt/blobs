package geom

type Vec3 struct {
	X, Y, Z float64
}

func (a Vec3) Plus(b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vec3) Minus(b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a *Vec3) PlusEquals(b Vec3) {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
}
