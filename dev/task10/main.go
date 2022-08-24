package main

import "os"

func main() {
	telnet := ParseArgs(os.Args)
	telnet.Run()
}
