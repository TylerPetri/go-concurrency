package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() // decrements wait group by 1

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg) // reference and pointer because really shouldn't be copying waitgroups
	}

	wg.Wait() // waits until the inital wg.Add(9) goes down to zero

	wg.Add(1) // without this, the wg.Done() in the printSomething() function turns the wait groups to a negative number. Thus add 1
	printSomething("This is the second thing that's being printed", &wg)

}
