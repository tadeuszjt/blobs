package table

import (
	"reflect"
	"testing"
)

func testPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("funcion didn't panic")
		}
	}()
	f()
}

func tableIdentical(a, b T) bool {
	return reflect.DeepEqual(a, b)
}

func TestTIdentical(t *testing.T) {
	cases := []struct {
		a, b   T
		result bool
	}{
		{T{}, T{}, true},
		{T{}, T{[]float64{}}, false},
		{T{[]float64{}}, T{[]float64{}}, true},
		{T{[]float64{}}, T{[]float64{1}}, false},
		{T{[]float64{1}}, T{[]float64{1}}, true},
		{T{[]float64{1}, []int{}}, T{[]float64{1}}, false},
		{T{[]float64{1}, []int{}}, T{[]float64{1}, []int{}}, true},
		{T{[]float64{1}, []int{2}}, T{[]float64{1}, []int{2}}, true},
		{T{[]float64{1}, []int{2, 3}}, T{[]float64{1}, []int{2, 2}}, false},
		{T{[]float64{1}, []int{2, 3}}, T{[]float64{1}, []int{2, 3}}, true},
	}

	for _, c := range cases {
		expected := c.result
		actual := tableIdentical(c.a, c.b)
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestAppend(t *testing.T) {
	cases := []struct {
		a, b, result T
	}{
		{
			T{},
			T{},
			T{},
		},
		{
			T{[]float64{}},
			T{[]float64{}},
			T{[]float64{}},
		},
		{
			T{[]float64{1, 2, 3}},
			T{[]float64{}},
			T{[]float64{1, 2, 3}},
		},
		{
			T{[]float64{}},
			T{[]float64{4, 5, 6}},
			T{[]float64{4, 5, 6}},
		},
		{
			T{[]float64{1, 2}},
			T{[]float64{4, 5}},
			T{[]float64{1, 2, 4, 5}},
		},
		{
			T{[]float64{1, 2}, []int{3, 4}},
			T{[]float64{4, 5}, []int{8, 9}},
			T{[]float64{1, 2, 4, 5}, []int{3, 4, 8, 9}},
		},
		{
			T{[]uint{0, 1, 2, 3}, []rune{'a', 'b', 'c', 'd'}},
			T{[]uint{4}, []rune{'e'}},
			T{[]uint{0, 1, 2, 3, 4}, []rune{'a', 'b', 'c', 'd', 'e'}},
		},
		{
			T{[]uint{0, 1, 2, 3}},
			T{uint(4)},
			T{[]uint{0, 1, 2, 3, 4}},
		},
		{
			T{'j'},
			T{[]rune{'e', 'n', 'd'}},
			T{[]rune{'j', 'e', 'n', 'd'}},
		},
		{
			T{'j'},
			T{'e'},
			T{[]rune{'j', 'e'}},
		},
		{
			T{'j', 3, 0.2},
			T{[]rune{'a', 'b', 'c'}, []int{1, 2, 3}, []float64{0.1, 0, -0.1}},
			T{[]rune{'j', 'a', 'b', 'c'}, []int{3, 1, 2, 3}, []float64{0.2, 0.1, 0, -0.1}},
		},
		{
			T{[]rune{'a', 'b', 'c'}, []int{1, 2, 3}, []float64{0.1, 0, -0.1}},
			T{'j', 3, 0.2},
			T{[]rune{'a', 'b', 'c', 'j'}, []int{1, 2, 3, 3}, []float64{0.1, 0, -0.1, 0.2}},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := Append(c.a, c.b)
		if !tableIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}


	for _, f := range []func() {
		func() { Append(T{}, T{[]float64{}}) },
		func() { Append(T{[]float64{}}, T{}) },
		func() { Append(T{[]float64{}}, T{[]int{}}) },
		func() { Append(T{[]float64{}}, T{[]float64{}, []int{}}) },
	}{
		testPanic(t, f)
	}
}

func TestDelete(t *testing.T) {
	cases := []struct {
		index         int
		table, result T
	}{
		{
			0,
			T{[]int{1, 2, 3}},
			T{[]int{2, 3}},
		},
		{
			1,
			T{[]int{1, 2, 3}},
			T{[]int{1, 3}},
		},
		{
			2,
			T{[]int{1, 2, 3}},
			T{[]int{1, 2}},
		},
		{
			1,
			T{[]int{1, 2, 3}, []rune{'a', 'b', 'c'}},
			T{[]int{1, 3}, []rune{'a', 'c'}},
		},
		{
			0,
			T{[]int{1}, []rune{'a'}, []float64{0.2}},
			T{[]int{}, []rune{}, []float64{}},
		},
		{
			0,
			T{0.2, 'b', 3},
			T{[]float64{}, []rune{}, []int{}},
		},
	}

	for _, c := range cases {
		expected := c.result
		c.table.Delete(c.index)
		actual := c.table
		if !tableIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestDeleteUnordered(t *testing.T) {
	cases := []struct {
		index         int
		table, result T
	}{
		{
			0,
			T{[]int{1, 2, 3}},
			T{[]int{3, 2}},
		},
		{
			1,
			T{[]int{1, 2, 3}},
			T{[]int{1, 3}},
		},
		{
			2,
			T{[]int{1, 2, 3}},
			T{[]int{1, 2}},
		},
		{
			0,
			T{[]int{1, 2, 3}, []rune{'a', 'b', 'c'}},
			T{[]int{3, 2}, []rune{'c', 'b'}},
		},

		{
			1,
			T{[]int{1, 2, 3}, []rune{'a', 'b', 'c'}},
			T{[]int{1, 3}, []rune{'a', 'c'}},
		},
		{
			0,
			T{[]int{1}, []rune{'a'}, []float64{0.2}},
			T{[]int{}, []rune{}, []float64{}},
		},
		{
			0,
			T{0.2, 'b', 3},
			T{[]float64{}, []rune{}, []int{}},
		},
	}

	for _, c := range cases {
		expected := c.result
		c.table.DeleteUnordered(c.index)
		actual := c.table
		if !tableIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestSlice(t *testing.T) {
	cases := []struct {
		i, j          int
		table, result T
	}{
		{0, 0, T{}, T{}},
		{0, 0, T{[]int{1, 2, 3}}, T{[]int{}}},
		{0, 1, T{[]int{1, 2, 3}}, T{[]int{1}}},
		{0, 2, T{[]int{1, 2, 3}}, T{[]int{1, 2}}},
		{0, 3, T{[]int{1, 2, 3}}, T{[]int{1, 2, 3}}},
		{1, 1, T{[]int{1, 2, 3}}, T{[]int{}}},
		{1, 2, T{[]int{1, 2, 3}}, T{[]int{2}}},
		{1, 3, T{[]int{1, 2, 3}}, T{[]int{2, 3}}},
		{2, 2, T{[]int{1, 2, 3}}, T{[]int{}}},
		{2, 3, T{[]int{1, 2, 3}}, T{[]int{3}}},
		{
			1, 4,
			T{
				[]uint{1, 2, 3, 4, 5},
				[]rune{'1', '2', '3', '4', '5'},
			},
			T{
				[]uint{2, 3, 4},
				[]rune{'2', '3', '4'},
			},
		},
		{
			0, 0,
			T{'a', 0.2, 3},
			T{[]rune{}, []float64{}, []int{}},
		},
		{
			0, 1,
			T{'a', 0.2, 3},
			T{[]rune{'a'}, []float64{0.2}, []int{3}},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.table.Slice(c.i, c.j)
		if !tableIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestLen(t *testing.T) {
	cases := []struct {
		length int
		table  T
	}{
		{0, T{}},
		{0, T{
			[]float64{},
		}},
		{0, T{
			[]float64{},
			[]rune{},
		}},
		{1, T{
			[]float64{1},
			[]rune{'1'},
		}},
		{3, T{
			[]float64{1, 2, 3},
			[]rune{'1', '2', '3'},
		}},
		{100, T{
			make([]int, 100),
		}},
		{1, T{
			0.1,
			2,
			'3',
		}},
	}

	for _, c := range cases {
		expected := c.length
		actual := c.table.Len()
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestFilter(t *testing.T) {
	cases := []struct {
		function      func(column T) bool
		table, result T
	}{
		{
			func(T) bool { return false },
			T{[]int{1, 2 ,3}},
			T{[]int{}},
		},
		{
			func(T) bool { return true },
			T{[]int{1, 2 ,3}},
			T{[]int{1, 2, 3}},
		},
		{
			func(col T) bool { return col[0].(int) > 3 },
			T{[]int{1, 2, 3, 4, 5, 6}},
			T{[]int{4, 5, 6}},
		},
		{
			func(col T) bool { return col[0].(int) % 2 == 0 },
			T{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
			T{[]int{2, 4, 6, 8, 10, 12, 14, 16}},
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := Filter(c.table, c.function)
		if !tableIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}


func BenchmarkFilter(b *testing.B) {
	t := T{
		make([]int, 10000),
		make([]float64, 10000),
		make([]rune, 10000),
	}

	for n := 0; n < b.N; n++ {
		Filter(t, func(column T) bool { return true })
	}
}

func BenchmarkAppend(b *testing.B) {
	t := T{
		[]int{},
		[]float64{},
		[]rune{},
	}

	for n := 0; n < b.N; n++ {
		t = Append(t, T{
			n,
			float64(n),
			rune(n),
		})

		t = Append(t, T{
			[]int{1, 2, 3},
			[]float64{1, 2, 3},
			[]rune{'a', 'b', 'c'},
		})
	}
}
