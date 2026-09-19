package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1. 不同的goroutine 会对counter资源竞争，造成数据累计的丢失
	/*var counter int
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter++
			}
		}()
	}
	// 阻塞当前goroutine,直接计算器为0
	wg.Wait()
	fmt.Println(counter)*/

	// 2.使用互斥锁解决资源竞争的问题
	/*var a adder
	wg := sync.WaitGroup{}
	for range 10 {
		wg.Add(1)
		go a.Add(&wg)
	}
	wg.Wait()
	fmt.Println(a.counter)*/

	// 3.使用原子锁解决竞态问题
	/*var counter atomic.Int64
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.Add(1)
			}
		}()

	}
	wg.Wait()
	fmt.Println(counter.Load())*/

	// 使用channel
	ch := make(chan int)
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				ch <- 1
			}
		}()
	}

	// 另起协助 去关联协程
	go func() {
		wg.Wait()
		close(ch)
	}()

	var counter int
	for val := range ch {
		counter += val
	}

	fmt.Println("counter:", counter)
}

type adder struct {
	mu      sync.Mutex
	counter int
}

func (a *adder) Add(wg *sync.WaitGroup) {
	a.mu.Lock()
	defer a.mu.Unlock()
	defer wg.Done()
	for range 1000 {
		a.counter++
	}
}
