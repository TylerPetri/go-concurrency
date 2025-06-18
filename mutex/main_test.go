package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	wg.Add(1)
	// go updateMessage("x")
	go updateMessage("Goobye, cruel world!")
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Error("incorrect value in msg")
	}
}

// wg.Add(2) -> test will fail, race condition
// go test -race .
