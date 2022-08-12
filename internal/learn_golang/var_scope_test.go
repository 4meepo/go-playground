package learn_golang

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarScopre(t *testing.T) {
	var a int = 30
	assert.Equal(t, 30, a)

	if true {
		// 因为这里用了 :=， 等于说是个新的局部变量，所以不会影响外部的变量a的值
		a, _ := strconv.Atoi("10")
		assert.Equal(t, 10, a)
	}
	// 所以这里还是 30
	assert.Equal(t, 30, a)
}
