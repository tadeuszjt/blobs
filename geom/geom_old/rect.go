package geom

type Rect struct {
	Min, Max Vec2
}

func RectCentre(w, h float64) Rect {
	w /= 2
	h /= 2
	return Rect{
		Vec2{-w, -h},
		Vec2{w, h},
	}
}

func RectOrigin(w, h float64) Rect {
	return Rect{
		Vec2{0, 0},
		Vec2{w, h},
	}
}

func (r Rect) Width() float64 {
	return r.Max.X - r.Min.X
}

func (r Rect) Height() float64 {
	return r.Max.Y - r.Min.Y
}

func (r Rect) Area() float64 {
	return r.Width() * r.Height()
}

func (r Rect) Contains(p Vec2) bool {
	return p.X > r.Min.X &&
		p.X < r.Max.X &&
		p.Y > r.Min.Y &&
		p.Y < r.Max.Y
}

func (r Rect) Vertices() [4]Vec2 {
	return [4]Vec2{
		r.Min,
		Vec2{r.Max.X, r.Min.Y},
		r.Max,
		Vec2{r.Min.X, r.Max.Y},
	}
}
