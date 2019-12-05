package geom

type Mat44 [4][4]float64

func Mat44Identity() Mat44 {
	return Mat44{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func (a Mat44) Times(b Mat44) (ret Mat44) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			ret[r][c] = a[r][0]*b[0][c] + a[r][1]*b[1][c] + a[r][2]*b[2][c] + a[r][3]*b[3][c]
		}
	}
	return
}

func (m Mat44) TimesVec3(v Vec3, bias float64) Vec3 {
	return Vec3{
		m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3]*bias,
		m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3]*bias,
		m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3]*bias,
	}
}

func (m Mat44) Transposed() Mat44 {
	return Mat44{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
		{m[0][3], m[1][3], m[2][3], m[3][3]},
	}
}
