package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(ctx context.Context, task <-chan int, resultCh chan<- int, id int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-task:
			if !ok {
				return
			}
			resultCh <- v * v
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*20)
	defer cancel()

	wg := &sync.WaitGroup{}

	task, resultCh := make(chan int, 5), make(chan int, 5)

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, task, resultCh, i)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			task <- i
		}
		close(task)
	}()

	go func() {
		wg.Wait() // --- блокируется горутина ---
		close(resultCh)
	}()

	for res := range resultCh {
		fmt.Println(res)
	}
}

// 22 урок
