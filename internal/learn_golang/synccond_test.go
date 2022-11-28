package learn_golang

import (
	"fmt"
	"sync"
	"testing"
)

// TestSyncCond 测试sync.Cond 包的用法
// 写一个捐赠的小场景
func TestSyncCond(t *testing.T) {
	type donation struct {
		cond    *sync.Cond
		balance int
	}

	d := &donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	fn := func(target int) {
		d.cond.L.Lock()
		for target <= d.balance {
			d.cond.Wait()
		}
		fmt.Printf("donation balance is %d\n", target)
		d.cond.L.Unlock()
	}

	go fn(10)
	go fn(15)

	for i := 0; i < 20; i++ {
		// <-time.After(time.Millisecond * 100)
		d.cond.L.Lock()
		d.balance++
		d.cond.L.Unlock()
		d.cond.Broadcast()
	}

	// type Donation struct {
	// 	cond    *sync.Cond
	// 	balance int
	// }

	// donation := &Donation{
	// 	cond: sync.NewCond(&sync.Mutex{}),
	// }

	// // Listener goroutines
	// f := func(goal int) {
	// 	donation.cond.L.Lock()
	// 	for donation.balance < goal {
	// 		donation.cond.Wait()
	// 	}
	// 	fmt.Printf("%d$ goal reached\n", donation.balance)
	// 	donation.cond.L.Unlock()
	// }
	// go f(10)
	// go f(15)

	// // Updater goroutine
	// for {
	// 	time.Sleep(time.Second)
	// 	donation.cond.L.Lock()
	// 	donation.balance++
	// 	donation.cond.L.Unlock()
	// 	donation.cond.Broadcast()
	// }

}
