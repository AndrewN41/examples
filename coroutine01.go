package main

import (
	"fmt"
	"time"
)

const (
	fillCount  = 10 // number of elements in each input channel
	numWorkers = 3  // number of consumers.
)

func Wait() {
	time.Sleep(2000 * time.Millisecond)
}

func fillChannel(work chan string, name string) {
	for i := 0; i < fillCount; i++ {
		work <- fmt.Sprintf("%s%d", name, i)
	}
	close(work) // we're finished
}

func doWork(id int, ch1 chan string, ch2 chan string, done chan bool) {

	fmt.Println("Running worker", id)
	defer fmt.Println("Ending worker", id)

	for ch1 != nil || ch2 != nil {

		cnt1 := len(ch1)
		cnt2 := len(ch2)

		select {
		case str, more := <-ch1:
			if more {
				fmt.Printf("%d: ch1(%d) %s\n", id, cnt1, str)
			} else {
				ch1 = nil
				fmt.Printf("%d: ch1 closed\n", id)
			}

		case str, more := <-ch2:
			if more {
				fmt.Printf("%d: ch2(%d) %s\n", id, cnt2, str)
			} else {
				ch2 = nil
				fmt.Printf("%d: ch2 closed\n", id)
			}
		}
	}
	done <- true

}

func main() {

	a := make(chan string, 2) // a small channel
	b := make(chan string, 5) // a bigger channel

	// generate work
	go fillChannel(a, "A")
	go fillChannel(b, "B")

	// launch the consumers
	done := make(chan bool)

	for i := 0; i < numWorkers; i++ {
		go doWork(i, a, b, done)
	}

	// wait for the goroutines to finish.
	for i := 0; i < numWorkers; i++ {
		<-done
	}
	fmt.Println("All workers done.")

	Wait() // without this the defered prints from the workers doesn't flush
}
