package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) { // <-chan is receive only, chan<- is send only, chan is both
	for {
		s, ok := <-ping
		if !ok {
			// do something
		}
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// create two channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		// wait for a response
		response := <-pong
		fmt.Println("Response:", response)
	}

	fmt.Println("All done. Closing channels.")
	close(ping)
	close(pong)
}
