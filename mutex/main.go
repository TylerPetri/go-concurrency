package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) { // again as pointer, we don't want to copy
	defer wg.Done()

	msg = s
}

func main() {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos!")
	wg.Wait()

	fmt.Println(msg)
}

// func updateMessage(s string, m *sync.Mutex) { // again as pointer, we don't want to copy
// 	defer wg.Done()

// 	m.Lock()
// 	msg = s
// 	m.Unlock() // "thread safe operation"
// }

// func main() {
// 	msg = "Hello, world!"

// 	var mutex sync.Mutex

// 	wg.Add(2)
// 	go updateMessage("Hello, universe!", &mutex)
// 	go updateMessage("Hello, cosmos!", &mutex)
// 	wg.Wait()

// 	fmt.Println(msg)
// }

// go run -race .
// no locks, two goroutines running at same time: one can finish before other it's random -> 1 data race (2 go routines)
// added mutex lock/unlock, now no race condition
