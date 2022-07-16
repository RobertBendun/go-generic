package generic

import (
	"reflect"
)

func All[Lhs, Rhs, T any](lhs Lhs, rhs Rhs, iteration func(lhs, rhs T) bool) bool {
	lhsV, rhsV := reflect.ValueOf(lhs), reflect.ValueOf(rhs)
	fieldsCount := lhsV.NumField()
	if fieldsCount < rhsV.NumField() {
		fieldsCount = rhsV.NumField()
	}

	for i := 0; i < fieldsCount; i++ {
		lhsF, lhsOk := tryConvert[T](lhsV.Field(i))
		rhsF, rhsOk := tryConvert[T](rhsV.Field(i))
		if lhsOk && rhsOk {
			if !iteration(lhsF, rhsF) {
				return false
			}
		}
	}
	return true
}
