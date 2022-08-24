package main

import (
	"os"
	"task1/timeLib"
)

func main() {
	err := timeLib.ShowTime()
	if err != nil {
		os.Exit(1)
	}
	return
}
