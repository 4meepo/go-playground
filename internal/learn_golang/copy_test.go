package learn_golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {

	a := []int{1, 2, 3}
	assert.Equal(t, len(a), 3)
	assert.Equal(t, cap(a), 3)

	b := make([]int, 1, 10)

	// 复制 a,b 中len较小的
	n := copy(b[1:], a)
	assert.Equal(t, 0, n)
}
