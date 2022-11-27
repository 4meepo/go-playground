package learn_golang

import (
	"fmt"
	"testing"
)

func Test进制(t *testing.T) {
	fmt.Println(0x10)  // 16
	fmt.Println(0o70)  // 56
	fmt.Println(0b100) // 4
}
