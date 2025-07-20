package main

import (
	"fmt"
	"sync"
	"time"
)

func testCounterWithWg() {
	now := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	fmt.Println("start counter", counter)

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("finish counter:", counter)
	fmt.Println(time.Since(now).Seconds())
}

func main() {
	testCounterWithWg()
}
