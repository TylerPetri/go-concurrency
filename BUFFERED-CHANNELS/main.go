package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// print a got data message
		i := <-ch
		fmt.Println("Got", i, "from channel")

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10) // buffered channel: now able to have 10 things on channel (first 10 executed fast, then 1 every second like unbuffered)

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("sending", i, "to channel...")
		ch <- i
		fmt.Println("sent", i, "to channel...")
	}

	fmt.Println("Done!")
	close(ch)
}
