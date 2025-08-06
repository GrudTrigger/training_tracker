package main

import (
	"fmt"
	"time"
)

func square(ch chan int) {
	v := <- ch
	time.Sleep(1 * time.Second)
	ch <- v * v
}

func cube(ch chan int) {
	v := <- ch
	time.Sleep(1 * time.Second)
	ch <- v * v * v
}

func double(ch chan int) {
	v := <- ch
	ch <- v * 2
}

func main() {
	sqr_ch := make(chan int)
	cybe_ch := make(chan int)
	double_ch := make(chan int)

	go square(sqr_ch)
	go cube(cybe_ch)
	go double(double_ch)

	sqr_ch <- 3
	cybe_ch <- 5
	double_ch <- 10
	count := 0
	for {
		select {
		case v := <- sqr_ch:
			count++
			fmt.Printf("Square = %d\n", v)
			if(count == 3) {
				return
			}
		case v := <- cybe_ch:
			count++
			fmt.Printf("Cube = %d\n", v)
			if(count == 3) {
				return
			}
		case v := <- double_ch:
			count++
			fmt.Printf("Double = %d\n", v)
			if(count == 3) {
				return
			}
		case <- time.After(2 * time.Second):
			fmt.Println("time After")
			return
		}
	}
}