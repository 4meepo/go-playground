package learn_golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 参考 https://go.dev/doc/faq#nil_error
// An interface value is nil only if the V and T are both unset, (T=nil, V is not set),
// In particular, a nil interface will always hold a nil type.
// If we store a nil pointer of type *int inside an interface
// value, the inner type will be *int regardless of the value of the pointer: (T=*int, V=nil).
// Such an interface value will therefore be non-nil even when the pointer value V inside is nil
func TestNil(t *testing.T) {
	err1 := GiveMeAError1()
	fmt.Printf("err1: %T\n", err1)
	assert.True(t, nil == err1)

	err2 := GiveMeAError2()
	fmt.Printf("err2: %T\n", err2)
	assert.True(t, nil == err2)
	assert.True(t, (*MyError)(nil) == err2)

	err3 := GiveMeAError3()
	fmt.Printf("err3: %T\n", err3)
	assert.False(t, nil == err3)

	// interface的zero value == nil
	var err4 error
	fmt.Printf("err4: %T\n", err4)
	// In particular, a nil interface will always hold a nil type
	assert.IsType(t, nil, err4)
	assert.True(t, nil == err4)

	// 指针类型的zero value != nil
	var p1 *MyError
	fmt.Printf("p1: %T\n", p1)
	assert.IsType(t, (*MyError)(nil), p1)
	assert.True(t, nil == p1)

}

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

// 正确的写法√
func GiveMeAError1() error {
	return nil
}

func GiveMeAError2() *MyError {
	return nil
}

func GiveMeAError3() error {
	// 直接声明e，会导致e的type为*MyError 所以会发生 nil != nil
	var e *MyError
	return e
}

func TestNilPointer(t *testing.T) {
	p := giveMeAPerson()
	fmt.Println(p == nil)

	e := GiveMeAError2()
	fmt.Println(e == nil)
}

type person struct {
	name string
}

func giveMeAPerson() *person {
	var p *person
	return p
}

// 测试 nil != nil
func TestNilNotEqualsNil(t *testing.T) {
	e := returnANil(10)

	assert.False(t, e == nil)
	assert.True(t, e == (*MyError)(nil))
}

func returnANil(i int) error {
	var rtn *MyError
	if i >= 0 {
		return rtn
	}
	rtn = &MyError{}
	return rtn
}
