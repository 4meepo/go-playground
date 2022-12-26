package learn_golang

import (
	"fmt"
	"testing"
)

func TestMergeChan(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int, 100)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		close(c1)
	}()
	go func() {
		for i := 10; i < 20; i++ {
			c2 <- i
		}
		close(c2)
	}()

	// go func() {
	for c1 != nil || c2 != nil {
		select {
		case v, open := <-c1:
			if !open {
				c1 = nil
				break
			}
			// <-time.After(100 * time.Millisecond)
			c3 <- v
		case v, open := <-c2:
			if !open {
				c2 = nil
				break
			}
			// <-time.After(100 * time.Millisecond)
			c3 <- v
		}
	}
	// }()

	// close(c3)
	for v := range c3 {
		fmt.Println(v)
	}
	fmt.Println("done") // will never be exectured, 因为 range loop channel时, channel只有关闭后, range才会退出

}

func TestReadClosedChannel(t *testing.T) {
	c := make(chan bool, 3)
	c <- true
	c <- true
	close(c)

	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)

	// 当且仅当 channel closed后, 读取到channel的零值时, 第二个参数为false
	v, ok = <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)

	// 关闭closed的channel,会panic
	//close(c)
}
