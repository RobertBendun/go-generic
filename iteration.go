package generic

import "reflect"

// Each iterates by value over provided structure
func Each[Struct, T any](in Struct, iteration func(T)) {
	inV := reflect.ValueOf(in)

	switch inV.Type().Kind() {
	case reflect.Array, reflect.Slice, reflect.String:
		for i := 0; i < inV.Len(); i++ {
			if v, ok := tryConvert[T](inV.Index(i)); ok {
				iteration(v)
			}
		}
	case reflect.Struct:
		for i := 0; i < inV.NumField(); i++ {
			if v, ok := tryConvert[T](inV.Field(i)); ok {
				iteration(v)
			}
		}
	}
}

// EachWithKey iterates by value over provided structure
func EachWithKey[Struct any, K, V any](in Struct, iteration func(K, V)) {
	inV := reflect.ValueOf(in)
	keyT := reflect.TypeOf((*K)(nil)).Elem()

	switch inV.Type().Kind() {
	case reflect.Array, reflect.Slice, reflect.String:
		if reflect.TypeOf(0).AssignableTo(keyT) {
			for i := 0; i < inV.Len(); i++ {
				if v, ok := tryConvert[V](inV.Index(i)); ok {
					iteration(interface{}(i).(K), v)
				}
			}
		}
	case reflect.Struct:
		if reflect.TypeOf("").AssignableTo(keyT) {
			for i := 0; i < inV.NumField(); i++ {
				if v, ok := tryConvert[V](inV.Field(i)); ok {
					name := inV.Type().Field(i).Name
					iteration(interface{}(name).(K), v)
				}
			}
		}
	}

}

// Modify iterates by pointer over provided structure
func Modify[Struct, T any](in Struct, iteration func(T)) Struct {
	inV := reflect.Indirect(reflect.ValueOf(in))
	switch inV.Type().Kind() {
	case reflect.Array, reflect.Slice, reflect.String:
		for i := 0; i < inV.Len(); i++ {
			if v, ok := tryConvert[T](inV.Index(i).Addr()); ok {
				iteration(v)
			}
		}
	case reflect.Struct:
		for i := 0; i < inV.NumField(); i++ {
			if v, ok := tryConvert[T](inV.Field(i).Addr()); ok {
				iteration(v)
			}
		}
	}
	return in
}

// Map creates new value transformed by provided iteration on provided fields
func Map[Struct, T any](in Struct, iteration func(T) T) (out Struct) {
	out = in
	inV, outV := reflect.ValueOf(in), reflect.Indirect(reflect.ValueOf(&out))
	for i := 0; i < inV.NumField(); i++ {
		inF, outF := inV.Field(i), outV.Field(i)
		if v, ok := tryConvert[T](inF); ok {
			outF.Set(reflect.ValueOf(iteration(v)))
		}
	}
	return
}
