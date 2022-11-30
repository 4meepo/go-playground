package learn_golang

import (
	"fmt"
	"testing"
)

func TestDataRace(t *testing.T) {
	err := Foo()
	if err != nil {
		// ...
	}

	go func() {
		var y int
		y, err = Bar() // 捕获了外面的 err, 应该用 :=
		if err != nil {
			// ...
			fmt.Println(y)
		}
	}()

	var z int
	z, err = Bar() // GG, 并发写
	if err != nil {
		// ...
		fmt.Println(z)
	}
}

func Foo() error {
	return nil
}

func Bar() (int, error) {
	return 0, nil
}
