package learn_golang

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}

}
