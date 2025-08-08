package main

import (
	"fmt"
	"time"
)

func makeRequest(num int) <-chan string {
	responseChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		responseChan <- fmt.Sprintf("response number %d", num)
	}()
	return responseChan
}

//func chanAsPromise() {
//	firstResponseChan := makeRequest(1)
//	secondResponseChan := makeRequest(2)
//	fmt.Println("Non blocking")
//	fmt.Println(<-firstResponseChan, <-secondResponseChan)
//}

//func chanAsMutex() {
//	var counter int
//	mutexChan := make(chan struct{}, 1)
//	wg := sync.WaitGroup{}
//
//	for i := 0; i < 1000; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			mutexChan <- struct{}{}
//			counter++
//			<-mutexChan
//		}()
//	}
//	wg.Wait()
//	fmt.Println(counter)
//}

func main() {
	//chanAsMutex()
}
