package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
)

const (
	requestText = "Give me some wisdom, please"
)

var addrFlag string

var errUsage = errors.New("usage")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %s [option...] COMMAND

Commands:
  connect

Options:
`, os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&addrFlag, "addr", "127.0.0.1:8088", "`address` to connect to")
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	if err := run(); err != nil {
		if errors.Is(err, errUsage) {
			flag.Usage()
			os.Exit(2)
		}
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	conn, err := net.Dial("tcp", addrFlag)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send request to server.
	message := []byte(requestText)
	_, err = conn.Write(message)
	if err != nil {
		return err
	}

	fmt.Printf("Sent: %s\n", message)

	// Read response from server.
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	response := string(buffer[:n])
	fmt.Printf("Received: %s\n", response)

	return nil
}
