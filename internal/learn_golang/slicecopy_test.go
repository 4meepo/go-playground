package learn_golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceCopy(t *testing.T) {
	s1 := []int{0, 1, 2}
	s2 := make([]int, 0)
	copy(s2, s1)

	fmt.Println(s2)

	s3 := make([]int, 2)
	copy(s3, s1[1:3])
	fmt.Println(s3)

	s4 := []int{3, 6, 4, 6, 9, 0, 3, 2}
	s5 := s4[3:6:6]

	assert.Len(t, s5, 3)
	assert.Equal(t, 3, cap(s5))
	assert.Equal(t, s5[2], 0)
	s5[2] = 33
	assert.Equal(t, s4[5], 33)

}
