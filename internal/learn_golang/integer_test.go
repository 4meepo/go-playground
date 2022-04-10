package learn_golang

import (
	"fmt"
	"testing"
)

func TestInteger(t *testing.T) {
	var aa int64 = -1          // 有符号数-1 即   11111111111111111111111111
	var bb uint64 = uint64(aa) // 转型后即为最大的uint64 1111111111111111111
	fmt.Println(bb)
	fmt.Println(bb + 1) // 溢出  expect 0

	// 直接输出最大的uint64
	fmt.Println(uint32(aa))

	// 测试uint8向上转型
	var cc uint8 = 255
	fmt.Println(uint64(cc)) // 还是255

	// 测试uint64向下转型
	var dd uint64 = 511
	fmt.Println(uint8(dd)) // 输出255，即会丢失精度！

	// 不同类型的数不能直接比较，需要转型
	// fmt.Println(dd == cc) // compile error!

}

// 测试int的使用
func TestInteger2(t *testing.T) {
	// 一个 int 的大小是基于操作系统特定实现的，它可能是 32 位，
	// 也可能是 64 位，因此从 int 转换为 int64 时，您不会丢失任何信息。
	var maxInt int = 1<<31 -1
	fmt.Println(maxInt)   // 64位机器上 expect: 2147483647

	var b uint64 = uint64(maxInt)
	fmt.Println(b)    // same as above
}
