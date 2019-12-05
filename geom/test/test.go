package geomTest

import (
	"math"
)

var (
	nan  = math.NaN()
	pInf = math.Inf(0)
	nInf = math.Inf(-1)
)

func floatIdentical(a, b float64) bool {
	return math.IsNaN(a) && math.IsNaN(b) ||
		math.IsInf(a, -1) && math.IsInf(b, -1) ||
		math.IsInf(a, 1) && math.IsInf(b, 1) ||
		math.Abs(a-b) < 0.0000001
}
