package iokit

import (
	"reflect"
)

func option(t interface{}, o interface{}) reflect.Value {
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

func mini(a int, b ...int) int {
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
func maxi(a int, b ...int) int {
	q := a
	for _, x := range b {
		if x > q {
			q = x
		}
	}
	return q
}

func ife(expr bool, x interface{}, y interface{}) interface{} {
	if expr {
		return x
	}
	return y
}

func ifei(expr bool, x int, y int) int {
	if expr {
		return x
	}
	return y
}

func ifes(expr bool, x string, y string) string {
	if expr {
		return x
	}
	return y
}
