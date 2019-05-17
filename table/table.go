package table

import (
	"reflect"
)

type T []interface{}

func valueOf(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice {
		return reflect.Append(reflect.MakeSlice(reflect.SliceOf(v.Type()), 0, 1), v)
	}
	return v
}

func Append(a, b T) T {
	if len(a) != len(b) {
		panic("tables don't match")
	}
	t := make(T, len(b))

	for row := range t {
		va := reflect.ValueOf(a[row])
		if va.Kind() != reflect.Slice {
			va = reflect.Append(
				reflect.MakeSlice(reflect.SliceOf(va.Type()), 0, 1),
				va,
			)
		}

		vb := reflect.ValueOf(b[row])
		if vb.Kind() == reflect.Slice {
			t[row] = reflect.AppendSlice(va, vb).Interface()
		} else {
			t[row] = reflect.Append(va, vb).Interface()
		}
	}

	return t
}

func Filter(t T, f func(column T) bool) T {
	ret := t.Slice(0, 0)
	col := make(T, len(t))

	for i := 0; i < t.Len(); i++ {
		for row := range t {
			vt := reflect.ValueOf(t[row])
			if vt.Kind() == reflect.Slice {
				col[row] = vt.Index(i).Interface()
			} else {
				col[row] = t[row]
			}
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

func (t T) Delete(column int) {
	for row := range t {
		v := valueOf(t[row])
		s1 := v.Slice(0, column)
		s2 := v.Slice(column+1, v.Len())
		t[row] = reflect.AppendSlice(s1, s2).Interface()
	}
}

func (t T) DeleteUnordered(column int) {
	for row := range t {
		v := valueOf(t[row])
		end := v.Len() - 1
		reflect.Swapper(v.Interface())(column, end)
		t[row] = v.Slice(0, end).Interface()
	}
}

func (t T) Slice(i, j int) T {
	r := make(T, 0, len(t))
	for row := range t {
		v := valueOf(t[row])
		r = append(r, v.Slice(i, j).Interface())
	}
	return r
}

func (t T) Len() int {
	if len(t) > 0 {
		v := reflect.ValueOf(t[0])
		if v.Kind() == reflect.Slice {
			return v.Len()
		} else {
			return 1
		}
	}
	return 0
}
