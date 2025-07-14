package main

import "fmt"

func abc(args ...int) {
	fmt.Println(args)
}

func main() {
	abc(1,2,3,4,5)
}