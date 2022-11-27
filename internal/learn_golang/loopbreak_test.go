package learn_golang

import (
	"fmt"
	"testing"
)

func TestBreakInLoop(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5}
loop:
	for _, d := range data {
		switch d {
		case 3:
			break loop
		}
		fmt.Println(d)
	}

}
