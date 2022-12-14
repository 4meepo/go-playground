package learn_golang

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/golang/groupcache/lru"
)

func TestLRU(t *testing.T) {

	cache := lru.New(32)

	for i := 0; i < 100; i++ {
		cache.Add(i, i)
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i++ {
				v, ok := cache.Get(i % 100)
				if ok {
					fmt.Print(v)
				}
			}
		}()
	}

	<-time.After(time.Second)
	// data race
	go func() {
		cache.Add(2, 2)
	}()

	wg.Wait()

}
