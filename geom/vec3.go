package geom

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Vec2() Vec2 {
	return Vec2{v.X, v.Y}
}
