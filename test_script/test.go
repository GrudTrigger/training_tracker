package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomWait(ch chan int) {
	workSeconds := rand.Intn(5 + 1)
	time.Sleep(time.Duration(workSeconds) * time.Second)
	ch <- workSeconds
}

func main() {
	mainSeconds := time.Now()
	totalWordSeconds := 0
	ch := make(chan int)
	for range 100 {
		go randomWait(ch)
	}
	for range 100 {
		totalWordSeconds += <-ch
	}
	fmt.Println("main:", time.Since(mainSeconds))
	fmt.Println("total:", totalWordSeconds)
}
