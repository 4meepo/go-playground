package learn_golang

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m1 := make(map[string]string)
	m2 := m1

	fmt.Println(m2)
	// fmt.Println(m1 == m2)  compile error , map 只可以和nil 用来 == 比较

}
