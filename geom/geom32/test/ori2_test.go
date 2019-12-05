package geomTest

import (
	"github.com/tadeuszjt/geom/geom32"
	"math"
	"testing"
)

func ori2Identical(a, b geom.Ori2) bool {
	return floatIdentical(a.X, b.X) &&
		floatIdentical(a.Y, b.Y) &&
		floatIdentical(a.Theta, b.Theta)
}

func TestOri2(t *testing.T) {
	expected := geom.Ori2{0, 0, 0}
	actual := geom.Ori2{}
	if !ori2Identical(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestVec2(t *testing.T) {
	cases := []struct {
		o geom.Ori2
		v geom.Vec2
	}{
		{geom.Ori2{}, geom.Vec2{}},
		{geom.Ori2{1, 2, 3}, geom.Vec2{1, 2}},
		{geom.Ori2{.1, .2, .3}, geom.Vec2{.1, .2}},
		{geom.Ori2{-1, -2, 3}, geom.Vec2{-1, -2}},
		{geom.Ori2{nan, pInf, nInf}, geom.Vec2{nan, pInf}},
	}

	for _, c := range cases {
		expected := c.v
		actual := c.o.Vec2()
		if !vec2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2PlusEquals(t *testing.T) {
	cases := []struct {
		v, b, result geom.Ori2
	}{
		{geom.Ori2{}, geom.Ori2{}, geom.Ori2{}},
		{geom.Ori2{0, 0, 0}, geom.Ori2{1, 2, 3}, geom.Ori2{1, 2, 3}},
		{geom.Ori2{0, 0, 0}, geom.Ori2{-1, -2, -3}, geom.Ori2{-1, -2, -3}},
		{geom.Ori2{1, 2, 3}, geom.Ori2{4, 5, 6}, geom.Ori2{5, 7, 9}},
		{geom.Ori2{1, 2, 3}, geom.Ori2{-4, -5, -6}, geom.Ori2{-3, -3, -3}},
		{geom.Ori2{nan, 2, 3}, geom.Ori2{4, 5, 6}, geom.Ori2{nan, 7, 9}},
	}

	for _, c := range cases {
		expected := c.result
		c.v.PlusEquals(c.b)
		actual := c.v
		if !ori2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2ScaledBy(t *testing.T) {
	cases := []struct {
		scalar    float32
		o, result geom.Ori2
	}{
		{0, geom.Ori2{}, geom.Ori2{}},
		{0, geom.Ori2{1, 2, 3}, geom.Ori2{0, 0, 0}},
		{1, geom.Ori2{1, 2, 3}, geom.Ori2{1, 2, 3}},
		{-2, geom.Ori2{1, 2, 3}, geom.Ori2{-2, -4, -6}},
		{nan, geom.Ori2{1, 2, 3}, geom.Ori2{nan, nan, nan}},
		{0.001, geom.Ori2{1, 2, 3}, geom.Ori2{0.001, 0.002, 0.003}},
		{pInf, geom.Ori2{1, -2, 3}, geom.Ori2{pInf, nInf, pInf}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.o.ScaledBy(c.scalar)
		if !ori2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2Mat3Transform(t *testing.T) {
	cases := []struct {
		o      geom.Ori2
		v      geom.Vec2
		result geom.Vec3
	}{
		{geom.Ori2{}, geom.Vec2{}, geom.Vec3{0, 0, 1}},
		{geom.Ori2{1, 2, 0}, geom.Vec2{0, 0}, geom.Vec3{1, 2, 1}},
		{geom.Ori2{1, 2, 0}, geom.Vec2{3, 4}, geom.Vec3{4, 6, 1}},
		{geom.Ori2{3, 4, math.Pi / 2}, geom.Vec2{1, 2}, geom.Vec3{1, 5, 1}},
		{geom.Ori2{3, 4, -math.Pi / 2}, geom.Vec2{1, 2}, geom.Vec3{5, 3, 1}},
		{geom.Ori2{-2, 8, math.Pi}, geom.Vec2{3, -2}, geom.Vec3{-5, 10, 1}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.o.Mat3Transform().TimesVec2(c.v, 1)
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2Vec2(t *testing.T) {
	cases := []struct {
		o      geom.Ori2
		result geom.Vec3
	}{
		{geom.Ori2{}, geom.Vec3{}},
		{geom.Ori2{1, 2, 3}, geom.Vec3{1, 2, 3}},
		{geom.Ori2{-1, -2, -3}, geom.Vec3{-1, -2, -3}},
		{geom.Ori2{0.001, 0.002, 0.003}, geom.Vec3{0.001, 0.002, 0.003}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.o.Vec3()
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2Times(t *testing.T) {
	cases := []struct {
		a, b, result geom.Ori2
	}{
		{geom.Ori2{}, geom.Ori2{}, geom.Ori2{}},
		{geom.Ori2{0, 0, 0}, geom.Ori2{1, 2, 3}, geom.Ori2{0, 0, 0}},
		{geom.Ori2{1, 0.2, 3}, geom.Ori2{0.4, 5, 0.6}, geom.Ori2{0.4, 1, 1.8}},
		{geom.Ori2{-1, -2, -3}, geom.Ori2{4, 5, 6}, geom.Ori2{-4, -10, -18}},
		{geom.Ori2{nan, pInf, nInf}, geom.Ori2{-4, -5, -6}, geom.Ori2{nan, nInf, pInf}},
		{geom.Ori2{nan, pInf, nInf}, geom.Ori2{0, 0, 0}, geom.Ori2{nan, nan, nan}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.a.Times(c.b)
		if !ori2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}
