package learn_golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructValue(t *testing.T) {
	// 在对struct类型赋值时,go在内存中进行了内存的复制

	yifei := Person{"Yifei", &Address{"中潭路"}}
	assert.Equal(t, "Yifei", yifei.Name)

	yifei.Name = "Yifei2"
	assert.Equal(t, "Yifei2", yifei.Name)

	chengrui := yifei
	chengrui.Name = "chengrui"
	chengrui.Address.Street = "两湾城"
	assert.Equal(t, "chengrui", chengrui.Name)

	assert.Equal(t, "Yifei2", yifei.Name)

	assert.Equal(t, "两湾城", yifei.Address.Street)

}

type Person struct {
	Name    string
	Address *Address
}

type Address struct {
	Street string
}
