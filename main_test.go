package main

import "testing"

func TestCallSayHello(t *testing.T) {
	name := "Vinicius"

	message, err := callSayHello(name)
	if err != nil {
		t.Fatal(err)
	}

	expectedMessage := "Hello Vinicius!"
	if message != expectedMessage {
		t.Errorf("got %q but expected %q", message, expectedMessage)
	}
}
