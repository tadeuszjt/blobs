package geomTest

import (
	"github.com/tadeuszjt/geom"
	"testing"
)

func mat3Identical(a, b geom.Mat3) bool {
	for i := range a {
		if !floatIdentical(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestMat3Identity(t *testing.T) {
	expected := geom.Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
	actual := geom.Mat3Identity()
	if !mat3Identical(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestMat3TimesVec2(t *testing.T) {
	for _, c := range []struct {
		mat    geom.Mat3
		vec    geom.Vec2
		bias   float64
		result geom.Vec3
	}{
		{
			geom.Mat3Identity(),
			geom.Vec2{1, 1},
			1,
			geom.Vec3{1, 1, 1},
		},
		{
			geom.Mat3{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			geom.Vec2{10, 11},
			1,
			geom.Vec3{35, 101, 167},
		},
		{
			geom.Mat3{
				-3, pInf, 2.2,
				0, -38, 7,
				nan, 8, -0.1,
			},
			geom.Vec2{-1, -2},
			-3,
			geom.Vec3{nInf, 55, nan},
		},
		{
			geom.Mat3{
				pInf, 0, 0,
				nInf, 0, 0,
				0.001, -0.002, 0.003,
			},
			geom.Vec2{0, 1},
			2,
			geom.Vec3{nan, nan, 0.004},
		},
	} {
		expected := c.result
		actual := c.mat.TimesVec2(c.vec, c.bias)
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestMat3Camera2D(t *testing.T) {
	camera := geom.Rect{
		Min: geom.Vec2{10, 16},
		Max: geom.Vec2{50, 32},
	}
	display := geom.Rect{
		Min: geom.Vec2{-1, -2},
		Max: geom.Vec2{3, 4},
	}
	mat := geom.Mat3Camera2D(camera, display)

	cases := []struct {
		point, result geom.Vec2
	}{
		{geom.Vec2{30, 24}, geom.Vec2{1, 1}},
		{geom.Vec2{10, 16}, geom.Vec2{-1, -2}},
		{geom.Vec2{50, 16}, geom.Vec2{3, -2}},
		{geom.Vec2{50, 32}, geom.Vec2{3, 4}},
		{geom.Vec2{10, 32}, geom.Vec2{-1, 4}},
	}

	for _, c := range cases {
		actual := mat.TimesVec2(c.point, 1).Vec2()
		expected := c.result
		if !vec2Identical(expected, actual) {
			t.Errorf("point: %v: expected: %v, got: %v", c.point, expected, actual)
		}
	}

}
