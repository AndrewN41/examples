package main

import (
	"fmt"
	"time"
)

func Wait() {
	time.Sleep(2000 * time.Millisecond)
}

func Print(c chan string, d chan string, id int, done chan bool) {
	fmt.Println("Running go-routine", id)
	defer fmt.Println("Ending go-routine", id)

loop:
	for {
		fmt.Print(id, " len(c): ", len(c), " \"")
		select {
		case str, ok := <-c:
			fmt.Println(str, "\"")
			if !ok {
				break loop
			}
		case str, ok := <-d:
			fmt.Println(str, "\"")
			if !ok {
				break loop
			}
		}

	}
	done <- true
}

func fillChannel(ch chan string, name string) {
	ch <- name + "0"
	ch <- name + "1"
	ch <- name + "2"
	ch <- name + "3"
	close(ch)
}

func main() {
	numRoutines := 3

	done := make(chan bool)

	c := make(chan string, 2)
	d := make(chan string, 2)

	go fillChannel(c, "c")
	go fillChannel(d, "d")

	fmt.Println("Before go Print(c)")
	for i := 0; i < numRoutines; i++ {
		go Print(c, d, i, done)
	}
	fmt.Println("After go Print(c)")
	for i := 0; i < numRoutines; i++ {
		<-done
	}
	fmt.Println("All goroutines done.")
}
