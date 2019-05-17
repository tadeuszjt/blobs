package geomTest

import (
	"github.com/tadeuszjt/geom"
	"testing"
)

func vec3Identical(a, b geom.Vec3) bool {
	return floatIdentical(a.X, b.X) &&
		floatIdentical(a.Y, b.Y) &&
		floatIdentical(a.Z, b.Z)
}

func TestVec3Vec2(t *testing.T) {
	for _, c := range []struct {
		geom.Vec3
		geom.Vec2
	}{
		{geom.Vec3{0, 0, 0}, geom.Vec2{0, 0}},
		{geom.Vec3{1, 2, 3}, geom.Vec2{1, 2}},
		{geom.Vec3{-1, -2, -3}, geom.Vec2{-1, -2}},
		{geom.Vec3{nan, nInf, pInf}, geom.Vec2{nan, nInf}},
	} {
		expected := c.Vec2
		actual := c.Vec3.Vec2()
		if !vec2Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}
