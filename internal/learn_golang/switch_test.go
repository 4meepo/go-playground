package learn_golang

import (
	"fmt"
	"testing"
)



// 测试switch和fallthorough的第一种使用方式（case为固定值，非表达式）
func TestSwitch(t *testing.T){
	var flag int = 3
	switch flag {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
		fallthrough   // fallthrough会穿透case一次。
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("default branch")
	}
	// Output:
	// 3
	// 4
}


// 测试带表达式的switch的使用,区别在于后者是在case中使用表达式
// 来计算true or false
func TestSwitchWithExpression(t *testing.T){
	var value int8 = 10
	switch {
	case value < 0:
		fmt.Println("less than zero")
	case value == 0:
		fmt.Println("equal to zero")
	case value> 0 :
		fmt.Println("greater than zero")
	default:
		fmt.Println("should not be executed")
	}
}