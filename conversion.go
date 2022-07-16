package generic

import "reflect"

// tryConvert tries to convert provided value to given type or interface
func tryConvert[T any](source reflect.Value) (result T, success bool) {
	itArg := reflect.TypeOf((*T)(nil)).Elem()
	switch itArg.Kind() {
	case reflect.Interface:
		if source.Type().Implements(itArg) {
			return source.Interface().(T), true
		}
	default:
		if source.CanConvert(itArg) {
			return source.Convert(itArg).Interface().(T), true
		}
	}
	return
}
