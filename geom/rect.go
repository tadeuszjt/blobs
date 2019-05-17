package geom

type Rect struct {
	Min, Max Vec2
}

func RectOrigin(w, h float64) Rect {
	return Rect{
		Vec2{0, 0},
		Vec2{w, h},
	}
}

func RectCentered(w, h float64, pos Vec2) Rect {
	wh := w / 2
	hh := h / 2
	return Rect{
		Vec2{pos.X - wh, pos.Y - hh},
		Vec2{pos.X + wh, pos.Y + hh},
	}
}

func (r Rect) Width() float64 {
	return r.Max.X - r.Min.X
}

func (r Rect) Height() float64 {
	return r.Max.Y - r.Min.Y
}

func (r Rect) Contains(v Vec2) bool {
	return v.X >= r.Min.X &&
		v.X <= r.Max.X &&
		v.Y >= r.Min.Y &&
		v.Y <= r.Max.Y
}
