package learn_golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	io_a = iota + 1
	io_b
	io_c
)

func TestIota(t *testing.T) {
	assert.Equal(t, 1, io_a)
	assert.Equal(t, 2, io_b)
	assert.Equal(t, 3, io_c)
}

// -------

// iota的作用域是一个 const 块儿中
const (
	ioo_a = 1 + iota
	ioo_b
	ioo_c
)

func TestIota2(t *testing.T) {
	assert.Equal(t, 1, ioo_a)
	assert.Equal(t, 2, ioo_b)
	assert.Equal(t, 3, ioo_c)
}
