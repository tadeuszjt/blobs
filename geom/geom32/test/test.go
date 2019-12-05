package geomTest

import (
	"math"
)

var (
	nan  = float32(math.NaN())
	pInf = float32(math.Inf(0))
	nInf = float32(math.Inf(-1))
)

func floatIdentical(a32, b32 float32) bool {
	a := float64(a32)
	b := float64(b32)
	return math.IsNaN(a) && math.IsNaN(b) ||
		math.IsInf(a, -1) && math.IsInf(b, -1) ||
		math.IsInf(a, 1) && math.IsInf(b, 1) ||
		math.Abs(a-b) < 0.000001
}
