package learn_golang

import (
	"fmt"
	"testing"
)

// go中只有值传递！！！！

// slice运行时是一个结构体，其中有一个指针指向底层的数组 @see reflect.SliceHeader
func TestSlice(t *testing.T) {
	var dataArray []Data
	fmt.Println(nil == dataArray)
	fmt.Printf("%v %p\n", dataArray, dataArray)

	var dataArrayPointer = &dataArray
	*dataArrayPointer = []Data{{1, "a"}}
	fmt.Printf("%v %p\n", dataArray, dataArray)

}

type Data struct {
	a int
	b string
}
