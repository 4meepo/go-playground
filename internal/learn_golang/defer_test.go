package learn_golang

import (
	"fmt"
	"runtime"
	"testing"
)

func TestDefer(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}

}

func TestDeferPointerRec(t *testing.T) {
	a := &p2{"foo"}
	defer a.NameP() // bar
	defer a.Name()  // foo
	a.name = "bar"

	a = nil
}

type p2 struct {
	name string
}

func (p *p2) NameP() {
	fmt.Println(p.name)
}
func (p p2) Name() {
	fmt.Println(p.name)
}

func TestDeferClosure(t *testing.T) {
	var status string
	var code int
	defer func(code int) {
		fmt.Println(status) // 引用的外部变量是在运行时才确定
		fmt.Println(code)
	}(code) // args 在闭包中是立即计算的
	status = "foo"
	code = 500
	// output:
	// foo
	// 0
	runtime.KeepAlive(code)
}
