package generic

import (
	"reflect"
)

func isIterable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String, reflect.Struct:
		return true
	}
	return false
}

func isString(t reflect.Type) bool {
	return t.Kind() == reflect.String
}

func isAny(v interface{}, f func(reflect.Type) bool) bool {
	return v != nil && f(reflect.TypeOf(v))
}

func IsIterable[V any]() bool {
	return isIterable(reflect.TypeOf((*V)(nil)).Elem())
}

func IsString[V any]() bool {
	return isString(reflect.TypeOf((*V)(nil)).Elem())
}

func IsAnyIterable(v interface{}) bool {
	return isAny(v, isIterable)
}

func IsAnyString(v interface{}) bool {
	return isAny(v, isString)
}
