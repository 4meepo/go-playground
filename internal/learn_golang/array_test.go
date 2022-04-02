package learn_golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeclareArray(t *testing.T) {
	var array1 [3]int
	assert.Len(t, array1, 3)

	var array2 []int
	assert.Len(t, array2, 0)

	var array3 = [...]int{1, 2, 3}
	assert.Len(t, array3, 3)
	assert.IsType(t, [3]int{}, array3)

	var a = 13
	var aa = &a
	assert.Equal(t, a, *aa)
	assert.Equal(t, a, 13)

	*aa = 14
	assert.Equal(t, a, 14)

}
