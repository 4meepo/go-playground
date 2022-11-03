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

func TestSelectSlice(t *testing.T) {
	// 测试选择部分切片
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, 10, len(sl))

	// 选择前两个
	f2 := sl[:2]
	assert.Equal(t, []int{0, 1}, f2)

	// 选择最终两个
	l2 := sl[len(sl)-2:]
	assert.Equal(t, []int{8, 9}, l2) // 111111111111111

	// 未发生任何数组复制, 切片底层是同一个数组, 所以会影响原数组
	l2[0] = 100
	assert.Equal(t, 100, sl[8])

	m2 := sl[5:7:10]
	assert.Equal(t, 2, len(m2))
	assert.Equal(t, 5, cap(m2))
	assert.Equal(t, 5, m2[0])
	assert.Equal(t, 6, m2[1])

	// 如果访问数组下标超过len(), 会越界报错
	m2[1] = 99
	assert.Equal(t, 99, sl[6]) //  1111111111111111

}

func TestDeleteItem(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 删除第5,6
	assert.Equal(t, 10, cap(arr[:5])) // 第三个元素为空时，默认为cap()

	deleted := append(arr[:5:5], arr[7:]...)
	fmt.Println(deleted)
	fmt.Println(arr)
}

func TestDeclareSlice(t *testing.T) {
	var s []int
	assert.Nil(t, s, "这种声明方式的 slice 的值为nil")

	// s[0] = 1 // 这里会 index out of range
}
