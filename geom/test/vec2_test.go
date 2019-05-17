package geomTest

import (
	"github.com/tadeuszjt/geom"
	"math"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func vec2Identical(a, b geom.Vec2) bool {
	return floatIdentical(a.X, b.X) && floatIdentical(a.Y, b.Y)
}

func TestZero(t *testing.T) {
	expected := geom.Vec2{0, 0}
	actual := geom.Vec2{}
	if !vec2Identical(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestPlus(t *testing.T) {
	cases := []struct {
		A, B, result geom.Vec2
	}{
		{geom.Vec2{0, 0}, geom.Vec2{0, 0}, geom.Vec2{0, 0}},
		{geom.Vec2{1, 2}, geom.Vec2{0, 0}, geom.Vec2{1, 2}},
		{geom.Vec2{-1, 2}, geom.Vec2{3, 4}, geom.Vec2{2, 6}},
		{geom.Vec2{-1, -2}, geom.Vec2{-3, -4}, geom.Vec2{-4, -6}},
		{geom.Vec2{nan, -2}, geom.Vec2{-3, -4}, geom.Vec2{nan, -6}},
		{geom.Vec2{pInf, -2}, geom.Vec2{-3, -4}, geom.Vec2{pInf, -6}},
		{geom.Vec2{nInf, -2}, geom.Vec2{-3, -4}, geom.Vec2{nInf, -6}},
		{geom.Vec2{pInf, -2}, geom.Vec2{nInf, -4}, geom.Vec2{nan, -6}},
	}
	for _, c := range cases {
		expected := c.result
		actual := c.A.Plus(c.B)
		if !vec2Identical(expected, actual) {
			t.Errorf("%v.Plus(%v): expected: %v, got: %v", c.A, c.B, expected, actual)
		}
	}
}

func TestMinus(t *testing.T) {
	cases := []struct {
		a, b, result geom.Vec2
	}{
		{geom.Vec2{0, 0}, geom.Vec2{0, 0}, geom.Vec2{0, 0}},
		{geom.Vec2{1, 2}, geom.Vec2{2, 6}, geom.Vec2{-1, -4}},
		{geom.Vec2{8, 9}, geom.Vec2{3, 4}, geom.Vec2{5, 5}},
		{geom.Vec2{-1, -2}, geom.Vec2{-3, -4}, geom.Vec2{2, 2}},
		{geom.Vec2{pInf, -2}, geom.Vec2{-3, -4}, geom.Vec2{pInf, 2}},
		{geom.Vec2{nInf, -2}, geom.Vec2{-3, -4}, geom.Vec2{nInf, 2}},
		{geom.Vec2{pInf, -2}, geom.Vec2{nInf, -4}, geom.Vec2{pInf, 2}},
		{geom.Vec2{pInf, -2}, geom.Vec2{pInf, -4}, geom.Vec2{nan, 2}},
	}
	for _, c := range cases {
		expected := c.result
		actual := c.a.Minus(c.b)
		if !vec2Identical(expected, actual) {
			t.Errorf("%v.Minus(%v): expected: %v, got: %v", c.a, c.b, expected, actual)
		}
	}
}

func TestScaled(t *testing.T) {
	cases := []struct {
		scalar    float64
		v, result geom.Vec2
	}{
		{0, geom.Vec2{0, 0}, geom.Vec2{0, 0}},
		{0, geom.Vec2{1, 2}, geom.Vec2{0, 0}},
		{2, geom.Vec2{1, 2}, geom.Vec2{2, 4}},
		{-2, geom.Vec2{1, 2}, geom.Vec2{-2, -4}},
		{2, geom.Vec2{-9, 2}, geom.Vec2{-18, 4}},
		{0.001, geom.Vec2{-9, 2}, geom.Vec2{-0.009, 0.002}},
		{0.001, geom.Vec2{pInf, 0}, geom.Vec2{pInf, 0}},
		{0.001, geom.Vec2{nInf, 0}, geom.Vec2{nInf, 0}},
		{0.001, geom.Vec2{nan, 0}, geom.Vec2{nan, 0}},
		{0, geom.Vec2{nInf, 0}, geom.Vec2{nan, 0}},
	}
	for _, c := range cases {
		expected := c.result
		actual := c.v.Scaled(c.scalar)
		if !vec2Identical(expected, actual) {
			t.Errorf("%v.Scaled(%v): expected: %v, got: %v", c.v, c.scalar, expected, actual)
		}
	}
}

func TestVec2Rand(t *testing.T) {
	cases := []geom.Rect{
		geom.Rect{},
		geom.Rect{geom.Vec2{10, 10}, geom.Vec2{20, 20}},
		geom.Rect{geom.Vec2{-10, 20}, geom.Vec2{50, 30}},
		geom.Rect{geom.Vec2{-0.001, 0}, geom.Vec2{0.001, 0.0001}},
		geom.Rect{geom.Vec2{0, 0}, geom.Vec2{10000, 100000}},
		geom.Rect{geom.Vec2{0, 0}, geom.Vec2{1, pInf}},
		geom.Rect{geom.Vec2{nInf, 0}, geom.Vec2{0, 0}},
		geom.Rect{geom.Vec2{nInf, 0}, geom.Vec2{pInf, 0}},
		geom.Rect{geom.Vec2{nan, 0}, geom.Vec2{1, 2}},
	}

	for _, rect := range cases {
		for i := 0; i < 4; i++ {
			vec := geom.Vec2Rand(rect)

			if vec.X < rect.Min.X ||
				vec.X > rect.Max.X ||
				vec.Y < rect.Min.Y ||
				vec.Y > rect.Max.Y {
				t.Errorf("%v does not contain %v", rect, vec)
			}
		}
	}
}

func TestPlusEquals(t *testing.T) {
	cases := []struct {
		A, B, result geom.Vec2
	}{
		{geom.Vec2{}, geom.Vec2{}, geom.Vec2{}},
		{geom.Vec2{}, geom.Vec2{1, 2}, geom.Vec2{1, 2}},
		{geom.Vec2{}, geom.Vec2{-1, -2}, geom.Vec2{-1, -2}},
		{geom.Vec2{0.002, -9.32}, geom.Vec2{0, 43.2}, geom.Vec2{0.002, 33.88}},
		{geom.Vec2{nan, pInf}, geom.Vec2{0, 43.2}, geom.Vec2{nan, pInf}},
	}

	for _, c := range cases {
		v := c.A
		v.PlusEquals(c.B)
		expected := c.result
		actual := v
		if !vec2Identical(expected, actual) {
			t.Errorf("%v.PlusEquals(%v): expected: %v, got: %v", c.A, c.B, expected, actual)
		}
	}
}

func TestRandNormal(t *testing.T) {
	var sections [4]bool

	for i := 0; i < 100; i++ {
		v := geom.Vec2RandNormal()

		length := v.X*v.X + v.Y*v.Y
		if !floatIdentical(length, 1) {
			t.Errorf("%v: expected length 1, got %v", v, length)
		}

		theta := math.Atan2(v.Y, v.X)
		switch {
		case theta > math.Pi*-1.0 && theta < math.Pi*-0.5:
			sections[0] = true
		case theta > math.Pi*-0.5 && theta < math.Pi*0:
			sections[1] = true
		case theta > math.Pi*0 && theta < math.Pi*0.5:
			sections[2] = true
		case theta > math.Pi*0.5 && theta < math.Pi*1:
			sections[3] = true
		}
	}

	for i := range sections {
		if sections[i] == false {
			t.Errorf("%d quarter of circle not covered", i)
		}
	}
}

func TestLen2(t *testing.T) {
	// a2 = b2 + c2
	cases := []struct {
		vec  geom.Vec2
		len2 float64
	}{
		{geom.Vec2{}, 0},
		{geom.Vec2{1, 0}, 1},
		{geom.Vec2{2, 0}, 4},
		{geom.Vec2{0, 2}, 4},
		{geom.Vec2{2, 2}, 8},
		{geom.Vec2{-3, 4}, 25},
		{geom.Vec2{3, 4}, 25},
	}

	for _, c := range cases {
		expected := c.len2
		actual := c.vec.Len2()
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, got %v", expected, actual)
		}
	}
}
