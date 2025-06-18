package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	msg = "Hello, world!"

	greetings := []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world!",
	}

	for _, x := range greetings {
		wg.Add(1)
		go updateMessage(x)
		wg.Wait()
		printMessage()
	}
}

// tried wg.Add(len(greetings)), then wait() after loop... but still jumbled
// deadload if wg.Add(len(greetings)) before loop with wait() inside lopo and defer every time
