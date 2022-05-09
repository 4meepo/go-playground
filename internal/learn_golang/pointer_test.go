package learn_golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	var d Dog
	changeValue(&d)
	assert.Equal(t, "dd", d.Name)
	d = Dog{"dog"}
	assert.Equal(t, "dog", d.Name)

	dp := &d
	*dp = Dog{"cat"}
	assert.Equal(t, "cat", d.Name)
}

func changeValue(a *Dog) {
	a.Name = "dd"
}

type Dog struct {
	Name string
}
