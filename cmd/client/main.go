package main

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strings"

	"faraway/internal/lib/logger/sl"
	"faraway/internal/lib/pow"
)

var (
	addrFlag string
	envFlag  string
)

var errUsage = errors.New("usage")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %s [option...] COMMAND

Commands:
  connect

Options:
`,
			os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&addrFlag, "addr", "127.0.0.1:8088", "`address` to connect to")
	flag.StringVar(&envFlag, "env", "local", "`env` environment for logger (local|prod)")
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}
	log := sl.SetupLogger(envFlag)

	if err := run(log); err != nil {
		if errors.Is(err, errUsage) {
			flag.Usage()
			os.Exit(2)
		}
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func run(log *slog.Logger) error {
	conn, err := net.Dial("tcp", addrFlag)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Reading the PoW task from the server.
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Error("Error reading PoW task:", "error", err.Error())
		return err
	}
	nonce := string(buffer[:n])
	log.Info("Received PoW task:", "nonce", nonce)

	// Calculating the PoW solution.
	proof := pow.FindProofOfWork(nonce)
	fmt.Println()
	log.Info("Found PoW solution:", "proof", proof)

	// Sending the PoW solution to the server.
	_, err = conn.Write([]byte(proof))
	if err != nil {
		return err
	}

	// Read response from server.
	n, err = conn.Read(buffer)
	if err != nil {
		log.Error("Error reading server response:", "error", err.Error())
		return err
	}
	response := strings.ReplaceAll(string(buffer[:n]), "\u0000", "")
	log.Info("Server response:", "response", response)

	return nil
}
