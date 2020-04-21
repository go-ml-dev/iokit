package fu

import (
	"reflect"
)

func Option(t interface{}, o interface{}) reflect.Value {
	xs := reflect.ValueOf(o)
	tv := reflect.ValueOf(t)
	for i := 0; i < xs.Len(); i++ {
		x := xs.Index(i)
		if x.Kind() == reflect.Interface {
			x = x.Elem()
		}
		if x.Type() == tv.Type() {
			return x
		}
	}
	return tv
}

func Mini(a int, b ...int) int {
	q := a
	for _, x := range b {
		if x < q {
			q = x
		}
	}
	return q
}

/*
Maxi returns maximal int value
*/
func Maxi(a int, b ...int) int {
	q := a
	for _, x := range b {
		if x > q {
			q = x
		}
	}
	return q
}

func Ife(expr bool, x interface{}, y interface{}) interface{} {
	if expr {
		return x
	}
	return y
}

func Ifei(expr bool, x int, y int) int {
	if expr {
		return x
	}
	return y
}

func Ifes(expr bool, x string, y string) string {
	if expr {
		return x
	}
	return y
}

func Fvs(a ...interface{}) string {
	return a[0].(string)
}
