package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {
	var n int64
	n = 256 * 1024 * 1024
	fmt.Println(n)
	debug.SetMemoryLimit(n)

	i := 0
	for {
		select {
		case <-time.After(time.Nanosecond * 1):
			//
		case <-time.After(time.Millisecond * 1):
			//
		case <-time.After(time.Second * 100):
			//
		}
		i++
		if i%10000 == 0 {
			printMem()
		}
	}

}
func printMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc/1024/1024, "MB")
}
