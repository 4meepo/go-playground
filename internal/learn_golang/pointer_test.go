package learn_golang

import (
	"fmt"
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

	var nilIntPointer *int
	assert.Nil(t, nilIntPointer)

	pointerByNew := new(Dog)
	assert.NotNil(t, pointerByNew)
	assert.Equal(t, "", pointerByNew.Name)
}

func changeValue(a *Dog) {
	a.Name = "dd"
}

type Dog struct {
	Name string
}

func TestIntPointer(t *testing.T) {
	a := new(int)
	fmt.Println(a == nil)

	fn := func(ap *int) {
		*ap = 0
	}

	fn(a)

	fmt.Println(a == nil)
}
