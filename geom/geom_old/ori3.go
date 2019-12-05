package geom

type Ori3 struct {
	Vec3
	Quarternion
}

func (o Ori3) Mat44Transform() Mat44 {
	r := o.Quarternion.Mat33Rotation()
	return Mat44{
		{r[0][0], r[0][1], r[0][2], o.Vec3.X},
		{r[1][0], r[1][1], r[1][2], o.Vec3.Y},
		{r[2][0], r[2][1], r[2][2], o.Vec3.Z},
		{0, 0, 0, 1},
	}
}

func (o Ori3) Mat44InvTransform() Mat44 {
	r := o.Quarternion.Mat33Rotation().Transposed()
	p := r.TimesVec3(o.Vec3)
	return Mat44{
		{r[0][0], r[0][1], r[0][2], -p.X},
		{r[1][0], r[1][1], r[1][2], -p.Y},
		{r[2][0], r[2][1], r[2][2], -p.Z},
		{0, 0, 0, 1},
	}
}
