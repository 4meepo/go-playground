package learn_golang

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeIdentity(t *testing.T) {
	// struct is comparable only when all its fields are comparable
	// struct is equal only when all its fields are equal
	{

		var a = struct {
			Name string
			Age  int
		}{
			"Tom",
			18,
		}

		var b = struct {
			Name string
			Age  int
		}{
			"Tom",
			18,
		}

		assert.True(t, a == b)

		b.Name = "Jerry"
		assert.False(t, a == b)
	}

	{
		var a = struct {
			Name string
			Map  map[string]int
		}{
			"Tom",
			nil,
		}
		var b = struct {
			Name string
			Map  map[string]int
		}{
			"Tom",
			nil,
		}
		assert.Nil(t, a.Map)
		assert.Nil(t, b.Map)

		// map is not comparable
		// so struct with map is not comparable
		// assert.False(t, a == b) // compile error
	}

}

type CustomTypeA []string

func TestTypeCompare(t *testing.T) {
	{

		aType := reflect.TypeOf(new(CustomTypeA))
		bType := reflect.TypeOf([]string{})

		fmt.Println(aType)
		fmt.Println(bType)

		assert.False(t, aType == bType)
	}
}

type customError struct {
	msg      string
	metadata map[string]string
}

func (e customError) Error() string {
	return e.msg
}

var customErrorA = customError{}

func TestErrorCompare(t *testing.T) {
	warpError := fmt.Errorf("warp error: %w", customErrorA)

	// customError is not comparable 因为它包含了一个map.
	assert.False(t, errors.Is(warpError, customErrorA))

}

func TestComparablePanics(t *testing.T) {
	var a, b error
	a = customError{}
	b = a

	// A comparison of two interface values with identical dynamic types
	// causes a run-time panic if values of that type are not comparable.
	assert.Panicsf(t, func() {
		_ = a == b
	}, "comparing uncomparable type error")

	assert.False(t, &a == &b)

	{
		var aa, bb error = (*customError)(nil), nil
		fmt.Println(reflect.TypeOf(aa))
		fmt.Println(reflect.TypeOf(bb))
		fmt.Println(reflect.TypeOf(nil))
		assert.False(t, aa == bb)
	}
	{
		var aa = (*customError)(nil)
		fmt.Println(reflect.TypeOf(aa))
		assert.True(t, aa == nil)
	}

}

type comparableType struct {
	name string
	age  int
}

func (comparableType) String() string {
	return "comparableType"
}

type uncomparableType struct {
	name string
	age  int
	meta map[string]string
}

func (uncomparableType) String() string {
	return "uncomparableType"
}

func TestCompareUncomparableStruct(t *testing.T) {
	{
		ca := comparableType{}
		cb := comparableType{}
		assert.True(t, ca == cb)

		assert.False(t, &ca == &cb)

		fmt.Println(reflect.TypeOf(&ca).Comparable())
	}

	{
		// compile error! uncomparableType is not comparable
		// ua := uncomparableType{}
		// ub := uncomparableType{}
		// assert.False(t, ua == ub)
	}

	{
		var aa fmt.Stringer = comparableType{}
		var bb fmt.Stringer = uncomparableType{}
		var cc fmt.Stringer = uncomparableType{}
		assert.False(t, aa == bb)
		assert.Panics(t, func() {
			//A comparison of two interface values with identical dynamic types causes a run-time panic if values of that type are not comparable. //
			// This behavior applies not only to direct interface value comparisons but also when comparing arrays of interface values or structs with interface-valued fields.
			_ = bb == cc
		}, "comparing uncomparable type error")
	}
}

type customTypeA struct {
	name string
}

func (customTypeA) String() string {
	return "customTypeA"
}

type customTypeB struct {
	name string
}

func (customTypeB) String() string {
	return "customTypeB"
}

func TestA(t *testing.T) {
	a := customTypeA{}
	aa := a
	fmt.Println(a == aa)
}
