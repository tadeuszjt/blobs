package geom

type Mat33 [3][3]float64

func Mat33Identity() Mat33 {
	return Mat33{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

func Mat33Scale(sx, sy, sz float64) Mat33 {
	return Mat33{
		{sx, 0, 0},
		{0, sy, 0},
		{0, 0, sz},
	}
}

func Mat33Translation(tx, ty float64) Mat33 {
	return Mat33{
		{1, 0, tx},
		{0, 1, ty},
		{0, 0, 1},
	}
}

func (m Mat33) Scaled(f float64) Mat33 {
	return Mat33{
		{f * m[0][0], f * m[0][1], f * m[0][2]},
		{f * m[1][0], f * m[1][1], f * m[1][2]},
		{f * m[2][0], f * m[2][1], f * m[2][2]},
	}
}

func (a *Mat33) PlusEquals(b Mat33) {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			a[r][c] += b[r][c]
		}
	}
}

func (a Mat33) Times(b Mat33) (ret Mat33) {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			ret[r][c] = a[r][0]*b[0][c] + a[r][1]*b[1][c] + a[r][2]*b[2][c]
		}
	}
	return
}

func (a Mat33) TimesVec3(b Vec3) Vec3 {
	return Vec3{
		a[0][0]*b.X + a[0][1]*b.Y + a[0][2]*b.Z,
		a[1][0]*b.X + a[1][1]*b.Y + a[1][2]*b.Z,
		a[2][0]*b.X + a[2][1]*b.Y + a[2][2]*b.Z,
	}
}

func (a Mat33) TimesVec2(b Vec2, bias float64) Vec2 {
	return Vec2{
		a[0][0]*b.X + a[0][1]*b.Y + a[0][2]*bias,
		a[1][0]*b.X + a[1][1]*b.Y + a[1][2]*bias,
	}
}

func (a Mat33) TimesOri2(b Ori2) Ori2 {
	return Ori2{
		a[0][0]*b.X + a[0][1]*b.Y + a[0][2]*b.Theta,
		a[1][0]*b.X + a[1][1]*b.Y + a[1][2]*b.Theta,
		a[2][0]*b.X + a[2][1]*b.Y + a[2][2]*b.Theta,
	}
}

func (m Mat33) Transposed() Mat33 {
	return Mat33{
		{m[0][0], m[1][0], m[2][0]},
		{m[0][1], m[1][1], m[2][1]},
		{m[0][2], m[1][2], m[2][2]},
	}
}

func (m Mat33) Determinant() float64 {
	return m[0][0]*(m[1][1]*m[2][2]-m[1][2]*m[2][1]) -
		m[0][1]*(m[1][0]*m[2][2]-m[1][2]*m[2][0]) +
		m[0][2]*(m[1][0]*m[2][1]-m[1][1]*m[2][0])
}

func (m Mat33) Inverse() Mat33 {
	/* 1. calculate determinant */
	det := m.Determinant()

	/* 2. find transpose of matrix */
	m = m.Transposed()

	/* 3. create matrix of cofactors */
	var cof Mat33
	sign := 1.0
	skip := [3][2]int{{1, 2}, {0, 2}, {0, 1}}
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			r1, r2 := skip[r][0], skip[r][1]
			c1, c2 := skip[c][0], skip[c][1]
			cof[r][c] = sign * (m[r1][c1]*m[r2][c2] - m[r2][c1]*m[r1][c2])
			sign = -sign // alternate signs
		}
	}

	/* 4. Divide each term by determinant */
	return cof.Scaled(1.0 / det)
}
