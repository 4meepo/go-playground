package learn_golang

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	s := "H汉llo"
	for i, c := range s {
		// fmt.Println(i, string(s[i]))
		fmt.Println(i, string(c))
	}

	fmt.Println("")
	
	for i, c := range []rune(s) {
		fmt.Println(i, string(c))
	}
}
