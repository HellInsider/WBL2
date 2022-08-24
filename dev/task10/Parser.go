package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

func ParseArgs(args []string) *telnet {

	if len(args) < 3 {
		log.Fatal("N.e. arguments")
	}

	var sTimeout string
	telnet := NewTelnet()

	flag.StringVar(&sTimeout, "timeout", "10s", "sets timeout")
	flag.Parse()

	telnet.timeout = ParseTimeout(sTimeout)
	telnet.host = args[len(args)-2]
	telnet.port = args[len(args)-1]

	return telnet
}

func ParseTimeout(s string) (t time.Duration) {
	var err error
	var i int
	if s[len(s)-2:] == "ms" {
		i, err = strconv.Atoi(s[:len(s)-2])
		t = time.Millisecond * time.Duration(i)
	} else if s[len(s)-1:] == "s" {
		i, err = strconv.Atoi(s[:len(s)-1])
		t = time.Second * time.Duration(i)
	}

	if err != nil {
		fmt.Println("Error while parsing timeout value.")
		t = 10 * time.Second
	}

	return t
}
