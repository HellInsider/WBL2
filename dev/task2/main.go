package main

import (
	"fmt"
	"testing"
)

func main() {
	var data string = "a4bc2d5e"
	decode(data)

	fmt.Printf("Testing...")
	errHandler := testing.T{}
	testDecode(&errHandler)

	fmt.Printf("Everything ok!")
}
