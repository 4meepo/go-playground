package learn_golang

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test进制(t *testing.T) {
	fmt.Println(0x10)  // 16
	fmt.Println(0o70)  // 56
	fmt.Println(0b100) // 4
}

func TestIntOverflow(t *testing.T) {
	var a struct {
		I int64
		N string
	}
	a.I = 1000
	a.N = "dsf"
	bs, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	fmt.Println("====== ", string(bs))

}
