package learn_golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	var a string
	assert.Equal(t, "", a)

	var b = "hello"
	var c = b
	assert.Equal(t, "hello", b)
	assert.Equal(t, "hello", c)

	c = "world"
	assert.Equal(t, "hello", b)
	assert.Equal(t, "world", c)

	bp := &b
	assert.Equal(t, "hello", *bp)

	*bp = "FFF"
	assert.Equal(t, "FFF", b)
}

// string 也是一个结构体，它的内部包含了一个指向底层数组的指针，
// 并且还有一个长度和容量，这两个属性都是不可变的。
// 在多个变量之间赋值时，传递的是结构体的值
func TestString2(t *testing.T) {
	var a = "hello"
	var b = a

	assert.Equal(t, "hello", a)
	assert.Equal(t, "hello", b)

	var ap = &a
	// 利用a的指针修改a的值
	*ap = "world"
	assert.Equal(t, "world", a)

	// b的值不变
	assert.Equal(t, "hello", b)
}
