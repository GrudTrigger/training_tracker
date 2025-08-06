package main

import (
	"fmt"
	"time"
)


func main() {
	ch := make(chan string)

	go func(){
		for i := range 10000 {
			time.Sleep(300 * time.Millisecond)
			ch <- fmt.Sprintf("value = %d", i)
		}
		close(ch)
	}()

	go func (){
		for v:= range ch {
			fmt.Println(v)
		}
	}()
 // |(main)   |(g1)   |(g2)
	for range 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Выполняется какой-то следующий код, пока в фоне работает асинхронность")
	}
}

// посмотреть и прорешать вместе с автором про каналы(buff unbuff)
