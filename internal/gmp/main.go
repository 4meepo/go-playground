package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	go func() {
		i++
	}()
	t := time.NewTimer(time.Second)
	<-t.C
	fmt.Println(i)
}
