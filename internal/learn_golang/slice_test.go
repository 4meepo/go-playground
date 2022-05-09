package learn_golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

	// append 修改的slice需要接收返回值
	intSlice := []int{1, 2, 3}
	intSlice2 := append(intSlice, 4)
	assert.Len(t, intSlice, 3)
	assert.Len(t, intSlice2, 4)

}

type Data struct {
	a int
	b string
}
