package learn_golang

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSeed(t *testing.T) {
	rand.Seed(100)

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int())
	}
	// Output: rand包 的seed 如果相同，则生成的随机数序列永远相同
	// 7530908113823513298
	// 8856843482472596024
	// 557474217132159819
	// 2877121541507365039
	// 6731085038606570514
	// 1816137624592625432
	// 5282924661529872796
	// 4077108211906745666
	// 4808754755028350735
	// 8704958202260308128

	// 生成 [-100,100] 的随机数
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(101-(-100)) + (-100))
	}
}
