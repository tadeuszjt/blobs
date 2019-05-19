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

func testTableIdentical(t *testing.T, expected, actual T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func testTableCapacity(t *testing.T, expected int, table T) {
	if len(table) == 0 {
	} else {
		for row := range table {
			v := reflect.ValueOf(table[row])
			actual := v.Cap()
			if expected != actual {
				t.Errorf("table: %v, row: %v, expected cap: %v, actual cap: %v",
					table, row, expected, actual)
			}
		}
	}
}

func TestAppend(t *testing.T) {
	cases := []struct {
		a, b, result T
		resultCap    int
	}{
		{
			T{},
			T{},
			T{}, 0,
		},
		{
			T{[]int{}},
			T{[]int{}},
			T{[]int{}}, 0,
		},
		{
			T{[]int{}},
			T{[]int{0}},
			T{[]int{0}}, 1,
		},
		{
			T{[]int{1}},
			T{[]int{2}},
			T{[]int{1, 2}}, 2,
		},
		{
			T{[]int{1, 2}},
			T{[]int{3}},
			T{[]int{1, 2, 3}}, 4,
		},
		{
			T{[]int{1, 2, 3}},
			T{[]int{4}},
			T{[]int{1, 2, 3, 4}}, 6,
		},
		{
			T{make([]int, 4, 16)},
			T{[]int{5}},
			T{[]int{0, 0, 0, 0, 5}}, 16,
		},
		{
			T{[]int{}, []float64{}},
			T{[]int{}, []float64{}},
			T{[]int{}, []float64{}}, 0,
		},
		{
			T{[]int{1}, []float64{1}},
			T{[]int{2}, []float64{2}},
			T{[]int{1, 2}, []float64{1, 2}}, 2,
		},
		{
			T{[]int{1, 2, 3, 4}},
			T{5},
			T{[]int{1, 2, 3, 4, 5}}, 8,
		},
		{
			T{[]int{1, 2, 3, 4, 5, 6}, []float64{1, 2, 3, 4, 5, 6}},
			T{7, 7.0},
			T{[]int{1, 2, 3, 4, 5, 6, 7}, []float64{1, 2, 3, 4, 5, 6, 7}}, 12,
		},
	}

	for _, c := range cases {
		actual := Append(c.a, c.b)
		testTableIdentical(t, c.result, actual)
		testTableCapacity(t, c.resultCap, actual)
	}
}

func TestFilter(t *testing.T) {
	cases := []struct {
		function      func(column T) bool
		table, result T
	}{
		{
			func(T) bool { return false },
			T{[]int{1, 2, 3, 4}},
			T{[]int{}},
		},
		{
			func(T) bool { return true },
			T{[]int{1, 2, 3, 4}},
			T{[]int{1, 2, 3, 4}},
		},
		{
			func(col T) bool { return col[0].(int) > 3 },
			T{[]int{1, 2, 3, 4, 5, 6}},
			T{[]int{4, 5, 6}},
		},
		{
			func(col T) bool { return col[0].(int)%2 == 0 },
			T{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
			T{[]int{2, 4, 6, 8, 10, 12, 14, 16}},
		},
	}

	for _, c := range cases {
		actual := Filter(c.table, c.function)
		testTableIdentical(t, c.result, actual)
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
	}

	for _, c := range cases {
		expected := c.result
		actual := c.table.Slice(c.i, c.j)
		testTableIdentical(t, expected, actual)
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
	}

	for _, c := range cases {
		expected := c.length
		actual := c.table.Len()
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestSwap(t *testing.T) {
	cases := []struct {
		i, j          int
		table, result T
	}{
		{0, 0, T{[]int{1}}, T{[]int{1}}},
		{0, 1, T{[]int{1, 2}}, T{[]int{2, 1}}},
		{0, 1, T{[]int{1, 2, 3}}, T{[]int{2, 1, 3}}},
		{0, 2, T{[]int{1, 2, 3}}, T{[]int{3, 2, 1}}},
		{1, 2, T{[]int{1, 2, 3}}, T{[]int{1, 3, 2}}},
		{2, 2, T{[]int{1, 2, 3}}, T{[]int{1, 2, 3}}},
		{
			1, 2,
			T{
				[]int{1, 2, 3},
				[]rune{'1', '2', '3'},
			},
			T{
				[]int{1, 3, 2},
				[]rune{'1', '3', '2'},
			},
		},
	}

	for _, c := range cases {
		expected := c.result
		c.table.Swap(c.i, c.j)
		actual := c.table
		testTableIdentical(t, expected, actual)
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
