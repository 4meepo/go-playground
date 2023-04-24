package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用N个goroutine交替打印出1-N, 比如N=3， 打印123123123123....

func main() {

	const n = 3

	channels := make([]chan struct{}, 0, n)
	wg := &sync.WaitGroup{}

	for i := 1; i <= n; i++ {
		channels = append(channels, make(chan struct{}, 2))
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-channels[id-1]:
					fmt.Print(id)
					channels[id%n] <- struct{}{}
				default:
					<-time.After(time.Millisecond * 100)
				}
			}
		}(i)
	}

	channels[0] <- struct{}{}

	wg.Wait()

}
