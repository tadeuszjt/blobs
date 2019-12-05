package geom

type Quarternion struct {
	W, X, Y, Z float64
}

func (q Quarternion) Mat33Rotation() Mat33 {
	qx2 := q.X * q.X
	qy2 := q.Y * q.Y
	qz2 := q.Z * q.Z

	return Mat33{
		{1 - 2*qy2 - 2*qz2, 2*q.X*q.Y - 2*q.Z*q.W, 2*q.X*q.Z + 2*q.Y*q.W},
		{2*q.X*q.Y + 2*q.Z*q.W, 1 - 2*qx2 - 2*qz2, 2*q.Y*q.Z - 2*q.X*q.W},
		{2*q.X*q.Z - 2*q.Y*q.W, 2*q.Y*q.Z + 2*q.X*q.W, 1 - 2*qx2 - 2*qy2},
	}
}
