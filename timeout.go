package main

import (
	"fmt"
	"time"
)

func main() {
	done := time.After(1 * time.Millisecond)

	numbers := make(chan int)
	go func() {
		for n := 0; ; {
			numbers <- n
			n++
		}
	}()

myexit:
	for {
		select {
		case <-done:
			break myexit
		case num := <-numbers:
			fmt.Println(num)
		}
	}
}
