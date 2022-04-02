package learn_golang

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
	assert.Equal(t, nil, err1)

	err2 := GiveMeAError2()
	fmt.Printf("err2: %T\n", err2)
	assert.Equal(t, (*MyError)(nil), err2)

	err3 := GiveMeAError3()
	fmt.Printf("err3: %T\n", err3)
	assert.Equal(t, (*MyError)(nil), err3)

	// interface的zero value == nil
	var err4 error
	fmt.Printf("err4: %T\n", err4)
	// In particular, a nil interface will always hold a nil type
	assert.IsType(t, nil, err4)
	assert.Equal(t, nil, err4)

	// 指针类型的zero value != nil
	var p1 *MyError
	fmt.Printf("p1: %T\n", p1)
	assert.IsType(t, (*MyError)(nil), p1)
	assert.NotEqual(t, nil, p1)

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
	// 这里的nil的type会被隐式转换为 *MyError 所以会发生 nil != nil
	return nil
}

func GiveMeAError3() error {
	// 直接声明e，会导致e的type为*MyError 所以会发生 nil != nil
	var e *MyError
	return e
}
