package table

import (
	"reflect"
)

type T []interface{}

func Append(a, b T) T {
	if len(a) != len(b) {
		panic("Append(a, b, T) T: table mismatch")
	}
	ret := make(T, len(a))

	for row := range ret {
		va := reflect.ValueOf(a[row])
		vb := reflect.ValueOf(b[row])

		if vb.Kind() == reflect.Slice {
			ret[row] = reflect.AppendSlice(va, vb).Interface()
		} else {
			ret[row] = reflect.Append(va, vb).Interface()
		}
	}

	return ret
}

func Filter(t T, f func(column T) bool) T {
	ret := t.Slice(0, 0)
	col := make(T, len(t))

	for i := 0; i < t.Len(); i++ {
		for row := range t {
			v := reflect.ValueOf(t[row])
			col[row] = v.Index(i).Interface()
		}

		if f(col) {
			for row := range ret {
				vRet := reflect.ValueOf(ret[row])
				vCol := reflect.ValueOf(col[row])
				ret[row] = reflect.Append(vRet, vCol).Interface()
			}
		}
	}
	return ret
}

func (t T) Slice(i, j int) T {
	ret := make(T, len(t))
	for row := range t {
		v := reflect.ValueOf(t[row])
		ret[row] = v.Slice(i, j).Interface()
	}
	return ret
}

func (t T) Swap(i, j int) {
	for row := range t {
		reflect.Swapper(t[row])(i, j)
	}
}

func (t T) Len() int {
	if len(t) == 0 {
		return 0
	}

	return reflect.ValueOf(t[0]).Len()
}
