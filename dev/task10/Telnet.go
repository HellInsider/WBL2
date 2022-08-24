package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type telnet struct {
	input      io.Reader
	output     io.Writer
	timeout    time.Duration
	host       string
	port       string
	conn       net.Conn
	errHandler chan error
}

func NewTelnet() *telnet {
	var tel telnet
	tel.input = os.Stdin
	tel.output = os.Stdout
	tel.errHandler = make(chan error)
	return &tel
}

func (t *telnet) connect() {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(t.host, t.port), t.timeout)
	if err != nil {
		log.Fatalln("Connection error: ", err)
	}
	t.conn = conn
	fmt.Println("Connected!")
}

func (t *telnet) disconnect() {
	if err := t.conn.Close(); err != nil {
		log.Fatal("Disconnect error")
	}
}

func (t *telnet) send() error {
	if _, err := io.Copy(t.conn, t.input); err != nil {
		return err
	}
	log.Println("Msg sent")
	return nil
}

func (t *telnet) receive() error {
	if _, err := io.Copy(t.output, t.conn); err != nil {
		return err
	}
	log.Println("disconnect from server")
	return nil
}

func (t *telnet) Run() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	t.connect()

	go func() {
		if err := t.send(); err != nil {
			t.errHandler <- err
			log.Println(err)
		}
	}()

	go func() {
		if err := t.receive(); err != nil {
			t.errHandler <- err
			log.Println(err)
		}
	}()

	select {
	case err := <-t.errHandler:
		t.disconnect()
		log.Println("Err occured: ", err)
	case <-sigint:
		t.disconnect()
		log.Println("Interrupt execution entered")
	}
}
