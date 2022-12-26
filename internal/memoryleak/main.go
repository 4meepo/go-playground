package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	bytes := make([]byte, 100000000)

	fmt.Println(len(bytes))
	b2 := make([]byte, 5)
	copy(b2, bytes)
	fmt.Println("cap of b2: ", cap(b2))

	printMem()

	time.Sleep(time.Second * 10)
	runtime.GC()

	printMem()

	runtime.KeepAlive(b2)

}

func printMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println(m.Alloc / 1024)

}
