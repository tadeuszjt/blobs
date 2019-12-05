package geomTest

import (
	"github.com/tadeuszjt/geom/geom32"
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

func TestVec3Dot(t *testing.T) {
	cases := []struct {
		a, b   geom.Vec3
		result float32
	}{
		{geom.Vec3{}, geom.Vec3{}, 0},
		{geom.Vec3{1, 2, 3}, geom.Vec3{4, 5, 6}, 32},
		{geom.Vec3{0, 0, 0}, geom.Vec3{4, 5, 6}, 0},
		{geom.Vec3{-1, -2, -3}, geom.Vec3{4, 5, 6}, -32},
		{geom.Vec3{-1, 2, -3}, geom.Vec3{4, 5, 6}, -12},
		{geom.Vec3{-1, nan, -3}, geom.Vec3{4, 5, 6}, nan},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.a.Dot(c.b)
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestVec3Times(t *testing.T) {
	cases := []struct {
		a, b, result geom.Vec3
	}{
		{geom.Vec3{}, geom.Vec3{}, geom.Vec3{}},
		{geom.Vec3{1, 2, 3}, geom.Vec3{4, 5, 6}, geom.Vec3{4, 10, 18}},
		{geom.Vec3{-1, 2, -3}, geom.Vec3{4, -5, 6}, geom.Vec3{-4, -10, -18}},
		{geom.Vec3{nan, pInf, nInf}, geom.Vec3{-4, -5, -6}, geom.Vec3{nan, nInf, pInf}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.a.Times(c.b)
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestVec3ScaledBy(t *testing.T) {
	cases := []struct {
		scalar    float32
		v, result geom.Vec3
	}{
		{0, geom.Vec3{}, geom.Vec3{}},
		{0, geom.Vec3{1, 2, 3}, geom.Vec3{0, 0, 0}},
		{1, geom.Vec3{1, 2, 3}, geom.Vec3{1, 2, 3}},
		{-1, geom.Vec3{1, 2, 3}, geom.Vec3{-1, -2, -3}},
		{100, geom.Vec3{1, 2, 3}, geom.Vec3{100, 200, 300}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.v.ScaledBy(c.scalar)
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestVec3Ori2(t *testing.T) {
	cases := []struct {
		v geom.Vec3
		o geom.Ori2
	}{
		{geom.Vec3{}, geom.Ori2{}},
		{geom.Vec3{1, 2, 3}, geom.Ori2{1, 2, 3}},
		{geom.Vec3{-1, -2, -3}, geom.Ori2{-1, -2, -3}},
		{geom.Vec3{nan, pInf, nInf}, geom.Ori2{nan, pInf, nInf}},
	}

	for _, c := range cases {
		expected := c.o
		actual := c.v.Ori2()
		if !ori2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}
