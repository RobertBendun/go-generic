package generic_test

import (
	"generic"
	"testing"
)

type exampleStruct1 struct {
	Field1 string
	Field2 int
	Field3 int
}

func TestModifyOnlyMatchingType(t *testing.T) {
	data := exampleStruct1{
		Field1: "foo",
		Field2: 1,
		Field3: 2,
	}
	generic.Modify(&data, func(x *int) { *x += 10 })
	if data.Field1 != "foo" {
		t.Errorf("Modification was incorrect, got: %s, want: %s", data.Field1, "foo")
	}
	if data.Field2 != 11 {
		t.Errorf("Modification was incorrect, got: %d, want: %d", data.Field2, 11)
	}
	if data.Field3 != 12 {
		t.Errorf("Modification was incorrect, got: %d, want: %d", data.Field3, 12)
	}
}
